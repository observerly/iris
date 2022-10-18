package iris

import (
	"image/jpeg"
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

func TestNewRGGBExpsourePixels(t *testing.T) {
	rggb := NewRGGBExposure(ex, 800, 600, "RGGB")

	var got int = rggb.Pixels

	var want int = 480000

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

func TestNewRGGBDebayerBilinearInterpolation(t *testing.T) {
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

	rggb := NewRGGBExposure(ex, 16, 16, "RGGB")

	xOffset, yOffset, err := rggb.GetBayerMatrixOffset()

	if err != nil {
		t.Errorf("Expected the CFA string to be valid, but got %q", err)
	}

	err = rggb.DebayerBilinearInterpolation(xOffset, yOffset)

	if err != nil {
		t.Errorf("Expected the debayering to be successful, but got %q", err)
	}

	// Encode the image as a JPEG:
	err = jpeg.Encode(&rggb.Buffer, rggb.Image, &jpeg.Options{Quality: 100})

	if err != nil {
		t.Errorf("Expected the JPEG encoding to be successful, but got %q", err)
	}

	if err != nil {
		t.Errorf("Expected to be able to preprocess the RGGB CFA image, but got %q", err)
	}
}

func TestNewRGGBPreprocess(t *testing.T) {
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

	rggb := NewRGGBExposure(ex, 16, 16, "RGGB")

	_, err := rggb.Preprocess()

	if err != nil {
		t.Errorf("Expected to be able to preprocess the RGGB CFA image, but got %q", err)
	}
}
