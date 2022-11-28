package photometry

import (
	"image"
	"image/jpeg"
	"os"
	"testing"

	stats "github.com/observerly/iris/pkg/statistics"
	"github.com/observerly/iris/pkg/utils"
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

func TestNewStarsExtractor(t *testing.T) {
	var ex = [][]uint32{
		{123, 6, 117, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{89, 123, 81, 123, 8, 128, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{123, 8, 82, 7, 89, 7, 97, 7, 111, 7, 7, 7, 7, 9, 8, 7},
		{6, 123, 8, 129, 6, 114, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{87, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 129, 8, 212, 8, 117, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 111, 9, 7, 7, 7, 7, 7, 7, 7, 7, 121, 7, 9, 8, 7},
		{102, 7, 8, 6, 111, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 98, 8, 108, 8, 173, 8, 8, 123, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 109, 6, 105, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 121, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 191},
	}

	xs := 16

	ys := 16

	data := make([]float32, xs*ys)

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range ex {
		for _, col := range row {
			data = append(data, float32(col))
		}
	}

	s := NewStarsExtractor(data, xs, ys, 2.5, 255)

	if s.Height != 16 {
		t.Errorf("Height is %d, expected 16", s.Height)
	}

	if s.Width != 16 {
		t.Errorf("Width is %d, expected 16", s.Width)
	}

	if s.Radius != 2.5 {
		t.Errorf("Radius is %f, expected 2.5", s.Radius)
	}

	if s.Threshold != 0 {
		t.Errorf("Threshold is %f, expected 0", s.Threshold)
	}

	if s.ADU != 255 {
		t.Errorf("ADU is %d, expected 255", s.ADU)
	}

	if s.Stars == nil {
		t.Errorf("Expected there to be a holding array for stars")
	}
}

func TestNewGetBrightPixels(t *testing.T) {
	var ex = [][]uint32{
		{123, 6, 117, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{89, 123, 81, 123, 8, 128, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{123, 8, 82, 7, 89, 7, 97, 7, 111, 7, 7, 7, 7, 9, 8, 7},
		{6, 123, 8, 129, 6, 114, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{87, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 129, 8, 212, 8, 117, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 111, 9, 7, 7, 7, 7, 7, 7, 7, 7, 121, 7, 9, 8, 7},
		{102, 7, 8, 6, 111, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 98, 8, 108, 8, 173, 8, 8, 123, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 109, 6, 105, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 121, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 191},
	}

	xs := 16

	ys := 16

	data := make([]float32, xs*ys)

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range ex {
		for _, col := range row {
			data = append(data, float32(col))
		}
	}

	s := NewStarsExtractor(data, xs, ys, 2.5, 255)

	s.Threshold = 100

	stars := s.GetBrightPixels()

	if len(stars) != 24 {
		t.Error("Expected 24 bright pixels, got ", len(stars))
	}
}

func TestNewGetBrightPixelsFrom2DData(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	xs := bounds.Dx()

	ys := bounds.Dy()

	d := utils.Flatten2DUInt32Array(data)

	radius := float32(16.0)

	sigma := float32(8.0)

	s := NewStarsExtractor(d, xs, ys, radius, 65535)

	st := stats.NewStats(d, 65535, xs)

	location, scale := st.FastApproxSigmaClippedMedianAndQn()

	s.Threshold = location + scale*sigma

	stars := s.GetBrightPixels()

	if len(stars) != 2084 {
		t.Error("Expected 2084 bright pixels, got ", len(stars))
	}
}

func TestNewRejectBadPixelsFrom2DData(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	xs := bounds.Dx()

	ys := bounds.Dy()

	d := utils.Flatten2DUInt32Array(data)

	radius := float32(16.0)

	sigma := float32(8.0)

	s := NewStarsExtractor(d, xs, ys, radius, 65535)

	s.Sigma = sigma

	st := stats.NewStats(d, 65535, xs)

	location, scale := st.FastApproxSigmaClippedMedianAndQn()

	s.Threshold = location + scale*sigma

	s.Stars = s.GetBrightPixels()

	stars := s.RejectBadPixels()

	if len(stars) > 2084 {
		t.Error("Expected less than 2084 bright pixels, got ", len(stars))
	}

	if len(stars) < 2030 || len(stars) > 2035 {
		t.Error("Expected to reject about 50 bad pixels, got ", len(stars))
	}
}

func TestNewFilterOverlappingStarsFrom2DData(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	xs := bounds.Dx()

	ys := bounds.Dy()

	d := utils.Flatten2DUInt32Array(data)

	radius := float32(16.0)

	sigma := float32(8.0)

	s := NewStarsExtractor(d, xs, ys, radius, 65535)

	s.Sigma = sigma

	st := stats.NewStats(d, 65535, xs)

	location, scale := st.FastApproxSigmaClippedMedianAndQn()

	s.Threshold = location + scale*sigma

	s.Stars = s.GetBrightPixels()

	s.Stars = s.RejectBadPixels()

	stars := s.FilterOverlappingPixels()

	if len(stars) > 2084 {
		t.Error("Expected less than 2084 bright pixels, got ", len(stars))
	}

	if len(stars) > 2035 {
		t.Error("Expected to filter out a number of overlapping stars", len(stars))
	}

	if len(stars) >= 230 {
		t.Error("Expected to filter out ~1860 number of overlapping stars, but filtered out ", 2084-len(stars))
	}
}

func TestNewShiftToBrightestPixelStarsFrom2DData(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	xs := bounds.Dx()

	ys := bounds.Dy()

	d := utils.Flatten2DUInt32Array(data)

	radius := float32(16.0)

	sigma := float32(8.0)

	s := NewStarsExtractor(d, xs, ys, radius, 65535)

	s.Sigma = sigma

	st := stats.NewStats(d, 65535, xs)

	location, scale := st.FastApproxSigmaClippedMedianAndQn()

	s.Threshold = location + scale*sigma

	s.Stars = s.GetBrightPixels()

	s.Stars = s.RejectBadPixels()

	s.Stars = s.FilterOverlappingPixels()

	stars := s.ShiftToCenterOfMass()

	if len(stars) > 2084 {
		t.Error("Expected less than 2084 bright pixels, got ", len(stars))
	}

	if len(stars) > 2035 {
		t.Error("Expected to filter out a number of overlapping stars", len(stars))
	}

	if len(stars) >= 230 {
		t.Error("Expected to filter out ~1860 number of overlapping stars, but filtered out ", 2084-len(stars))
	}

	if len(stars) != len(s.Stars) {
		t.Error("Expected to shift all stars, but shifted ", len(stars))
	}
}

func TestNewExtractAndFilterHalfFluxRadiusStarsFrom2DData(t *testing.T) {
	data, bounds := GetTestDataFromImage()

	xs := bounds.Dx()

	ys := bounds.Dy()

	d := utils.Flatten2DUInt32Array(data)

	radius := float32(16.0)

	sigma := float32(8.0)

	s := NewStarsExtractor(d, xs, ys, radius, 65535)

	s.Sigma = sigma

	st := stats.NewStats(d, 65535, xs)

	location, scale := st.FastApproxSigmaClippedMedianAndQn()

	s.Threshold = location + scale*sigma

	s.Stars = s.GetBrightPixels()

	s.Stars = s.RejectBadPixels()

	s.Stars = s.FilterOverlappingPixels()

	s.Stars = s.ShiftToCenterOfMass()

	s.Stars = s.FilterOverlappingPixels()

	stars := s.ExtractAndFilterHalfFluxRadius(location, 2.0)

	if len(stars) > 2084 {
		t.Error("Expected less than 2084 bright pixels, got ", len(stars))
	}

	if len(stars) > 2035 {
		t.Error("Expected to filter out a number of overlapping stars", len(stars))
	}

	if len(stars) >= 230 {
		t.Error("Expected to filter out ~1860 number of overlapping stars, but filtered out ", 2084-len(stars))
	}

	if s.HFR == 0 {
		t.Error("Expected to calculate HFR, but got ", s.HFR)
	}

	if s.HFR > 8.0 {
		t.Error("Expected to calculate HFR less than 2.0, but got ", s.HFR)
	}
}
