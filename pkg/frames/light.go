package frames

import (
	"time"

	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/utils"
)

type CalibratedLightFrame struct {
	Type             string           // The type of master frame (e.g., bias, dark, flat)
	Count            int              // The number of frames used to create the master frame
	Pixels           int32            // The number of pixels in the master frame
	Frames           []fits.FITSImage // The individual frames used to create the master frame
	Combined         *fits.FITSImage  // The calibrated master light frame
	MaterBias        *MasterFrame     // The master bias frame used to create the master flat frame
	MasterFlat       *MasterFlatFrame // The master flat frame used to create the master light frame
	MasterDark       *MasterDarkFrame // The master dark frame used to create the master light frame
	CreatedTimestamp int64
}

/*
NewCalibratedLightFrame()

Creates a new calibrated light frame from a light frame, master bias,
master dark, and master flat.

A calibrated light frame is a light frame that has been calibrated by
subtracting the master bias, master dark, and then multiplying by the
averaged master flat divided by the master flat.
*/
func NewCalibratedLightFrame(
	frame *fits.FITSImage,
	masterBias *MasterFrame,
	masterDark *MasterDarkFrame,
	masterFlat *MasterFlatFrame,
	naxis int32,
	naxis1 int32,
	naxis2 int32,
	adu int32,
	exposureTime float32,
) (*CalibratedLightFrame, error) {
	pixels := naxis1 * naxis2

	// Subtract the master bias from the light frame:
	light, err := utils.SubtractFloat32Array(frame.Data, masterBias.Combined.Data)

	if err != nil {
		return nil, err
	}

	// Subtract the master dark from the light frame:
	light, err = utils.SubtractFloat32Array(light, masterDark.Combined.Data)

	if err != nil {
		return nil, err
	}

	// Obtain the average master flat:
	averageMasterFlat, err := utils.AverageFloat32Array(masterFlat.Combined.Data)

	if err != nil {
		return nil, err
	}

	// Divide every pixel in the master flat by the average master flat:
	light, err = utils.DivideFloat32Array(light, masterFlat.Combined.Data, averageMasterFlat)

	if err != nil {
		return nil, err
	}

	// Create a new FITSImage from the master bias data
	f := fits.NewFITSImage(
		naxis,
		naxis1,
		naxis2,
		adu,
	)

	f.Data = light

	f.Pixels = pixels

	f.Header.Ints["ADU"] = struct {
		Value   int32
		Comment string
	}{
		Value:   adu,
		Comment: "Analog to Digital Units (ADU)",
	}

	f.Header.Floats["EXPOSURE"] = struct {
		Value   float32
		Comment string
	}{
		Value:   exposureTime,
		Comment: "The exposure time (s) of the flat frame",
	}

	f.Header.Strings["SENSOR"] = struct {
		Value   string
		Comment string
	}{
		Value:   "Monochrome",
		Comment: "ASCOM Alpaca Sensor Type",
	}

	// Create the new calibrated light frame:
	return &CalibratedLightFrame{
		Type:             "light",
		Count:            1,
		Pixels:           frame.Pixels,
		Frames:           []fits.FITSImage{*frame},
		Combined:         f,
		MaterBias:        masterBias,
		MasterFlat:       masterFlat,
		MasterDark:       masterDark,
		CreatedTimestamp: time.Now().Unix(),
	}, nil
}
