package fits

import (
	"bytes"
	"encoding/binary"
	"strings"
)

const FITS_STANDARD = "FITS Standard 4.0"

// FITS Image struct:
// @see https://fits.gsfc.nasa.gov/fits_primer.html
// @see https://fits.gsfc.nasa.gov/standard40/fits_standard40aa-le.pdf
type FITSImage struct {
	ID       int        // Sequential ID number, for log output. Counted upwards from 0 for light frames. By convention, dark is -1 and flat is -2
	Filename string     // Original file name, if any, for log output.
	Header   FITSHeader // The FITS Header with all keys, values, comments, history entries etc.
	Bitpix   int32      // Bits per pixel value from the header. Positive values are integral, negative floating.
	Bzero    float32    // Zero offset. (True pixel value is Bzero + Bscale * Data[i]).
	Bscale   float32    // Value scaler. (True pixel value is Bzero + Bscale * Data[i]).
	Naxisn   []int32    // Axis dimensions. Most quickly varying dimension first (i.e. X,Y)
	Pixels   int32      // Number of pixels in the image. Product of Naxisn[] or naxis1 and naxis2
	Data     []float32  // The image data
	Exposure float32    // Image exposure in seconds
}

// Creates a new instance of FITS image initialized with empty header
func NewFITSImage(bitpix int32, naxis int32, naxis1 int32, naxis2 int32) *FITSImage {
	h := NewFITSHeader(bitpix, naxis, naxis1, naxis2)

	h.Ints["PCOUNT"] = struct {
		Value   int32
		Comment string
	}{
		Value:   0,
		Comment: "",
	}

	h.Ints["GCOUNT"] = struct {
		Value   int32
		Comment string
	}{
		Value:   1,
		Comment: "",
	}

	return &FITSImage{
		Header: h,
		Bscale: 1,
	}
}

// Creates a new instance of FITS image from given naxisn:
// (Data is not copied, allocated if nil. naxisn is deep copied)
func NewFITSImageFromNaxisn(naxisn []int32, data []float32, bitpix int32, naxis int32, naxis1 int32, naxis2 int32) *FITSImage {
	numPixels := int32(1)

	for _, naxis := range naxisn {
		numPixels *= naxis
	}

	if data == nil {
		data = make([]float32, numPixels)
	}

	h := NewFITSHeader(bitpix, naxis, naxis1, naxis2)

	h.Ints["PCOUNT"] = struct {
		Value   int32
		Comment string
	}{
		Value:   0,
		Comment: "",
	}

	h.Ints["GCOUNT"] = struct {
		Value   int32
		Comment string
	}{
		Value:   1,
		Comment: "",
	}

	return &FITSImage{
		ID:       0,
		Filename: "",
		Header:   h,
		Bitpix:   -32,
		Bzero:    0,
		Bscale:   1,
		Naxisn:   append([]int32(nil), naxisn...), // clone slice
		Pixels:   numPixels,
		Data:     data,
		Exposure: 0,
	}
}

// Creates a new instance of FITS image from given 2D exposure array
// (Data is not copied, allocated if nil. naxisn is deep copied)
func NewFITSImageFrom2DData(ex [][]uint32, bitpix int32, naxis int32, naxis1 int32, naxis2 int32) *FITSImage {
	pixels := naxis1 * naxis2

	var data []float32

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range ex {
		for _, col := range row {
			data = append(data, float32(col))
		}
	}

	if data == nil {
		data = make([]float32, pixels)
	}

	f := NewFITSImage(bitpix, naxis, naxis1, naxis2)

	return &FITSImage{
		ID:       f.ID,
		Filename: f.Filename,
		Header:   f.Header,
		Bitpix:   bitpix,
		Bzero:    f.Bzero,
		Bscale:   f.Bscale,
		Naxisn:   []int32{naxis1, naxis2},
		Pixels:   pixels,
		Data:     data,
		Exposure: 0,
	}
}

// Creates a new instance of FITS image from given image:
// (New data array will be allocated)
func NewFITSImageFromImage(img *FITSImage) *FITSImage {
	data := make([]float32, img.Pixels)

	return &FITSImage{
		ID:       img.ID,
		Filename: img.Filename,
		Header:   img.Header,
		Bitpix:   img.Bitpix,
		Bzero:    img.Bzero,
		Bscale:   img.Bscale,
		Naxisn:   append([]int32(nil), img.Naxisn...), // clone slice
		Pixels:   img.Pixels,
		Data:     data,
		Exposure: img.Exposure,
	}
}

// Writes FITS binary body data in network byte order to buffer
func writeFloat32ArrayToBuffer(data []float32) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, data)

	if err != nil {
		return nil, err
	}

	// Complete the last partial block, for strictly FITS compliant software
	totalBytes := len(data) << 2

	partial := totalBytes % 2880

	if partial != 0 {
		sb := strings.Builder{}

		for i := partial; i < 2880; i++ {
			sb.WriteRune(' ')
		}

		err := binary.Write(buf, binary.BigEndian, []byte(sb.String()))

		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}
