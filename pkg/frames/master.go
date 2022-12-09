package frames

import (
	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/utils"
)

type MasterFrame struct {
	Type             string           // The type of master frame (e.g., bias, dark, flat)
	Count            int              // The number of frames used to create the master frame
	Pixels           int32            // The number of pixels in the master frame
	Frames           []fits.FITSImage // The individual frames used to create the master frame
	Combined         *fits.FITSImage  // The combined master frame
	CreatedTimestamp int64
}

func (m *MasterFrame) ApplyFrame(frame *fits.FITSImage) (*MasterFrame, error) {
	// Take the current combined master frame and apply the new frame to it:
	combined, err := utils.MeanFloat32Arrays([][]float32{m.Combined.Data, frame.Data})

	if err != nil {
		return nil, err
	}

	// Create a new FITSImage from the master data
	m.Combined.Data = combined

	return &MasterFrame{
		Type:             m.Type,
		Count:            m.Count + 1,
		Pixels:           m.Pixels,
		Frames:           append(m.Frames, *frame),
		Combined:         m.Combined,
		CreatedTimestamp: m.CreatedTimestamp,
	}, nil
}
