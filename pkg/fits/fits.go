package fits

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"strings"

	stats "github.com/observerly/iris/pkg/statistics"
)

const FITS_STANDARD = "FITS Standard 4.0"

// FITS Image struct:
// @see https://fits.gsfc.nasa.gov/fits_primer.html
// @see https://fits.gsfc.nasa.gov/standard40/fits_standard40aa-le.pdf
type FITSImage struct {
	ID       int          // Sequential ID number, for log output. Counted upwards from 0 for light frames. By convention, dark is -1 and flat is -2
	Filename string       // Original file name, if any, for log output.
	Header   FITSHeader   // The FITS Header with all keys, values, comments, history entries etc.
	Bitpix   int32        // Bits per pixel value from the header. Positive values are integral, negative floating.
	Bzero    float32      // Zero offset. (True pixel value is Bzero + Bscale * Data[i]).
	Bscale   float32      // Value scaler. (True pixel value is Bzero + Bscale * Data[i]).
	Naxisn   []int32      // Axis dimensions. Most quickly varying dimension first (i.e. X,Y)
	Pixels   int32        // Number of pixels in the image. Product of Naxisn[] or naxis1 and naxis2
	Data     []float32    // The image data
	ADU      int32        // The number of ADU (Analog to Digital Units) in the image.
	Exposure float32      // Image exposure in seconds
	Stats    *stats.Stats // Image statistics (mean, min, max, stdDev etc)
}

// Creates a new instance of FITS image initialized with empty header
func NewFITSImage(naxis int32, naxis1 int32, naxis2 int32, adu int32) *FITSImage {
	h := NewFITSHeader(naxis, naxis1, naxis2)

	return &FITSImage{
		Header: h,
		Bitpix: -32,
		Bzero:  0,
		Bscale: 1,
		ADU:    adu,
	}
}

// Creates a new instance of FITS image from given naxisn:
// (Data is not copied, allocated if nil. naxisn is deep copied)
func NewFITSImageFromNaxisn(naxisn []int32, data []float32, bitpix int32, naxis int32, naxis1 int32, naxis2 int32, adu int32) *FITSImage {
	pixels := int32(1)

	for _, naxis := range naxisn {
		pixels *= naxis
	}

	if data == nil {
		data = make([]float32, pixels)
	}

	f := NewFITSImage(naxis, naxis1, naxis2, adu)

	f.Stats = stats.NewStats(data, adu, int(naxis1))

	f.Header.Ints["ADU"] = struct {
		Value   int32
		Comment string
	}{
		Value:   adu,
		Comment: "Analog to Digital Units (ADU)",
	}

	return &FITSImage{
		ID:       f.ID,
		Filename: f.Filename,
		Header:   f.Header,
		Bitpix:   -32,
		Bzero:    f.Bzero,
		Bscale:   f.Bscale,
		Naxisn:   []int32{naxis1, naxis2},
		Pixels:   pixels,
		Data:     data,
		ADU:      adu,
		Exposure: 0,
		Stats:    f.Stats,
	}
}

// Creates a new instance of FITS image from given 2D exposure array
// (Data is not copied, allocated if nil. naxisn is deep copied)
func NewFITSImageFrom2DData(ex [][]uint32, naxis int32, naxis1 int32, naxis2 int32, adu int32) *FITSImage {
	pixels := naxis1 * naxis2

	data := make([]float32, 0)

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range ex {
		for _, col := range row {
			data = append(data, float32(col))
		}
	}

	if len(data) == 0 {
		data = make([]float32, pixels)
	}

	f := NewFITSImage(naxis, naxis1, naxis2, adu)

	f.Stats = stats.NewStats(data, adu, int(naxis1))

	f.Header.Ints["ADU"] = struct {
		Value   int32
		Comment string
	}{
		Value:   f.ADU,
		Comment: "Analog to Digital Units (ADU)",
	}

	return &FITSImage{
		ID:       f.ID,
		Filename: f.Filename,
		Header:   f.Header,
		Bitpix:   -32,
		Bzero:    f.Bzero,
		Bscale:   f.Bscale,
		Naxisn:   []int32{naxis1, naxis2},
		Pixels:   pixels,
		Data:     data,
		ADU:      adu,
		Exposure: 0,
		Stats:    f.Stats,
	}
}

// Read the FITS image from the given file.
func (f *FITSImage) Read(r io.Reader) error {
	// Read Header:
	err := f.Header.Read(r)

	if err != nil {
		return err
	}

	// Check that the mandatory SIMPLE header value exists as per FITS standard:
	if !f.Header.Bools["SIMPLE"].Value {
		return fmt.Errorf("%d: not a valid FITS file; SIMPLE=T missing in header", f.ID)
	}

	bitpix, ok := f.Header.Ints["BITPIX"]

	if !ok {
		return fmt.Errorf("%d: not a valid FITS Image file; BITPIX missing in header", f.ID)
	}

	// Check that the BITPIX value is valid for the IRIS module (only -32 supported):
	if bitpix.Value != -32 {
		return fmt.Errorf("%d: not a valid float32 FITS Image file; BITPIX must be -32", f.ID)
	}

	f.Header.Bitpix = bitpix.Value

	f.Bitpix = bitpix.Value

	naxis, ok := f.Header.Ints["NAXIS"]

	if !ok {
		return fmt.Errorf("%d: not a valid FITS Image file; NAXIS missing in header", f.ID)
	}

	f.Header.Naxis = naxis.Value

	naxis1, ok := f.Header.Ints["NAXIS1"]

	if !ok {
		return fmt.Errorf("%d: not a valid FITS Image file; NAXIS1 missing in header", f.ID)
	}

	// Set the NAXIS1 value:
	f.Header.Naxis1 = naxis1.Value

	naxis2, ok := f.Header.Ints["NAXIS2"]

	if !ok {
		return fmt.Errorf("%d: not a valid FITS Image file; NAXIS2 missing in header", f.ID)
	}

	// Set the NAXIS2 value:
	f.Header.Naxis2 = naxis2.Value

	// Set the NAXISn values:
	f.Naxisn = []int32{naxis1.Value, naxis2.Value}

	// Set the number of pixels:
	f.Pixels = f.Header.Naxis1 * f.Header.Naxis2

	data, err := readData(r, f.Bitpix, f.Pixels)

	if err != nil {
		return err
	}

	f.Data = data

	return nil
}

// Writes an in-memory FITS image to an io.Writer output stream
func (f *FITSImage) WriteToBuffer() (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	// Write the header:
	buf, err := f.Header.WriteToBuffer(buf)

	if err != nil {
		return nil, err
	}

	// Write the data:
	buf, err = writeFloat32ArrayToBuffer(buf, f.Data)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Writes FITS binary body data in network byte order to buffer
func writeFloat32ArrayToBuffer(buf *bytes.Buffer, data []float32) (*bytes.Buffer, error) {
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

/**
	Reads the FITS binary data from the given io.Reader stream and returns a
	slice of float32 values, or error

	Note: The data is read in network byte order and only supports 32-bitpix data
**/
func readData(r io.Reader, bitpix int32, pixels int32) ([]float32, error) {
	data := make([]float32, pixels)

	buf, err := io.ReadAll(r)

	// Convert []byte to bytes.Buffer:
	b := bytes.NewBuffer(buf)

	if err != nil {
		return nil, err
	}

	switch bitpix {

	// 32-bit floating point:
	case -32:
		err = readFloat32ArrayFromBuffer(b, data)

	// 64-bit floating point:
	case -64:
		// [TBI] Implement 64-bit floating point data type
		err = errors.New("64-bit floating point data not supported")

	// 8-bit unsigned integer:
	case 8:
		// [TBI] readUint8ArrayFromBuffer()
		err = errors.New("8-bit unsigned int data not supported")

	// 16-bit unsigned integer:
	case 16:
		// [TBI] readUint16ArrayFromBuffer()
		err = errors.New("16-bit unsigned int data not supported")

	// 32-bit unsigned integer:
	case 32:
		// [TBI] readUint32ArrayFromBuffer()
		err = errors.New("32-bit unsigned int data not supported")

		// 64-bit unsigned integer:
	case 64:
		// [TBI] readUint64ArrayFromBuffer()
		err = errors.New("64-bit unsigned int data not supported")
	}

	return data, err
}

// Reads FITS binary body float32 data in network byte order from buffer
func readFloat32ArrayFromBuffer(buf *bytes.Buffer, data []float32) error {
	return binary.Read(buf, binary.BigEndian, data)
}
