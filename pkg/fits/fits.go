/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/fits
//	@license	Copyright © 2021-2025 observerly

/*****************************************************************************************************************/

package fits

/*****************************************************************************************************************/

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/observerly/iris/pkg/photometry"
	stats "github.com/observerly/iris/pkg/statistics"
	"github.com/observerly/iris/pkg/utils"
)

/*****************************************************************************************************************/

const FITS_STANDARD = "FITS Standard 4.0"

/*****************************************************************************************************************/

// Represents a FITS Image
//
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

/*****************************************************************************************************************/

// Represents a FITS Observation
//
// @see https://fits.gsfc.nasa.gov/standard40/fits_standard40aa-le.pdf
type FITSObservation struct {
	DateObs    time.Time `json:"dateObs"`    // Date of observation e.g., 2022-05-15
	MJDObs     float32   `json:"mjdObs"`     // Modified Julian Date (JD − 2,400,000.5) of the observation
	Equinox    float32   `json:"equinox"`    // Equinox of observation e.g., J2000.0
	Epoch      float32   `json:"epoch"`      // Epoch of observation e.g., J2022.0
	RA         float32   `json:"ra"`         // Right Ascension of observation
	Dec        float32   `json:"dec"`        // Declination of observation
	Altitude   float32   `json:"altitude"`   // Altitude of the observation
	Azimuth    float32   `json:"azimuth"`    // Azimuth of the observation
	Airmass    float32   `json:"airmass"`    // Airmass of the observation
	Refraction float32   `json:"refraction"` // Refraction of the observation
	Object     string    `json:"object"`     // The name for the object observed
	Telescope  string    `json:"telescope"`  // The telescope used to acquire the data
	Instrument string    `json:"instrument"` // The instrument used to acquire the data
	Observer   string    `json:"observer"`   // Who acquired the data
}

/*****************************************************************************************************************/

// Represents a FITS Observer
type FITSObserver struct {
	Latitude  float32 `json:"latitude"`  // Latitude of the observer
	Longitude float32 `json:"longitude"` // Longitude of the observer
	Elevation float32 `json:"elevation"` // Elevation of the observer
}

/*****************************************************************************************************************/

// Creates a new instance of FITS image initialized with empty header
func NewFITSImage(naxis int32, naxis1 int32, naxis2 int32, adu int32) *FITSImage {
	h := NewFITSHeader(naxis, naxis1, naxis2)

	// Set the ADU (Analog to Digital Units) value:
	h.Set("ADU", adu, "Analog to Digital Units (ADU)")

	// Set the minimum valid physical value represented by the array:
	h.Set("DATAMIN", int32(0), "The minimum valid physical value represented by the array")

	// Set the maximum valid physical value represented by the array:
	h.Set("DATAMAX", adu, "The maximum valid physical value represented by the array")

	return &FITSImage{
		Header: h,
		Bitpix: -32,
		Bzero:  0,
		Bscale: 1,
		ADU:    adu,
	}
}

/*****************************************************************************************************************/

// Creates a new instance of FITS image initialized from an io.Reader:
func NewFITSImageFromReader(r io.Reader) *FITSImage {
	// Construct a blank FITS Image:
	f := NewFITSImage(0, 0, 0, 0)

	// Read appropriate data from the io.Reader:
	err := f.Read(r)

	if err != nil {
		return nil
	}

	return f
}

/*****************************************************************************************************************/

// Creates a new instance of FITS image from given 2D exposure array
// (Data is not copied, allocated if nil. naxisn is deep copied)
func NewFITSImageFrom2DData(ex [][]uint32, naxis int32, naxis1 int32, naxis2 int32, adu int32) *FITSImage {
	pixels := naxis1 * naxis2

	// Flatten the 2D Colour Filter Array array into a 1D array:
	data := utils.Flatten2DUInt32Array(ex)

	if len(data) == 0 {
		data = make([]float32, pixels)
	}

	f := NewFITSImage(naxis, naxis1, naxis2, adu)

	f.Stats = stats.NewStats(data, adu, int(naxis1))

	// Set the ADU (Analog to Digital Units) value:
	f.Header.Set("ADU", adu, "Analog to Digital Units (ADU)")

	// Set the minimum valid physical value represented by the array:
	f.Header.Set("DATAMIN", int32(0), "The minimum valid physical value represented by the array")

	// Set the maximum valid physical value represented by the array:
	f.Header.Set("DATAMAX", adu, "The maximum valid physical value represented by the array")

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

/*****************************************************************************************************************/

func (f *FITSImage) AddObservationEntry(observation *FITSObservation) *FITSImage {
	// Set the Object Name:
	f.Header.Set("OBJECT", observation.Object, "The name for the object observed")

	// Set the Date of Observation:
	f.Header.Set("DATE-OBS", observation.DateObs.Format("2006-01-02"), "Date of observation")

	// Set the Julian Date of the Observation:
	f.Header.Set("JD-OBS", observation.MJDObs+2400000.5, "Julian Date of the observation")

	// Set the Modified Julian Date of the Observation:
	f.Header.Set("MJD-OBS", observation.MJDObs, "Modified Julian Date of the observation")

	// Set the Equinox of the Observation:
	f.Header.Set("EQUINOX", fmt.Sprintf("%.1f", observation.Equinox), "Equinox of observation e.g., J2000.0")

	// Set the Epoch of the Observation:
	f.Header.Set("EPOCH", fmt.Sprintf("%.1f", observation.Epoch), "Epoch of observation")

	// Set the astrometric reference frame or celestial coordinate system used for the celestial coordinates:
	f.Header.Set("RADESYS", "ICRS", "International Celestial Reference System")

	// Set the Right Ascension of the Observation:
	f.Header.Set("RA", observation.RA, "Right Ascension (in degrees) at J2000.0")

	// Set the Declination of the Observation:
	f.Header.Set("DEC", observation.Dec, "Declination (in degrees) at J2000.0")

	// Set the local Altitude of the Observation:
	f.Header.Set("ALT", observation.Altitude, "Altitude (in degrees) of the observation")

	// Set the local Azimuth of the Observation:]
	f.Header.Set("AZ", observation.Azimuth, "Azimuth (in degrees) of the observation")

	// Set the Airmass of the Observation:
	f.Header.Set("AIRMASS", observation.Airmass, "Airmass of the observation (sec z)")

	// Set the Refractive Correction of the Observation:
	f.Header.Set("REFRACT", observation.Refraction, "Refraction correction (in degrees) of the observation")

	// Set the Telescope Name:
	f.Header.Set("TELESCOP", observation.Telescope, "The name of the telescope")

	// Set the Instrument Name:
	f.Header.Set("INSTRUME", observation.Instrument, "The name of the instrument")

	// Set the Observer Name:
	f.Header.Set("OBSERVER", observation.Observer, "Who owns the observation data")

	return f
}

/*****************************************************************************************************************/

func (f *FITSImage) AddObserverEntry(observer *FITSObserver) *FITSImage {
	// Set the Latitude of the Observer (in degrees):
	f.Header.Set("LATITUDE", observer.Latitude, "Latitude of the observer (in degrees)")

	// Set the Latitude of the Observer (in degrees):
	f.Header.Set("LONGITUD", observer.Longitude, "Longitude of the observer (in degrees)")

	// Set the elevation of the Observer (in meters):
	f.Header.Set("ELEVATIO", observer.Elevation, "Elevation of the observer (in meters)")

	return f
}

/*****************************************************************************************************************/

func (f *FITSImage) ExtractHFR(radius float32, sigma float32, starInOut float32) float32 {
	se := photometry.NewStarsExtractor(f.Data, int(f.Naxisn[0]), int(f.Naxisn[1]), radius, f.ADU)

	se.FindStars(f.Stats, sigma, starInOut)

	se.Stars = nil

	return se.HFR
}

/*****************************************************************************************************************/

func (f *FITSImage) ReadFromFile(fp string) error {
	// Check that the filename is not empty:
	if fp == "" {
		return fmt.Errorf("the filepath provided is empty")
	}

	// Attempt to open the file from the given filepath:
	file, err := os.Open(fp)

	if err != nil {
		return err
	}

	// Defer closing the file:
	defer file.Close()

	// Set the filename:
	f.Filename = path.Base(fp)

	return f.Read(file)
}

/*****************************************************************************************************************/

// Read the FITS image from the given file.
func (f *FITSImage) Read(r io.Reader) error {
	// Read Header:
	err := f.Header.Read(r)

	if err != nil {
		return err
	}

	// Check that the mandatory SIMPLE header OR XTENSION header value exists as per FITS standard:
	if !f.Header.Bools["SIMPLE"].Value && strings.TrimSpace(f.Header.Strings["XTENSION"].Value) != "IMAGE" {
		return fmt.Errorf("%d: not a valid FITS file; SIMPLE=T or XTENSION missing in header", f.ID)
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

	adu, ok := f.Header.Ints["ADU"]

	if !ok {
		return fmt.Errorf("%d: not a valid FITS Image file; ADU missing in header", f.ID)
	}

	f.ADU = adu.Value

	data, err := readData(r, f.Bitpix, f.Pixels)

	if err != nil {
		return err
	}

	f.Data = data

	return nil
}

/*****************************************************************************************************************/

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

/*****************************************************************************************************************/

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
			sb.WriteRune(rune(0x00))
		}

		err := binary.Write(buf, binary.BigEndian, []byte(sb.String()))

		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}

/*****************************************************************************************************************/

// Reads the FITS binary data from the given io.Reader stream and returns a slice of float32 values, or error.
// Note: The data is read in network byte order and only supports 32-bitpix data
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

/*****************************************************************************************************************/

// Reads FITS binary body float32 data in network byte order from buffer
func readFloat32ArrayFromBuffer(buf *bytes.Buffer, data []float32) error {
	return binary.Read(buf, binary.BigEndian, data)
}

/*****************************************************************************************************************/
