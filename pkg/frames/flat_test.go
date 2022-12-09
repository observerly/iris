package frames

import (
	"testing"

	"github.com/observerly/iris/pkg/fits"
)

func TestNewMasterFlatFrame(t *testing.T) {
	var flat = [][]uint32{
		{255, 254, 255, 250, 255, 255, 255, 255, 255, 255, 255, 252, 253, 254, 252, 255},
		{251, 250, 255, 255, 253, 252, 250, 255, 254, 250, 251, 251, 255, 255, 250, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 250, 251, 255},
		{250, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 253, 255, 255},
		{255, 253, 252, 250, 255, 254, 250, 251, 251, 255, 255, 255, 254, 255, 255, 255},
		{255, 252, 250, 255, 254, 250, 251, 251, 255, 255, 255, 255, 255, 254, 254, 255},
		{255, 250, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
	}

	var flatFrame = fits.NewFITSImageFrom2DData(flat, 2, 16, 16, 255)

	var frames = []fits.FITSImage{
		*flatFrame,
		*flatFrame,
		*flatFrame,
		*flatFrame,
		*flatFrame,
	}

	masterFlat, err := NewMasterFlatFrame(frames, 2, 16, 16, 255, 130)

	if err != nil {
		t.Errorf("NewMasterDarkFrame() failed: %s", err)
	}

	if masterFlat.Count != 5 {
		t.Errorf("NewMasterFlatFrame() failed: expected count of 5, got %d", masterFlat.Count)
	}

	if masterFlat.Pixels != 256 {
		t.Errorf("NewMasterFlatFrame() failed: expected pixels of 256, got %d", masterFlat.Pixels)
	}

	if masterFlat.Combined.ADU != 255 {
		t.Errorf("NewMasterFlatFrame() failed: expected ADU of 255, got %d", masterFlat.Combined.ADU)
	}

	if masterFlat.Combined.Data[1] != 254 {
		t.Errorf("NewMasterFlatFrame() failed: expected data[0] of 254, got %f", masterFlat.Combined.Data[0])
	}

	if masterFlat.Combined.Exposure != 130 {
		t.Errorf("NewMasterFlatFrame() failed: expected exposure of 130, got %f", masterFlat.Combined.Exposure)
	}
}

func TestApplyFrameToNewMasterFlatFrame(t *testing.T) {
	var flat = [][]uint32{
		{255, 254, 255, 250, 255, 255, 255, 255, 255, 255, 255, 252, 253, 254, 252, 255},
		{251, 250, 255, 255, 253, 252, 250, 255, 254, 250, 251, 251, 255, 255, 250, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 250, 251, 255},
		{250, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 253, 255, 255},
		{255, 253, 252, 250, 255, 254, 250, 251, 251, 255, 255, 255, 254, 255, 255, 255},
		{255, 252, 250, 255, 254, 250, 251, 251, 255, 255, 255, 255, 255, 254, 254, 255},
		{255, 250, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
	}

	var flatFrame = fits.NewFITSImageFrom2DData(flat, 2, 16, 16, 255)

	var frames = []fits.FITSImage{
		*flatFrame,
		*flatFrame,
		*flatFrame,
		*flatFrame,
		*flatFrame,
	}

	masterFlat, err := NewMasterFlatFrame(frames, 2, 16, 16, 255, 130)

	if err != nil {
		t.Errorf("NewMasterDarkFrame() failed: %s", err)
	}

	if masterFlat.Count != 5 {
		t.Errorf("NewMasterFlatFrame() failed: expected count of 5, got %d", masterFlat.Count)
	}

	if masterFlat.Pixels != 256 {
		t.Errorf("NewMasterFlatFrame() failed: expected pixels of 256, got %d", masterFlat.Pixels)
	}

	if masterFlat.Combined.ADU != 255 {
		t.Errorf("NewMasterFlatFrame() failed: expected ADU of 255, got %d", masterFlat.Combined.ADU)
	}

	if masterFlat.Combined.Data[1] != 254 {
		t.Errorf("NewMasterFlatFrame() failed: expected data[0] of 254, got %f", masterFlat.Combined.Data[0])
	}

	if masterFlat.Combined.Exposure != 130 {
		t.Errorf("NewMasterFlatFrame() failed: expected exposure of 130, got %f", masterFlat.Combined.Exposure)
	}

	flat = [][]uint32{
		{255, 250, 255, 250, 255, 255, 255, 255, 255, 255, 255, 252, 253, 254, 252, 255},
		{251, 250, 255, 255, 253, 252, 250, 255, 254, 250, 251, 251, 255, 255, 250, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 250, 251, 255},
		{250, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 253, 255, 255},
		{255, 253, 252, 250, 255, 254, 250, 251, 251, 255, 255, 255, 254, 255, 255, 255},
		{255, 252, 250, 255, 254, 250, 251, 251, 255, 255, 255, 255, 255, 254, 254, 255},
		{255, 250, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
		{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
	}

	flatFrame = fits.NewFITSImageFrom2DData(flat, 2, 16, 16, 255)

	masterFlat, err = masterFlat.ApplyFrame(flatFrame)

	if err != nil {
		t.Errorf("NewMasterBiasFrame() failed: %s", err)
	}

	if masterFlat.Count != 6 {
		t.Errorf("NewmasterFlatFrame() failed: expected count of 5, got %d", masterFlat.Count)
	}

	if masterFlat.Pixels != 256 {
		t.Errorf("NewmasterFlatFrame() failed: expected pixels of 256, got %d", masterFlat.Pixels)
	}

	if masterFlat.Combined.ADU != 255 {
		t.Errorf("NewmasterFlatFrame() failed: expected ADU of 255, got %d", masterFlat.Combined.ADU)
	}

	if masterFlat.Combined.Data[1] != 252 {
		t.Errorf("NewmasterFlatFrame() failed: expected data[0] of 6, got %f", masterFlat.Combined.Data[0])
	}
}
