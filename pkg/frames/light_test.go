package frames

import (
	"testing"

	"github.com/observerly/iris/pkg/fits"
)

func TestNewCalibratedLightFrame(t *testing.T) {
	// Setup the Master Bias Frame:
	var bias = [][]uint32{
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
	}

	biasFrame := fits.NewFITSImageFrom2DData(bias, 2, 16, 16, 255)

	masterBias, err := NewMasterBiasFrame([]fits.FITSImage{*biasFrame}, 2, 16, 16, 255, 1)

	if err != nil {
		t.Errorf("NewMasterBiasFrame() failed: %s", err)
	}

	// Setup the Master Dark Frame:

	// Give me a random dark frame:
	var dark = [][]uint32{
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
		{3, 2, 1, 6, 1, 1, 5, 6, 6, 6, 6, 7, 8, 1, 3, 1},
	}

	darkFrame := fits.NewFITSImageFrom2DData(dark, 2, 16, 16, 255)

	masterDark, err := NewMasterDarkFrame([]fits.FITSImage{*darkFrame}, masterBias, 2, 16, 16, 255, 130)

	if err != nil {
		t.Errorf("NewMasterDarkFrame() failed: %s", err)
	}

	// Setup the Master Flat Frame:
	var flat = [][]uint32{
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
		{255, 232, 251, 252, 251, 252, 251, 252, 252, 253, 255, 251, 255, 250, 244, 252},
	}

	flatFrame := fits.NewFITSImageFrom2DData(flat, 2, 16, 16, 255)

	masterFlat, err := NewMasterFlatFrame([]fits.FITSImage{*flatFrame}, masterBias, 2, 16, 16, 255, 130)

	if err != nil {
		t.Errorf("NewMasterFlatFrame() failed: %s", err)
	}

	// Setup the Light Frame:
	var light = [][]uint32{
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
		{190, 187, 188, 183, 178, 169, 190, 191, 192, 193, 194, 195, 196, 197, 198, 199},
	}

	lightFrame := fits.NewFITSImageFrom2DData(light, 2, 16, 16, 255)

	// Create the Calibrated Light Frame:
	calibratedLightFrame, err := NewCalibratedLightFrame(lightFrame, masterBias, masterDark, masterFlat, 2, 16, 16, 255, 130)

	if err != nil {
		t.Errorf("NewCalibratedLightFrame() failed: %s", err)
	}

	// Check the calibrated light frame:
	if calibratedLightFrame.Type != "light" {
		t.Errorf("NewCalibratedLightFrame() failed: Type is not 'light'")
	}

	if calibratedLightFrame.Count != 1 {
		t.Errorf("NewCalibratedLightFrame() failed: Count is not 1")
	}

	if calibratedLightFrame.Pixels != 256 {
		t.Errorf("NewCalibratedLightFrame() failed: Pixels is not 256")
	}

	if calibratedLightFrame.Combined.Data[0] != 180.466049 {
		t.Errorf("NewCalibratedLightFrame() failed: Combined.Data[0] expected to be 180.466049, but got %f", calibratedLightFrame.Combined.Data[0])
	}

	if calibratedLightFrame.Combined.Data[1] != 196.312225 {
		t.Errorf("NewCalibratedLightFrame() failed: Combined.Data[255] expected to be 196.312225, but got %f", calibratedLightFrame.Combined.Data[1])
	}

	if calibratedLightFrame.Combined.Data[2] != 187.09541 {
		t.Errorf("NewCalibratedLightFrame() failed: Combined.Data[255] expected to be 187.09541, but got %f", calibratedLightFrame.Combined.Data[2])
	}

	for _, v := range calibratedLightFrame.Combined.Data {
		if v < 0 {
			t.Errorf("NewCalibratedLightFrame() failed: Combined.Data has a value less than 0, but got %f", v)
		}

		if v > 255 {
			t.Errorf("NewCalibratedLightFrame() failed: Combined.Data has a value greater than 255, but got %f", v)
		}
	}
}
