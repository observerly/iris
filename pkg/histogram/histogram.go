/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/histogram
//	@license	Copyright © 2021-2025 observerly

/*****************************************************************************************************************/

package histogram

/*****************************************************************************************************************/

import (
	"image"

	"github.com/observerly/iris/pkg/utils"
)

/*****************************************************************************************************************/

const hsize8 = 256

/*****************************************************************************************************************/

const hsize16 = 65535

/*****************************************************************************************************************/

// Computes the histogram for a grayscale image, and bins the pixel values according to their accumulated count.
//
// Returns an array of 256 uint64 values containing a distribution of the pixel values.

func HistogramGray(img *image.Gray) [hsize8]uint64 {
	bounds := img.Bounds()

	size := bounds.Size()

	var res [hsize8]uint64

	utils.DeferForEachPixel(size, func(x, y int) {
		pixel := img.GrayAt(x, y)
		res[pixel.Y]++
	})

	return res
}

/*****************************************************************************************************************/

/*
HistogramGray16

Computes the histogram for a 16 bit grayscale image, and bins the pixel values according to their accumulated count.

Returns an array of 256 uint64 values containing a distribution of the pixel values.
*/
func HistogramGray16(img *image.Gray16) [hsize16]uint64 {
	bounds := img.Bounds()

	size := bounds.Size()

	var res [hsize16]uint64

	utils.DeferForEachPixel(size, func(x, y int) {
		pixel := img.Gray16At(x, y)
		res[pixel.Y]++
	})

	return res
}
