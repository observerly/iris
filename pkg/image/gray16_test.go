/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package image

/*****************************************************************************************************************/

import (
	"image"
	"testing"
)

/*****************************************************************************************************************/

func TestNewGray16FromRawFloat32Pixels(t *testing.T) {
	// Create some test data:
	pixels := []float32{
		0.0, 0.5, 1.0,
		0.5, 1.0, 0.5,
		1.0, 0.5, 0.0,
	}

	// Create a new image:
	img, err := NewGray16FromRawFloat32Pixels(pixels, 3)

	// Check that the image was created successfully:
	if err != nil {
		t.Errorf("error creating image: %v", err)
	}

	// Check that the image has the correct bounds:
	if img.Bounds() != image.Rect(0, 0, 3, 3) {
		t.Errorf("incorrect image bounds: %v", img.Bounds())
	}
}

/*****************************************************************************************************************/
