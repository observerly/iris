package histogram

import (
	"image"

	"github.com/observerly/iris/pkg/utils"
)

const hsize = 256

/*
  HistogramGray

  Computes the histogram for a grayscale image, and bins the pixel values according to their accumulated count.

  Returns an array of 256 uint64 values containing a distribution of the pixel values.
*/
func HistogramGray(img *image.Gray) [hsize]uint64 {
	bounds := img.Bounds()

	size := bounds.Size()

	var res [hsize]uint64

	utils.DeferForEachPixel(size, func(x, y int) {
		pixel := img.GrayAt(x, y)
		res[pixel.Y]++
	})

	return res
}
