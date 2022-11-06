package iris

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"

	"github.com/observerly/iris/pkg/histogram"
)

var ex = [][]uint32{}

func TestNewMonochromeExposureWidth(t *testing.T) {
	mono := NewMonochromeExposure(ex, 1, 800, 600)

	var got int = mono.Width

	var want int = 800

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochromeExposureHeight(t *testing.T) {
	mono := NewMonochromeExposure(ex, 1, 800, 600)

	var got int = mono.Height

	var want int = 600

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochromeExposurePixels(t *testing.T) {
	mono := NewMonochromeExposure(ex, 1, 800, 600)

	var got int = mono.Pixels

	var want int = 480000

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochromeExposurePreprocess4x4(t *testing.T) {
	var ex = [][]uint32{
		{6, 6, 6, 6},
		{6, 7, 8, 8},
		{7, 8, 9, 7},
		{6, 7, 8, 6},
	}

	mono := NewMonochromeExposure(ex, 1, 4, 4)

	var x int = mono.Width

	var y int = mono.Height

	if x != 4 {
		t.Errorf("got %q, wanted %q", x, 4)
	}

	if y != 4 {
		t.Errorf("got %q, wanted %q", y, 4)
	}

	buff, err := mono.Preprocess()

	if err != nil {
		t.Errorf("Expected the buffer to be created successfully, but got %q", err)
	}

	f, err := os.Create("4x4image.jpg")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("4x4image.jpg")
	}()

	n, err := f.Write(buff.Bytes())

	if err != nil {
		t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
	}

	if n >= 512 {
		t.Errorf("Expected the number of bytes to be approximately less than 128, but got %v", n)
	}
}

func TestNewMonochromeExposurePreprocess16x16(t *testing.T) {
	var ex = [][]uint32{
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
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	mono := NewMonochromeExposure(ex, 1, 16, 16)

	var x int = mono.Width

	var y int = mono.Height

	if x != 16 {
		t.Errorf("got %q, wanted %q", x, 16)
	}

	if y != 16 {
		t.Errorf("got %q, wanted %q", y, 16)
	}

	buff, err := mono.Preprocess()

	if err != nil {
		t.Errorf("Expected the buffer to be created successfully, but got %q", err)
	}

	f, err := os.Create("16x16image.jpg")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("16x16image.jpg")
	}()

	n, err := f.Write(buff.Bytes())

	if err != nil {
		t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
	}

	if n >= 1086 {
		t.Errorf("Expected the number of bytes to be approximately less than 1086, but got %q", n)
	}
}

func TestNewMonochromeExposureOtsuThreshold(t *testing.T) {
	var ex = [][]uint32{
		{6, 6, 6, 6, 6, 6, 6, 6, 9, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 31, 35, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 34, 36, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 213, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 9, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 212, 211, 213, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 213, 214, 213, 10, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 89, 211, 212, 211, 8, 8, 8, 8, 8, 8, 9, 8, 8, 7, 6},
		{7, 71, 100, 108, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	mono := NewMonochromeExposure(ex, 1, 16, 16)

	var x int = mono.Width

	var y int = mono.Height

	if x != 16 {
		t.Errorf("got %q, wanted %q", x, 16)
	}

	if y != 16 {
		t.Errorf("got %q, wanted %q", y, 16)
	}

	mono.Preprocess()

	buff, err := mono.ApplyOtsuThreshold()

	if err != nil {
		t.Errorf("Expected the buffer to be created successfully, but got %q", err)
	}

	f, err := os.Create("16x16otsuimage.jpg")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("16x16otsuimage.jpg")
	}()

	n, err := f.Write(buff.Bytes())

	if err != nil {
		t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
	}

	if n >= 1086 {
		t.Errorf("Expected the number of bytes to be approximately less than 1086, but got %q", n)
	}
}

func TestNewMonochromeExposureNoiseReduction16x16(t *testing.T) {
	var ex = [][]uint32{
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
		{7, 8, 9, 7, 7, 7, 7, 7, 200, 200, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 200, 200, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	mono := NewMonochromeExposure(ex, 1, 16, 16)

	var x int = mono.Width

	var y int = mono.Height

	if x != 16 {
		t.Errorf("got %q, wanted %q", x, 16)
	}

	if y != 16 {
		t.Errorf("got %q, wanted %q", y, 16)
	}

	mono.Preprocess()

	buff, err := mono.ApplyNoiseReduction()

	if err != nil {
		t.Errorf("Expected the buffer to be created successfully, but got %q", err)
	}

	f, err := os.Create("16x16image.jpg")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("16x16image.jpg")
	}()

	n, err := f.Write(buff.Bytes())

	if err != nil {
		t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
	}

	if n >= 1086 {
		t.Errorf("Expected the number of bytes to be approximately less than 1086, but got %q", n)
	}

	if mono.Noise <= 0.0 {
		t.Errorf("Noise is %f, expected > 0.0", mono.Noise)
	}

	if mono.Noise > 255 {
		t.Errorf("Noise is %f, expected <= 255", mono.Noise)
	}
}

func TestNewMonochromeExposureHistogramGray(t *testing.T) {
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

	mono := NewMonochromeExposure(ex, 1, 16, 16)

	mono.Preprocess()

	res := histogram.HistogramGray(mono.Image)

	if res[1] != 1 {
		t.Errorf("got %q, wanted %q", res[1], 1)
	}

	if res[6] != 119 {
		t.Errorf("got %q, wanted %q", res[6], 119)
	}

	if res[7] != 64 {
		t.Errorf("got %q, wanted %q", res[7], 64)
	}

	if res[8] != 64 {
		t.Errorf("got %q, wanted %q", res[8], 64)
	}

	if res[9] != 8 {
		t.Errorf("got %q, wanted %q", res[9], 16)
	}
}

func TestNewNoiseExtractorGaussianNoisePngImage(t *testing.T) {
	f, err := os.Open("../../images/noise.jpeg")

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	defer f.Close()

	img, err := jpeg.Decode(f)

	if err != nil {
		t.Errorf("Error decoding image: %s", err)
	}

	bounds := img.Bounds()

	ex := make([][]uint32, bounds.Dx())

	for y := 0; y < bounds.Dy(); y++ {
		row := make([]uint32, bounds.Dx())
		ex[y] = row
	}

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			ex[j][i] = uint32(lum / 256)
		}
	}

	mono := NewMonochromeExposure(ex, 1, bounds.Dx(), bounds.Dy())

	mono.Preprocess()

	// Extract the noise from the image:
	bytes, err := mono.ApplyNoiseReduction()

	if err != nil {
		t.Errorf("Error extracting noise from image: %s", err)
	}

	// Save the image to the root folder:
	f, err = os.Create("noise.jpg")

	if err != nil {
		t.Errorf("Error creating image: %s", err)
	}

	defer f.Close()

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("noise.jpg")
	}()

	_, err = f.Write(bytes.Bytes())

	if err != nil {
		t.Errorf("Error writing image: %s", err)
	}
}

func TestNewMonochrome16NoiseExtractorGaussianNoise16PngImage(t *testing.T) {
	f, err := os.Open("../../images/noise.jpeg")

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	defer f.Close()

	img, err := jpeg.Decode(f)

	if err != nil {
		t.Errorf("Error decoding image: %s", err)
	}

	bounds := img.Bounds()

	ex := make([][]uint32, bounds.Dx())

	for y := 0; y < bounds.Dy(); y++ {
		row := make([]uint32, bounds.Dx())
		ex[y] = row
	}

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			ex[j][i] = uint32(lum / 256)
		}
	}

	mono := NewMonochromeExposure(ex, 1, bounds.Dx(), bounds.Dy())

	mono.Preprocess()

	// Extract the noise from the image:
	bytes, err := mono.ApplyNoiseReduction()

	if err != nil {
		t.Errorf("Error extracting noise from image: %s", err)
	}

	// Save the image to the root folder:
	f, err = os.Create("noise.jpg")

	if err != nil {
		t.Errorf("Error creating image: %s", err)
	}

	defer f.Close()

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("noise.jpg")
	}()

	_, err = f.Write(bytes.Bytes())

	if err != nil {
		t.Errorf("Error writing image: %s", err)
	}
}

func TestNewMonochromeExposureGetFITSImage(t *testing.T) {
	f, err := os.Open("../../images/noise.jpeg")

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	defer f.Close()

	img, err := jpeg.Decode(f)

	if err != nil {
		t.Errorf("Error decoding image: %s", err)
	}

	bounds := img.Bounds()

	ex := make([][]uint32, bounds.Dx())

	for y := 0; y < bounds.Dy(); y++ {
		row := make([]uint32, bounds.Dx())
		ex[y] = row
	}

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			ex[j][i] = uint32(lum / 256)
		}
	}

	mono := NewMonochromeExposure(ex, 255, bounds.Dx(), bounds.Dy())

	mono.Preprocess()

	fit := mono.GetFITSImage()

	if fit == nil {
		t.Errorf("Expected the FITS image to be instantiated successfully, but got nil")
	}

	if fit.Data == nil {
		t.Errorf("Expected the FITS image data to be instantiated successfully, but got nil")
	}

	if len(fit.Data) != bounds.Dx()*bounds.Dy() {
		t.Errorf("Expected the FITS image data to be %d, but got %d", bounds.Dx()*bounds.Dy(), len(fit.Data))
	}

	if fit.Header.Naxis1 != int32(bounds.Dx()) {
		t.Errorf("Expected the FITS image header NAXIS1 to be %q, but got %q", bounds.Dx(), fit.Header.Naxis1)
	}

	if fit.Header.Naxis2 != int32(bounds.Dy()) {
		t.Errorf("Expected the FITS image header NAXIS2 to be %q, but got %q", bounds.Dy(), fit.Header.Naxis2)
	}

	f, err = os.OpenFile("noisemonochrome.fits", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	defer f.Close()

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("noisemonochrome.fits")
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

func TestNewMonochromeExposureFromASCOMGetFITSImage(t *testing.T) {
	type CameraExposure struct {
		BayerXOffset int32      `json:"bayerXOffset"`
		BayerYOffset int32      `json:"bayerYOffset"`
		CCDXSize     int32      `json:"ccdXSize"`
		CCDYSize     int32      `json:"ccdYSize"`
		Image        [][]uint32 `json:"exposure"`
		MaxADU       int32      `json:"maxADU"`
		Rank         uint32     `json:"rank"`
		SensorType   string     `json:"sensorType"`
	}

	file, err := ioutil.ReadFile("../../data/m42-800x600.json")

	if err != nil {
		t.Errorf("Error opening from JSON data: %s", err)
	}

	data := CameraExposure{}

	_ = json.Unmarshal([]byte(file), &data)

	fmt.Println("Should be 800", len(data.Image))

	xs := 800

	ys := 600

	mono := NewMonochromeExposure(data.Image, 65535, xs, ys)

	mono.PreprocessImageArray(800, 600)

	fit := mono.GetFITSImage()

	if fit == nil {
		t.Errorf("Expected the FITS image to be instantiated successfully, but got nil")
	}

	if fit.Data == nil {
		t.Errorf("Expected the FITS image data to be instantiated successfully, but got nil")
	}

	if len(fit.Data) != xs*ys {
		t.Errorf("Expected the FITS image data to be %d, but got %d", xs*ys, len(fit.Data))
	}

	if fit.Header.Naxis1 != int32(xs) {
		t.Errorf("Expected the FITS image header NAXIS1 to be %q, but got %q", xs, fit.Header.Naxis1)
	}

	if fit.Header.Naxis2 != int32(ys) {
		t.Errorf("Expected the FITS image header NAXIS2 to be %q, but got %q", ys, fit.Header.Naxis2)
	}

	f, err := os.OpenFile("m42-800x600.fits", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	defer f.Close()

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("m42-800x600.fits")
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
