package frames

import "github.com/observerly/iris/pkg/fits"

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
