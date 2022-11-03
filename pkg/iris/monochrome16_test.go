package iris

import (
	"image/jpeg"
	"os"
	"testing"
)

var ex16 = [][]uint32{}

func TestNewMonochrome16ExposureWidth(t *testing.T) {
	mono := NewMonochrome16Exposure(ex16, 1, 800, 600)

	var got int = mono.Width

	var want int = 800

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochrome16ExposureHeight(t *testing.T) {
	mono := NewMonochrome16Exposure(ex16, 1, 800, 600)

	var got int = mono.Height

	var want int = 600

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochrome16ExposurePixels(t *testing.T) {
	mono := NewMonochrome16Exposure(ex16, 1, 800, 600)

	var got int = mono.Pixels

	var want int = 480000

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochrome16ExposureGetBuffer(t *testing.T) {
	mono := NewMonochrome16Exposure(ex16, 1, 800, 600)

	_, err := mono.GetBuffer(mono.Image)

	if err != nil {
		t.Errorf("Expected no error when creating the output buffer, got %q", err)
	}
}

func TestNewMonochrome16ExposurePreprocess4x4(t *testing.T) {
	var ex = [][]uint32{
		{6, 6, 6, 6},
		{6, 7, 8, 8},
		{7, 8, 9, 7},
		{6, 7, 8, 6},
	}

	mono := NewMonochrome16Exposure(ex, 1, 4, 4)

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

	if n >= 1024 {
		t.Errorf("Expected the number of bytes to be approximately less than 1024, but got %v", n)
	}
}

func TestNewMonochrome16ExposurePreprocess16x16(t *testing.T) {
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

	mono := NewMonochrome16Exposure(ex, 1, 16, 16)

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

func TestNewNoiseExtractorGaussianNoise16PngImage(t *testing.T) {
	f, err := os.Open("../../images/noise16.jpeg")

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

	for x := 0; x < bounds.Dx(); x++ {
		col := make([]uint32, bounds.Dy())
		ex[x] = col
	}

	mono := NewMonochrome16Exposure(ex, 1, bounds.Dx(), bounds.Dy())

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			mono.Raw[i][j] = uint32(lum)
		}
	}

	mono.Preprocess()

	// Extract the noise from the image:
	bytes, err := mono.ApplyNoiseReduction()

	if err != nil {
		t.Errorf("Error extracting noise from image: %s", err)
	}

	// Save the image to the root folder:
	f, err = os.Create("noise16.jpg")

	if err != nil {
		t.Errorf("Error creating image: %s", err)
	}

	defer f.Close()

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("noise16.jpg")
	}()

	_, err = f.Write(bytes.Bytes())

	if err != nil {
		t.Errorf("Error writing image: %s", err)
	}
}

func TestNewMonochrome16ExposureOtsuThreshold(t *testing.T) {
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

	mono := NewMonochrome16Exposure(ex, 1, 16, 16)

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

func TestNewMonochrome16ExposureNoiseReduction16x16(t *testing.T) {
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

	mono := NewMonochrome16Exposure(ex, 1, 16, 16)

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

func TestNewMonochrome16ExposureGetFITSImage(t *testing.T) {
	f, err := os.Open("../../images/noise16.jpeg")

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

	for x := 0; x < bounds.Dx(); x++ {
		col := make([]uint32, bounds.Dy())
		ex[x] = col
	}

	mono := NewMonochrome16Exposure(ex, 65535, bounds.Dx(), bounds.Dy())

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			mono.Raw[i][j] = uint32(lum)
		}
	}

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

	if fit.Header.Ints["NAXIS1"].Value != int32(bounds.Dx()) {
		t.Errorf("Expected the FITS image header NAXIS1 to be %q, but got %q", bounds.Dx(), fit.Header.Ints["NAXIS1"])
	}

	if fit.Header.Ints["NAXIS2"].Value != int32(bounds.Dy()) {
		t.Errorf("Expected the FITS image header NAXIS2 to be %q, but got %q", bounds.Dy(), fit.Header.Ints["NAXIS2"])
	}
}
