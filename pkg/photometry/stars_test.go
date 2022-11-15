package photometry

import (
	"testing"
)

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

	s := NewStarsExtractor(data, 16, 16, 2.5)

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

	s := NewStarsExtractor(data, 16, 16, 2.5)

	s.Threshold = 100

	stars := s.GetBrightPixels()

	if len(stars) != 24 {
		t.Error("Expected 24 bright pixels, got ", len(stars))
	}
}
