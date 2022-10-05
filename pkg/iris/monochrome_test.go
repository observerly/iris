package iris

import "testing"

var ex = [][]uint32{}

func TestNewMonochromeExposureWidth(t *testing.T) {
	mono := NewMonochromeExposure(ex, 800, 600)

	var got int = mono.Width

	var want int = 800

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochromeExposureHeight(t *testing.T) {
	mono := NewMonochromeExposure(ex, 800, 600)

	var got int = mono.Height

	var want int = 600

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
