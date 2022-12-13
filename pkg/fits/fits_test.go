package fits

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"math"
	"os"
	"testing"
	"time"
)

func GetTestDataFromImage() ([][]uint32, image.Rectangle) {
	f, err := os.Open("../../images/noise16.jpeg")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	img, err := jpeg.Decode(f)

	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()

	data := make([][]uint32, bounds.Dx())

	for y := 0; y < bounds.Dy(); y++ {
		row := make([]uint32, bounds.Dx())
		data[y] = row
	}

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			data[j][i] = uint32(lum)
		}
	}

	return data, bounds
}

func TestNewDefaultFITSImageHeaderEnd(t *testing.T) {
	var img = NewFITSImage(2, 600, 800, 65535)

	var got = img.Header.End

	var want bool = false

	if got != want {
		t.Errorf("NewFITSImage() Header.End: got %v, want %v", got, want)
	}
}

func TestNewDefaultFITSImageBScale(t *testing.T) {
	var img = NewFITSImage(2, 600, 800, 65535)

	var got = img.Bscale

	var want float32 = 1

	if got != want {
		t.Errorf("NewFITSImage() Bscale: got %v, want %v", got, want)
	}
}

func TestNewFITSImageFromReader(t *testing.T) {
	// Attempt to open the file from the given filepath:
	file, err := os.Open("../../samples/noise16.fits")

	if err != nil {
		t.Errorf("NewFITSImageFromReader() os.Open(): %v", err)
	}

	// Defer closing the file:
	defer file.Close()

	// Attempt to read the file:
	fit := NewFITSImageFromReader(file)

	if fit.ADU != 65535 {
		t.Errorf("NewFITSImageFromReader() ADU: got %v, want %v", fit.ADU, 65535)
	}

	if fit.Bscale != 1 {
		t.Errorf("NewFITSImageFromReader() Bscale: got %v, want %v", fit.Bscale, 1)
	}

	if fit.Bzero != 0 {
		t.Errorf("NewFITSImageFromReader() Bzero: got %v, want %v", fit.Bzero, 32768)
	}

	if fit.Bitpix != -32 {
		t.Errorf("NewFITSImageFromReader() Bitpix: got %v, want %v", fit.Bzero, 32768)
	}

	// Check the header:
	if fit.Header.Naxis != 2 {
		t.Errorf("NewFITSImageFromReader() Header.Naxis: got %v, want %v", fit.Header.Naxis, 2)
	}

	if fit.Header.Naxis1 != 1463 {
		t.Errorf("NewFITSImageFromReader() Header.Naxis1: got %v, want %v", fit.Header.Naxis1, 600)
	}

	if fit.Header.Naxis2 != 1168 {
		t.Errorf("NewFITSImageFromReader() Header.Naxis2: got %v, want %v", fit.Header.Naxis2, 800)
	}
}

func TestNewFITSImageFrom2DDataID(t *testing.T) {
	var ex = [][]uint32{
		{1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	var img2DData = NewFITSImageFrom2DData(ex, 2, 16, 16, 255)

	var got = img2DData.ID

	var want int = 0

	if got != want {
		t.Errorf("NewFITSImageFrom2DData() ID: got %v, want %v", got, want)
	}
}

func TestNewFITSImageFrom2DDataPixels(t *testing.T) {
	var ex = [][]uint32{
		{1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	var img2DData = NewFITSImageFrom2DData(ex, 2, 16, 16, 255)

	var got = img2DData.Data

	var want = img2DData.Pixels

	if len(got) != int(want) {
		t.Errorf("NewFITSImageFrom2DData() Data Length should be 256 pixels: got %v, want %v", len(got), want)
	}
}

func TestNewFITSImageFrom2DDataData(t *testing.T) {
	var ex = [][]uint32{
		{1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	var img2DData = NewFITSImageFrom2DData(ex, 2, 16, 16, 255)

	var got = img2DData.Data

	var want int = 256

	if len(got) != want {
		t.Errorf("NewFITSImageFrom2DData() Data Length should be 256 pixels: got %v, want %v", got, want)
	}
}

func TestNewFITSImageFrom2DDataWriteFloatData(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	var fit = NewFITSImageFrom2DData(data, 2, int32(bounds.Dx()), int32(bounds.Dy()), 65535)

	var w io.Writer = os.Stdout

	buf := new(bytes.Buffer)

	buf, err := writeFloat32ArrayToBuffer(buf, fit.Data)

	if err != nil {
		t.Errorf("Error writing float32 array: %s", err)
	}

	if buf == nil {
		t.Errorf("Error writing float32 array: %s", err)
	}

	_, err = w.Write(buf.Bytes())

	if err != nil {
		t.Errorf("Error writing float32 array to standard output: %s", err)
	}
}

func TestNewFITSImageFrom2DDataWrite(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	var fit = NewFITSImageFrom2DData(data, 2, int32(bounds.Dx()), int32(bounds.Dy()), 65535)

	f, err := os.OpenFile("noise16.fits", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	defer f.Close()

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("noise16.fits")
	}()

	buf, err := fit.WriteToBuffer()

	if err != nil {
		t.Errorf("Error writing image: %s", err)
	}

	_, err = f.Write(buf.Bytes())

	if err != nil {
		t.Errorf("Error writing image: %s", err)
	}
}
func TestNewFITSImageFrom2DStats(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	var fit = NewFITSImageFrom2DData(data, 2, int32(bounds.Dx()), int32(bounds.Dy()), 65535)

	stats := fit.Stats

	if stats.ADU != 65535 {
		t.Errorf("Expected the ADU to be 65535, but got %d", stats.ADU)
	}

	if stats.Width != int(bounds.Dx()) {
		t.Errorf("Expected the width to be %d, but got %d", bounds.Dx(), stats.Width)
	}

	if stats.Min < 0 {
		t.Errorf("Expected the minimum pixel value to be greater than the minimum theoretical value")
	}

	if stats.Min != 0 {
		t.Errorf("Expected the minimum pixel value to be 0, but got %f", stats.Min)
	}

	if stats.Max != 65534 {
		t.Errorf("Expected the maximum pixel value to be 65534, but got %f", stats.Max)
	}

	if stats.Max > float32(stats.ADU) {
		t.Errorf("Expected that the maximum pixel value to be less than or equal to the maximum theoretical value, ADU")
	}

	if stats.Mean != 18514.215 {
		t.Errorf("Expected the mean pixel value to be 18514.215, but got %f", stats.Mean)
	}
}

func TestNewFITSRead(t *testing.T) {
	var fit = NewFITSImage(2, 1, 1, 65535)

	if fit == nil {
		t.Errorf("Expected the FITS image to be created, but got nil")
	}

	// Attempt to open the file from the given filepath:
	file, err := os.Open("../../samples/noise16.fits")

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	// Defer closing the file:
	defer file.Close()

	err = fit.Read(file)

	if err != nil {
		t.Errorf("Error reading FITS image: %s", err)
	}

	if fit == nil {
		t.Errorf("Expected the FITS image to be created, but got nil")
		return
	}

	if fit.Bitpix != -32 {
		t.Errorf("Expected the Bitpix to be -32, but got %d", fit.Bitpix)
	}

	if fit.Pixels != 1708784 {
		t.Errorf("Expected the Pixels to be 1, but got %d", fit.Pixels)
	}

	if fit.ADU != 65535 {
		t.Errorf("Expected the ADU to be 65535, but got %d", fit.ADU)
	}

	// Check that the mandatory DATAMIN header value exists and is set correctly as per FITS standard:
	if _, ok := fit.Header.Ints["DATAMIN"]; !ok {
		t.Errorf("Expected the DATAMIN header value to exist, but it does not")
	}

	if fit.Header.Ints["DATAMIN"].Value != 0 {
		t.Errorf("Expected the DATAMIN header value to be 0, but got %d", fit.Header.Ints["DATAMIN"].Value)
	}

	// Check that the mandatory DATAMAX header value exists as per FITS standard:
	if _, ok := fit.Header.Ints["DATAMAX"]; !ok {
		t.Errorf("Expected the DATAMAX header value to exist, but it does not")
	}

	if fit.Header.Ints["DATAMAX"].Value != fit.ADU {
		t.Errorf("Expected the DATAMAX header value to be %d, but got %d", fit.ADU, fit.Header.Ints["DATAMAX"].Value)
	}

	if fit.Header.Strings["PROGRAM"].Value != "@observerly/iris" {
		t.Errorf("Expected the PROGRAM to be @observerly/iris")
	}

	if !fit.Header.End {
		t.Errorf("Expected the End to be false, but got %t", fit.Header.End)
	}
}

func TestNewFITSFromFile(t *testing.T) {
	var fit = NewFITSImage(2, 1, 1, 65535)

	if fit == nil {
		t.Errorf("Expected the FITS image to be created, but got nil")
	}

	err := fit.ReadFromFile("../../samples/noise16.fits")

	if err != nil {
		t.Errorf("Error reading image: %s", err)
	}

	f, err := os.OpenFile("noise16.fits", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	defer f.Close()

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("noise16.fits")
	}()

	buf, err := fit.WriteToBuffer()

	if err != nil {
		t.Errorf("Error writing image: %s", err)
	}

	_, err = f.Write(buf.Bytes())

	if err != nil {
		t.Errorf("Error writing image: %s", err)
	}
}

func TestNewFindStarsFrom2DData(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	xs := bounds.Dx()

	ys := bounds.Dy()

	var fit = NewFITSImageFrom2DData(data, 2, int32(xs), int32(ys), 65535)

	radius := float32(16.0)

	sigma := float32(8.0)

	hfr := fit.ExtractHFR(radius, sigma, 2.0)

	if hfr == 0 {
		t.Error("Expected to calculate HFR, but got ", hfr)
	}

	if hfr > 8.0 {
		t.Error("Expected to calculate HFR less than 2.0, but got ", hfr)
	}

	if math.Abs(float64(hfr-6.601836)) > 0.000001 {
		t.Error("Expected to calculate HFR to an accuracy of 0.000001, but got ", hfr)
	}
}

func TestNewAddObservationEntry(t *testing.T) {
	var ex = [][]uint32{
		{1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	fit := NewFITSImageFrom2DData(ex, 2, 16, 16, 255)

	fit.AddObservationEntry(&FITSObservation{
		DateObs:    time.Date(2022, 5, 14, 0, 0, 0, 0, time.UTC),
		MJDObs:     59713,
		Equinox:    "2000.0 TT",
		Epoch:      "J2000",
		RA:         24.7122222,
		Dec:        41.2691667,
		Object:     "M31",
		Telescope:  "Namibiascope 1",
		Instrument: "20\" AG Optical iDK Planewave L-500s",
		Observer:   "Michael Roberts",
	})

	if fit.Header.Dates["DATE-OBS"].Value != "2022-05-14" {
		t.Errorf("Expected the DATE-OBS to be 2022-05-14, but got %s", fit.Header.Strings["DATE-OBS"].Value)
	}

	if fit.Header.Floats["MJD-OBS"].Value != 59713 {
		t.Errorf("Expected the MJD-OBS to be 59713, but got %f", fit.Header.Floats["MJD-OBS"].Value)
	}

	if fit.Header.Strings["EQUINOX"].Value != "2000.0 TT" {
		t.Errorf("Expected the EQUINOX to be 2000.0 TT, but got %s", fit.Header.Strings["EQUINOX"].Value)
	}

	if fit.Header.Strings["EPOCH"].Value != "J2000" {
		t.Errorf("Expected the EPOCH to be J2000, but got %s", fit.Header.Strings["EPOCH"].Value)
	}

	if fit.Header.Floats["RA"].Value != 24.7122222 {
		t.Errorf("Expected the RA to be 0.0, but got %f", fit.Header.Floats["RA"].Value)
	}

	if fit.Header.Floats["DEC"].Value != 41.2691667 {
		t.Errorf("Expected the DEC to be 41.2691667, but got %f", fit.Header.Floats["DEC"].Value)
	}

	if fit.Header.Strings["OBJECT"].Value != "M31" {
		t.Errorf("Expected the OBJECT to be M31, but got %s", fit.Header.Strings["OBJECT"].Value)
	}

	if fit.Header.Strings["TELESCOP"].Value != "Namibiascope 1" {
		t.Errorf("Expected the TELESCOP to be Namibiascope 1, but got %s", fit.Header.Strings["TELESCOP"].Value)
	}

	if fit.Header.Strings["INSTRUME"].Value != "20\" AG Optical iDK Planewave L-500s" {
		t.Errorf("Expected the INSTRUME to be 20\" AG Optical iDK Planewave L-500s, but got %s", fit.Header.Strings["INSTRUME"].Value)
	}

	if fit.Header.Strings["OBSERVER"].Value != "Michael Roberts" {
		t.Errorf("Expected the OBSERVER to be Michael Roberts, but got %s", fit.Header.Strings["OBSERVER"].Value)
	}
}
