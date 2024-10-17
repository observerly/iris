/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/astrotiff
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package astrotiff

/*****************************************************************************************************************/

import (
	"bytes"
	"compress/lzw"
	"compress/zlib"
	"encoding/binary"
	"image"
	"io"

	metadata "github.com/observerly/iris/pkg/ifd"
	"golang.org/x/image/tiff"
)

/*****************************************************************************************************************/

const (
	TiffLittleEndingHeader = "II\x2A\x00"
)

/*****************************************************************************************************************/

func FromCompressionType(c tiff.CompressionType) (t metadata.TagValueCompressionType) {
	switch c {
	case tiff.Uncompressed:
		return metadata.TagValueCompressionTypeNone
	case tiff.Deflate:
		return metadata.TagValueCompressionTypeDeflate
	case tiff.LZW:
		return metadata.TagValueCompressionTypeLZW
	case tiff.CCITTGroup3:
		return metadata.TagValueCompressionTypeG3
	case tiff.CCITTGroup4:
		return metadata.TagValueCompressionTypeG4
	default:
		return metadata.TagValueCompressionTypeNone
	}
}

/*****************************************************************************************************************/

func encodeGray(w io.Writer, pix []uint8, dx, dy, stride int, predictor bool) error {
	if !predictor {
		return writePixels(w, pix, dy, dx, stride)
	}
	buf := make([]byte, dx)
	for y := 0; y < dy; y++ {
		min := y*stride + 0
		max := y*stride + dx
		off := 0
		var v0 uint8
		for i := min; i < max; i++ {
			v1 := pix[i]
			buf[off] = v1 - v0
			v0 = v1
			off++
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

/*****************************************************************************************************************/

func encodeGray16(w io.Writer, pix []uint8, dx, dy, stride int, predictor bool) error {
	buf := make([]byte, dx*2)
	for y := 0; y < dy; y++ {
		min := y*stride + 0
		max := y*stride + dx*2
		off := 0
		var v0 uint16
		for i := min; i < max; i += 2 {
			// An image.Gray16's Pix is in big-endian order.
			v1 := uint16(pix[i])<<8 | uint16(pix[i+1])
			if predictor {
				v0, v1 = v1, v1-v0
			}
			// We only write little-endian TIFF files.
			buf[off+0] = byte(v1)
			buf[off+1] = byte(v1 >> 8)
			off += 2
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

/*****************************************************************************************************************/

func encodeRGBA(w io.Writer, pix []uint8, dx, dy, stride int, predictor bool) error {
	if !predictor {
		return writePixels(w, pix, dy, dx*4, stride)
	}
	buf := make([]byte, dx*4)
	for y := 0; y < dy; y++ {
		min := y*stride + 0
		max := y*stride + dx*4
		off := 0
		var r0, g0, b0, a0 uint8
		for i := min; i < max; i += 4 {
			r1, g1, b1, a1 := pix[i+0], pix[i+1], pix[i+2], pix[i+3]
			buf[off+0] = r1 - r0
			buf[off+1] = g1 - g0
			buf[off+2] = b1 - b0
			buf[off+3] = a1 - a0
			off += 4
			r0, g0, b0, a0 = r1, g1, b1, a1
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

/*****************************************************************************************************************/

func encodeRGBA64(w io.Writer, pix []uint8, dx, dy, stride int, predictor bool) error {
	buf := make([]byte, dx*8)
	for y := 0; y < dy; y++ {
		min := y*stride + 0
		max := y*stride + dx*8
		off := 0
		var r0, g0, b0, a0 uint16
		for i := min; i < max; i += 8 {
			// An image.RGBA64's Pix is in big-endian order.
			r1 := uint16(pix[i+0])<<8 | uint16(pix[i+1])
			g1 := uint16(pix[i+2])<<8 | uint16(pix[i+3])
			b1 := uint16(pix[i+4])<<8 | uint16(pix[i+5])
			a1 := uint16(pix[i+6])<<8 | uint16(pix[i+7])
			if predictor {
				r0, r1 = r1, r1-r0
				g0, g1 = g1, g1-g0
				b0, b1 = b1, b1-b0
				a0, a1 = a1, a1-a0
			}
			// We only write little-endian TIFF files.
			buf[off+0] = byte(r1)
			buf[off+1] = byte(r1 >> 8)
			buf[off+2] = byte(g1)
			buf[off+3] = byte(g1 >> 8)
			buf[off+4] = byte(b1)
			buf[off+5] = byte(b1 >> 8)
			buf[off+6] = byte(a1)
			buf[off+7] = byte(a1 >> 8)
			off += 8
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

/*****************************************************************************************************************/

func encode(w io.Writer, m image.Image, predictor bool) error {
	bounds := m.Bounds()
	buf := make([]byte, 4*bounds.Dx())
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		off := 0
		if predictor {
			var r0, g0, b0, a0 uint8
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := m.At(x, y).RGBA()
				r1 := uint8(r >> 8)
				g1 := uint8(g >> 8)
				b1 := uint8(b >> 8)
				a1 := uint8(a >> 8)
				buf[off+0] = r1 - r0
				buf[off+1] = g1 - g0
				buf[off+2] = b1 - b0
				buf[off+3] = a1 - a0
				off += 4
				r0, g0, b0, a0 = r1, g1, b1, a1
			}
		} else {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := m.At(x, y).RGBA()
				buf[off+0] = uint8(r >> 8)
				buf[off+1] = uint8(g >> 8)
				buf[off+2] = uint8(b >> 8)
				buf[off+3] = uint8(a >> 8)
				off += 4
			}
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

/*****************************************************************************************************************/

// Encode writes the image m to w. opt determines the options used for encoding, such as the compression
// type. If opt is nil, an uncompressed image is written.
func Encode(w io.Writer, m image.Image, opt *tiff.Options, ifdEntries []metadata.IFDEntry) error {
	d := m.Bounds().Size()

	_, err := io.WriteString(w, TiffLittleEndingHeader)

	if err != nil {
		return err
	}

	compression := tiff.Uncompressed

	predictor := false

	if opt != nil && opt.Compression != 0 {
		compression = opt.Compression
	}

	if opt != nil && opt.Predictor && compression == tiff.LZW {
		predictor = true
	}

	// Compressed data is written into a buffer first, so that we know the compressed size.
	var buf bytes.Buffer
	// dst holds the destination for the pixel data of the image - either w or a writer to buf.
	var dst io.Writer
	// imageLength is the length of the pixel data in bytes. The offset of the IFD is imageLength + 8 header bytes.
	var imageLength int

	switch compression {
	case tiff.Uncompressed:
		dst = w
		// Write IFD offset before outputting pixel data.
		switch m.(type) {
		case *image.Paletted:
			imageLength = d.X * d.Y * 1
		case *image.Gray:
			imageLength = d.X * d.Y * 1
		case *image.Gray16:
			imageLength = d.X * d.Y * 2
		case *image.RGBA64:
			imageLength = d.X * d.Y * 8
		case *image.NRGBA64:
			imageLength = d.X * d.Y * 8
		default:
			imageLength = d.X * d.Y * 4
		}
		err = binary.Write(w, enc, uint32(imageLength+8))
		if err != nil {
			return err
		}
	case tiff.Deflate:
		dst = zlib.NewWriter(&buf)
	case tiff.LZW:
		dst = lzw.NewWriter(&buf, lzw.MSB, 8)
	}

	pr := uint32(metadata.TagValuePredictorTypeNone)
	photometricInterpretation := uint32(metadata.TagValuePhotometricTypeRGB)
	samplesPerPixel := uint32(4)
	bitsPerSample := []uint32{8, 8, 8, 8}
	extraSamples := uint32(0)
	colorMap := []uint32{}

	if predictor {
		pr = uint32(metadata.TagValuePredictorTypeHorizontal)
	}

	switch m := m.(type) {
	case *image.Paletted:
		photometricInterpretation = uint32(metadata.TagValuePhotometricTypePaletted)
		samplesPerPixel = 1
		bitsPerSample = []uint32{8}
		colorMap = make([]uint32, 256*3)
		for i := 0; i < 256 && i < len(m.Palette); i++ {
			r, g, b, _ := m.Palette[i].RGBA()
			colorMap[i+0*256] = uint32(r)
			colorMap[i+1*256] = uint32(g)
			colorMap[i+2*256] = uint32(b)
		}
		err = encodeGray(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.Gray:
		photometricInterpretation = uint32(metadata.TagValuePhotometricTypeBlackIsZero)
		samplesPerPixel = 1
		bitsPerSample = []uint32{8}
		err = encodeGray(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.Gray16:
		photometricInterpretation = uint32(metadata.TagValuePhotometricTypeBlackIsZero)
		samplesPerPixel = 1
		bitsPerSample = []uint32{16}
		err = encodeGray16(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.NRGBA:
		extraSamples = 2 // Unassociated alpha.
		err = encodeRGBA(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.NRGBA64:
		extraSamples = 2 // Unassociated alpha.
		bitsPerSample = []uint32{16, 16, 16, 16}
		err = encodeRGBA64(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.RGBA:
		extraSamples = 1 // Associated alpha.
		err = encodeRGBA(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.RGBA64:
		extraSamples = 1 // Associated alpha.
		bitsPerSample = []uint32{16, 16, 16, 16}
		err = encodeRGBA64(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	default:
		extraSamples = 1 // Associated alpha.
		err = encode(dst, m, predictor)
	}

	if err != nil {
		return err
	}

	if compression != tiff.Uncompressed {
		if err = dst.(io.Closer).Close(); err != nil {
			return err
		}

		imageLength = buf.Len()

		if err = binary.Write(w, enc, uint32(imageLength+8)); err != nil {
			return err
		}

		if _, err = buf.WriteTo(w); err != nil {
			return err
		}
	}

	ifd := []metadata.IFDEntry{
		{
			Tag:      metadata.TagTypeImageWidth,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{uint32(d.X)},
		},
		{
			Tag:      metadata.TagTypeImageLength,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{uint32(d.Y)},
		},
		{
			Tag:      metadata.TagTypeBitsPerSample,
			DataType: metadata.DataTypeShort,
			Data:     bitsPerSample,
		},
		{
			Tag:      metadata.TagTypeCompression,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{uint32(FromCompressionType(compression))},
		},
		{
			Tag:      metadata.TagTypePhotometricInterpretation,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{photometricInterpretation},
		},
		{
			Tag:      metadata.TagTypeStripOffsets,
			DataType: metadata.DataTypeLong,
			Data:     []uint32{8},
		},
		{
			Tag:      metadata.TagTypeSamplesPerPixel,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{samplesPerPixel},
		},
		{
			Tag:      metadata.TagTypeRowsPerStrip,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{uint32(d.Y)},
		},
		{
			Tag:      metadata.TagTypeStripByteCounts,
			DataType: metadata.DataTypeLong,
			Data:     []uint32{uint32(imageLength)},
		},
		{
			Tag:      metadata.TagTypeXResolution,
			DataType: metadata.DataTypeRational,
			Data:     []uint32{72, 1},
		},
		{
			Tag:      metadata.TagTypeYResolution,
			DataType: metadata.DataTypeRational,
			Data:     []uint32{72, 1},
		},
		{
			Tag:      metadata.TagTypeResolutionUnit,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{uint32(metadata.TagValueResolutionUnitTypePerInch)},
		},
	}

	// Add predictor if needed:
	if pr != uint32(metadata.TagValuePredictorTypeNone) {
		ifd = append(ifd, metadata.IFDEntry{
			Tag:      metadata.TagTypePredictor,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{pr},
		})
	}

	// Add color map if needed:
	if len(colorMap) != 0 {
		ifd = append(ifd, metadata.IFDEntry{
			Tag:      metadata.TagTypeColorMap,
			DataType: metadata.DataTypeShort,
			Data:     colorMap,
		})
	}

	// Add extra samples if needed:
	if extraSamples > 0 {
		ifd = append(ifd, metadata.IFDEntry{
			Tag:      metadata.TagTypeExtraSamples,
			DataType: metadata.DataTypeShort,
			Data:     []uint32{extraSamples},
		})
	}

	// Extract and set the IFD entries from the options entry map:
	ifd = append(ifd, ifdEntries...)

	return writeIFD(w, imageLength+8, ifd)
}

/*****************************************************************************************************************/
