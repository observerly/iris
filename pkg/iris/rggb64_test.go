package iris

import "testing"

func TestNewRGGB64ExposureWidth(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "RGGB")

	var got int = rggb.Width

	var want int = 800

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRGGB64ExposureHeight(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "RGGB")

	var got int = rggb.Height

	var want int = 600

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRGGB64ExpsourePixels(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "RGGB")

	var got int = rggb.Pixels

	var want int = 480000

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
