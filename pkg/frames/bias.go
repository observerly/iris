/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/frames
//	@license	Copyright © 2021-2025 observerly

/*****************************************************************************************************************/

package frames

/*****************************************************************************************************************/

import (
	"time"

	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/utils"
)

/*****************************************************************************************************************/

// Creates a new master bias frame from a slice of bias frames.
// The idea of a bias frame is to take a series of exposures with the shutter closed, for the shortest exposure
// resolution supported by the camera with no light falling on the sensor. The resulting images are then averaged
// to produce a master bias frame, which is then subtracted from all subsequent images to remove the bias noise.
//
// The master bias frame is then created by taking the mean of all the bias frames.
//
// @see Image Calibration & Stack Woodhouse, C. (2017). The Astrophotography Manual. Taylor & Francis. p.203
func NewMasterBiasFrame(frames []fits.FITSImage, naxis int32, naxis1 int32, naxis2 int32, adu int32, resolution float32) (*MasterFrame, error) {
	pixels := naxis1 * naxis2

	// Create a slice of 2D data arrays from the slice of FITSImages
	data := make([][]float32, len(frames))

	for i, frame := range frames {
		data[i] = frame.Data
	}

	// Create a new FITSImage from the master bias data
	f := fits.NewFITSImage(
		naxis,
		naxis1,
		naxis2,
		adu,
	)

	// Combine the data arrays into a single array, by taking
	// the mean of the total of all the frames for each pixel:
	if len(frames) > 1 {
		combined, err := utils.MeanFloat32Arrays(data)

		if err != nil {
			return nil, err
		}

		f.Data = combined
	} else {
		combined := frames[0].Data

		f.Data = combined
	}

	f.Exposure = resolution

	f.Pixels = pixels

	f.Header.Ints["ADU"] = struct {
		Value   int32
		Comment string
	}{
		Value:   adu,
		Comment: "Analog to Digital Units (ADU)",
	}

	f.Header.Floats["RESOLUTION"] = struct {
		Value   float32
		Comment string
	}{
		Value:   resolution,
		Comment: "Smallest increment in exposure time (s)",
	}

	f.Header.Strings["SENSOR"] = struct {
		Value   string
		Comment string
	}{
		Value:   "Monochrome",
		Comment: "ASCOM Alpaca Sensor Type",
	}

	return &MasterFrame{
		Type:             "bias",
		Count:            len(frames),
		Pixels:           pixels,
		Frames:           frames,
		Combined:         f,
		CreatedTimestamp: time.Now().Unix(),
	}, nil
}

/*****************************************************************************************************************/
