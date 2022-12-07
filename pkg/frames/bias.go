package frames

import (
	"github.com/observerly/iris/pkg/fits"
)

type MasterBiasFrame struct {
	Count            int              // The number of bias frames used to create the master bias frame
	Pixels           int32            // The number of pixels in the master bias frame
	Frames           []fits.FITSImage // The individual bias frames used to create the master bias frame
	Combined         *fits.FITSImage  // The combined master bias frame
	CreatedTimestamp int64
}
