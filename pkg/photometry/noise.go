package photometry

import "math"

type NoiseExtractor struct {
	Width  int
	Height int
	Noise  float64
	Raw    []uint32
}

func NewNoiseExtractor(exposure [][]uint32, xs int, ys int) *NoiseExtractor {
	// Locate the brightest pixels in the data array above the threshold and return them as stars:
	var raw []uint32

	// Flatten the 2D Exposure Array array into a 1D array:
	for _, a := range exposure {
		raw = append(raw, a...)
	}

	return &NoiseExtractor{
		Width:  xs,
		Height: ys,
		Noise:  0,
		Raw:    raw,
	}
}

/*
  GetGaussianNoise()

  From J. Immerkær, “Fast Noise Variance Estimation”, Computer Vision and Image Understanding, Vol. 64, No. 2, pp. 300-302, Sep. 1996.
*/
func (n *NoiseExtractor) GetGaussianNoise() float64 {
	// Weights for the 3x3 noise estimate kernel:
	weight := []int32{
		1, -2, 1,
		-2, 4, -2,
		1, -2, 1,
	}

	xs := int32(n.Width)

	ys := int32(len(n.Raw)) / xs

	// Offsets for the 3x3 kernel:
	offset := []int32{
		-xs - 1, -xs, -xs + 1,
		-1, 0, 1,
		xs - 1, xs, xs + 1,
	}

	// The accumulated noise throughout the image:
	noise := 0.0

	for y := int32(1); y < ys; y++ {
		acc := 0.0
		for x := int32(1); x < xs; x++ {
			// Get pixel offset value:
			i := y*xs + x

			// Convolved value at pixel:
			conv := float64(0)

			// Convolve the pixel with the weight matrix:
			for j, o := range offset {
				if i+o >= 0 && i+o < int32(len(n.Raw)) {
					conv += float64(weight[j]) * float64(n.Raw[i+o])
				}
			}

			// Accumulate the convolved values
			acc += math.Abs(conv)
		}

		// Accumulate the noise
		noise += acc
	}

	// Calculate the fraction of the noise to apply to the image using a Gaussian distribution:
	fr := math.Sqrt(0.5*math.Pi) / (6.0 * float64(xs-2) * float64(ys-2))

	return fr * noise
}
