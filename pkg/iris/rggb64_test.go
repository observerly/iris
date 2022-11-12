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

func TestNewRGGB64GetBayerMatrixOffset(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "RGGB")

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

func TestNewGRBG64GetBayerMatrixOffset(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "GRBG")

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

func TestNewGBRG64GetBayerMatrixOffset(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "GBRG")

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

func TestNewBGGR64GetBayerMatrixOffset(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "BGGR")

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

func TestNewRGGB64GetBayerMatrixOffsetInvalid(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "INVALID")

	_, _, err := rggb.GetBayerMatrixOffset()

	if err == nil {
		t.Errorf("Expected the CFA string to be invalid, but got %q", err)
	}
}

func TestNewRGGB64ExposureGetBuffer(t *testing.T) {
	rggb := NewRGGB64Exposure(ex, 1, 800, 600, "RGGB")

	_, err := rggb.GetBuffer(rggb.Image)

	if err != nil {
		t.Errorf("Expected no error when creating the output buffer, got %q", err)
	}
}
