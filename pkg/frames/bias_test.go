package frames

import (
	"testing"

	"github.com/observerly/iris/pkg/fits"
)

func TestNewMasterBiasFrame(t *testing.T) {
	var bias = [][]uint32{
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

	var biasFrame = fits.NewFITSImageFrom2DData(bias, 2, 16, 16, 255)

	var frames = []fits.FITSImage{
		*biasFrame,
		*biasFrame,
		*biasFrame,
		*biasFrame,
		*biasFrame,
	}

	masterBias, err := NewMasterBiasFrame(frames, 2, 16, 16, 255, 0.05)

	if err != nil {
		t.Errorf("NewMasterBiasFrame() failed: %s", err)
	}

	if masterBias.Count != 5 {
		t.Errorf("NewMasterBiasFrame() failed: expected count of 5, got %d", masterBias.Count)
	}

	if masterBias.Pixels != 256 {
		t.Errorf("NewMasterBiasFrame() failed: expected pixels of 256, got %d", masterBias.Pixels)
	}

	if masterBias.Combined.ADU != 255 {
		t.Errorf("NewMasterBiasFrame() failed: expected ADU of 255, got %d", masterBias.Combined.ADU)
	}

	if masterBias.Combined.Data[1] != 6 {
		t.Errorf("NewMasterBiasFrame() failed: expected data[0] of 6, got %f", masterBias.Combined.Data[0])
	}
}

func TestApplyFrameToMasterBiasFrame(t *testing.T) {
	var bias = [][]uint32{
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

	var biasFrame = fits.NewFITSImageFrom2DData(bias, 2, 16, 16, 255)

	var frames = []fits.FITSImage{
		*biasFrame,
		*biasFrame,
		*biasFrame,
		*biasFrame,
		*biasFrame,
	}

	masterBias, err := NewMasterBiasFrame(frames, 2, 16, 16, 255, 0.05)

	if err != nil {
		t.Errorf("NewMasterBiasFrame() failed: %s", err)
	}

	if masterBias.Count != 5 {
		t.Errorf("NewMasterBiasFrame() failed: expected count of 5, got %d", masterBias.Count)
	}

	if masterBias.Pixels != 256 {
		t.Errorf("NewMasterBiasFrame() failed: expected pixels of 256, got %d", masterBias.Pixels)
	}

	if masterBias.Combined.ADU != 255 {
		t.Errorf("NewMasterBiasFrame() failed: expected ADU of 255, got %d", masterBias.Combined.ADU)
	}

	if masterBias.Combined.Data[1] != 6 {
		t.Errorf("NewMasterBiasFrame() failed: expected data[0] of 6, got %f", masterBias.Combined.Data[0])
	}

	bias = [][]uint32{
		{1, 2, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
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

	biasFrame = fits.NewFITSImageFrom2DData(bias, 2, 16, 16, 255)

	masterBias, err = masterBias.ApplyFrame(biasFrame)

	if err != nil {
		t.Errorf("NewMasterBiasFrame() failed: %s", err)
	}

	if masterBias.Count != 6 {
		t.Errorf("NewMasterBiasFrame() failed: expected count of 5, got %d", masterBias.Count)
	}

	if masterBias.Pixels != 256 {
		t.Errorf("NewMasterBiasFrame() failed: expected pixels of 256, got %d", masterBias.Pixels)
	}

	if masterBias.Combined.ADU != 255 {
		t.Errorf("NewMasterBiasFrame() failed: expected ADU of 255, got %d", masterBias.Combined.ADU)
	}

	if masterBias.Combined.Data[1] != 4 {
		t.Errorf("NewMasterBiasFrame() failed: expected data[0] of 6, got %f", masterBias.Combined.Data[0])
	}
}
