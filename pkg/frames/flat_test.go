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

	var bias = [][]uint32{
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
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

	var flatFrame = fits.NewFITSImageFrom2DData(flat, 2, 16, 16, 255)

	var biasFrame = fits.NewFITSImageFrom2DData(bias, 2, 16, 16, 255)

	var frames = []fits.FITSImage{
		*flatFrame,
		*flatFrame,
		*flatFrame,
		*flatFrame,
		*flatFrame,
	}

	masterBias, err := NewMasterBiasFrame([]fits.FITSImage{*biasFrame}, 2, 16, 16, 255, 130)

	if err != nil {
		t.Errorf("NewMasterBiasFrame() failed: %s", err)
	}

	masterFlat, err := NewMasterFlatFrame(frames, masterBias, 2, 16, 16, 255, 130)

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

	if masterFlat.Combined.Data[1] != 253 {
		t.Errorf("NewMasterFlatFrame() failed: expected data[0] of 253, got %f", masterFlat.Combined.Data[1])
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

	var bias = [][]uint32{
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
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

	var flatFrame = fits.NewFITSImageFrom2DData(flat, 2, 16, 16, 255)

	var biasFrame = fits.NewFITSImageFrom2DData(bias, 2, 16, 16, 255)

	var frames = []fits.FITSImage{
		*flatFrame,
		*flatFrame,
		*flatFrame,
		*flatFrame,
		*flatFrame,
	}

	masterBias, err := NewMasterBiasFrame([]fits.FITSImage{*biasFrame}, 2, 16, 16, 255, 130)

	if err != nil {
		t.Errorf("NewMasterBiasFrame() failed: %s", err)
	}

	masterFlat, err := NewMasterFlatFrame(frames, masterBias, 2, 16, 16, 255, 130)

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

	if masterFlat.Combined.Data[1] != 253 {
		t.Errorf("NewMasterFlatFrame() failed: expected data[0] of 253, got %f", masterFlat.Combined.Data[1])
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

	masterFlat, err = masterFlat.ApplyFlatFrame(flatFrame)

	if err != nil {
		t.Errorf("NewMasterFlatFrame() failed: %s", err)
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

	if masterFlat.Combined.Data[1] != 251 {
		t.Errorf("NewmasterFlatFrame() failed: expected data[0] of 251, got %f", masterFlat.Combined.Data[1])
	}
}
