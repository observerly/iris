package frames

import (
	"testing"

	"github.com/observerly/iris/pkg/fits"
)

func TestNewMasterDarkFrame(t *testing.T) {
	var dark = [][]uint32{
		{1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	var darkFrame = fits.NewFITSImageFrom2DData(dark, 2, 16, 16, 255)

	var frames = []fits.FITSImage{
		*darkFrame,
		*darkFrame,
		*darkFrame,
		*darkFrame,
		*darkFrame,
	}

	masterDark, err := NewMasterDarkFrame(frames, 2, 16, 16, 255, 130)

	if err != nil {
		t.Errorf("NewMasterDarkFrame() failed: %s", err)
	}

	if masterDark.Count != 5 {
		t.Errorf("NewMasterDarkFrame() failed: expected count of 5, got %d", masterDark.Count)
	}

	if masterDark.Pixels != 256 {
		t.Errorf("NewMasterDarkFrame() failed: expected pixels of 256, got %d", masterDark.Pixels)
	}

	if masterDark.Combined.ADU != 255 {
		t.Errorf("NewMasterDarkFrame() failed: expected ADU of 255, got %d", masterDark.Combined.ADU)
	}

	if masterDark.Combined.Data[1] != 6 {
		t.Errorf("NewMasterDarkFrame() failed: expected data[0] of 6, got %f", masterDark.Combined.Data[0])
	}

	if masterDark.Combined.Exposure != 130 {
		t.Errorf("NewMasterDarkFrame() failed: expected exposure of 130, got %f", masterDark.Combined.Exposure)
	}
}
