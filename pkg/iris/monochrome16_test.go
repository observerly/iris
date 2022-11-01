package iris

import (
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
