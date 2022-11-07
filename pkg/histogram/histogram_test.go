package histogram

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
	"testing"
)

func TestNewHistogramGray(t *testing.T) {
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

	gray := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			gray.Set(x, y, img.At(x, y))
		}
	}

	h := HistogramGray(gray)

	if len(h) != 256 {
		t.Errorf("Histogram length is not 256")
	}

	min, max := float64(h[0]), float64(h[0])
	for _, val := range h {
		min = math.Min(min, float64(val))
		max = math.Max(max, float64(val))
	}

	if min != 22 {
		t.Errorf("Histogram minimum is not 0")
	}

	if max != 259110 {
		t.Errorf("Histogram maximum is not 256")
	}
}

func TestNewHistogramGray16(t *testing.T) {
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

	gray := image.NewGray16(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			gray.Set(x, y, color.Gray16{uint16(lum)})
		}
	}

	h := HistogramGray16(gray)

	if len(h) != 65535 {
		t.Errorf("Histogram length is not 256")
	}

	min, max := float64(h[0]), float64(h[0])
	for _, val := range h {
		min = math.Min(min, float64(val))
		max = math.Max(max, float64(val))
	}

	if min != 0 {
		t.Errorf("Histogram minimum is not 0")
	}

	if max != 50189 {
		t.Errorf("Histogram maximum is not 256")
	}
}
