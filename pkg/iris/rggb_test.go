package iris

import (
	"testing"
)

func TestNewRGGBExposureWidth(t *testing.T) {
	rggb := NewRGGBExposure(ex, 800, 600, "RGGB")

	var got int = rggb.Width

	var want int = 800

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRGGBExposureHeight(t *testing.T) {
	rggb := NewRGGBExposure(ex, 800, 600, "RGGB")

	var got int = rggb.Height

	var want int = 600

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewRGGBGetBayerMatrixOffset(t *testing.T) {
	rggb := NewRGGBExposure(ex, 800, 600, "RGGB")

	xOffset, yOffset, err := rggb.GetBayerMatrixOffset()

	if err != nil {
		t.Errorf("Expected the CFA string to be valid, but got %q", err)
	}

	if xOffset != 0 {
		t.Errorf("got %q, wanted %q", xOffset, 0)
	}

	if yOffset != 0 {
		t.Errorf("got %q, wanted %q", yOffset, 0)
	}
}

func TestNewGRBGGetBayerMatrixOffset(t *testing.T) {
	rggb := NewRGGBExposure(ex, 800, 600, "GRBG")

	xOffset, yOffset, err := rggb.GetBayerMatrixOffset()

	if err != nil {
		t.Errorf("Expected the CFA string to be valid, but got %q", err)
	}

	if xOffset != 1 {
		t.Errorf("got %q, wanted %q", xOffset, 1)
	}

	if yOffset != 0 {
		t.Errorf("got %q, wanted %q", yOffset, 0)
	}
}

func TestNewGBRGGetBayerMatrixOffset(t *testing.T) {
	rggb := NewRGGBExposure(ex, 800, 600, "GBRG")

	xOffset, yOffset, err := rggb.GetBayerMatrixOffset()

	if err != nil {
		t.Errorf("Expected the CFA string to be valid, but got %q", err)
	}

	if xOffset != 0 {
		t.Errorf("got %q, wanted %q", xOffset, 0)
	}

	if yOffset != 1 {
		t.Errorf("got %q, wanted %q", yOffset, 1)
	}
}

func TestNewBGGRGetBayerMatrixOffset(t *testing.T) {
	rggb := NewRGGBExposure(ex, 800, 600, "BGGR")

	xOffset, yOffset, err := rggb.GetBayerMatrixOffset()

	if err != nil {
		t.Errorf("Expected the CFA string to be valid, but got %q", err)
	}

	if xOffset != 1 {
		t.Errorf("got %q, wanted %q", xOffset, 1)
	}

	if yOffset != 1 {
		t.Errorf("got %q, wanted %q", yOffset, 1)
	}
}

func TestNewRGGBGetBayerMatrixOffsetInvalid(t *testing.T) {
	rggb := NewRGGBExposure(ex, 800, 600, "INVALID")

	_, _, err := rggb.GetBayerMatrixOffset()

	if err == nil {
		t.Errorf("Expected the CFA string to be invalid, but got %q", err)
	}
}
