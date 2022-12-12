package frames

import (
	"time"

	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/utils"
)

type MasterDarkFrame struct {
	Type             string           // The type of master frame (e.g., bias, dark, flat)
	Count            int              // The number of frames used to create the master frame
	Pixels           int32            // The number of pixels in the master frame
	Frames           []fits.FITSImage // The individual frames used to create the master frame
	Combined         *fits.FITSImage  // The combined master frame
	MasterBias       *MasterFrame     // The master bias frame used to create the master dark frame
	CreatedTimestamp int64
}

/*
NewMasterDarkFrame()

Creates a new master dark frame from a slice of dark frames.

The idea of a dark frame is to take a series of exposures with the shutter closed,
for the shortest exposure resolution supported by the camera with no light falling
on the sensor, with the same sensor temperature, exposure duration and binning
level as the associated light frame. The resulting images are then averaged to
produce a master dark frame.

The master dark frame is then created by taking the mean of all the dark frames.

@retuns a new FITSImage containing the master dark frame.
@see Image Calibration & Stack Woodhouse, C. (2017). The Astrophotography Manual. Taylor & Francis. p.203
*/
func NewMasterDarkFrame(frames []fits.FITSImage, masterBias *MasterFrame, naxis int32, naxis1 int32, naxis2 int32, adu int32, exposureTime float32) (*MasterDarkFrame, error) {
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

		f.Data, err = utils.SubtractFloat32Array(combined, masterBias.Combined.Data)

		if err != nil {
			return nil, err
		}
	} else {
		var err error = nil

		combined := frames[0].Data

		f.Data = combined

		f.Data, err = utils.SubtractFloat32Array(combined, masterBias.Combined.Data)

		if err != nil {
			return nil, err
		}
	}

	f.Exposure = exposureTime

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
		Comment: "The exposure time (s) of the dark frame",
	}

	f.Header.Strings["SENSOR"] = struct {
		Value   string
		Comment string
	}{
		Value:   "Monochrome",
		Comment: "ASCOM Alpaca Sensor Type",
	}

	return &MasterDarkFrame{
		Type:             "dark",
		Count:            len(frames),
		Pixels:           pixels,
		Frames:           frames,
		Combined:         f,
		MasterBias:       masterBias,
		CreatedTimestamp: time.Now().Unix(),
	}, nil
}

func (m *MasterDarkFrame) ApplyDarkFrame(frame *fits.FITSImage) (*MasterDarkFrame, error) {
	// Add the current combined master frame to the master bias before averaging:
	combined, err := utils.AddFloat32Array(m.Combined.Data, m.MasterBias.Combined.Data)

	if err != nil {
		return nil, err
	}

	// Take the current combined master frame and apply the new frame to it:
	combined, err = utils.MeanFloat32Arrays([][]float32{combined, frame.Data})

	if err != nil {
		return nil, err
	}

	// Subtract the master bias from the combined master frame:
	combined, err = utils.SubtractFloat32Array(combined, m.MasterBias.Combined.Data)

	if err != nil {
		return nil, err
	}

	// Create a new FITSImage from the master data
	m.Combined.Data = combined

	return &MasterDarkFrame{
		Type:             m.Type,
		Count:            m.Count + 1,
		Pixels:           m.Pixels,
		Frames:           append(m.Frames, *frame),
		Combined:         m.Combined,
		MasterBias:       m.MasterBias,
		CreatedTimestamp: m.CreatedTimestamp,
	}, nil
}
