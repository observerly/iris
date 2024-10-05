/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package image

/*****************************************************************************************************************/

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/observerly/iris/pkg/utils"
)

/*****************************************************************************************************************/

func NewGray16FromRawFloat32Pixels(pixels []float32, width int) (*image.Gray16, error) {
	// Check that the number of pixels is a multiple of the width:
	if len(pixels)%width != 0 {
		return nil, fmt.Errorf("the number of pixels must be a multiple of the width")
	}

	// Calculate the height of the image:
	height := len(pixels) / width

	// Create a new image:
	img := image.NewGray16(image.Rect(0, 0, width, height))

	min, max := utils.BoundsFloat32Array(pixels)

	b := max - min

	if b == 0 {
		b = 1
	}

	// Set the pixels whilst normalizing them:
	for i := 0; i < len(pixels); i++ {
		x := i % width
		y := i / width
		// Normalize the value to 0-65535
		v := uint16(math.Round(float64(((pixels[i] - min) / b) * 65535)))
		// Set the pixel value at the correct position:
		img.SetGray16(x, y, color.Gray16{Y: v})
	}

	// Return the image:
	return img, nil
}

/*****************************************************************************************************************/
