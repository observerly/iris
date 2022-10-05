package iris

import (
	"testing"
)

func TestNewRGGBExposureWidth(t *testing.T) {
	mono := NewRGGBExposure(ex, 800, 600)

	var got int = mono.Width

	var want int = 800

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRGGBExposureHeight(t *testing.T) {
	mono := NewRGGBExposure(ex, 800, 600)

	var got int = mono.Height

	var want int = 600

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
