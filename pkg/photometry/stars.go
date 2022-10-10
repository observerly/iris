package photometry

type Star struct {
	Index     int32   // Index of the star in the data array. int32(x)+width*int32(y)
	Value     uint32  // Value of the star in the data array. data[index]
	X         float32 // Precise star x position
	Y         float32 // Precise star y position
	Intensity uint32  // Intensity of the star at position { X, Y }
	HFR       float32 // Half-Flux Radius of the star, in pixels
}

type StarsExtractor struct {
	Width     int
	Height    int
	Threshold float32
	Radius    float32
	Raw       []uint32
	Stars     []Star
	HFR       float32
}

func NewStarsExtractor(exposure [][]uint32, xs int, ys int, radius float32) *StarsExtractor {
	// Locate the brightest pixels in the data array above the threshold and return them as stars:
	var raw []uint32

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, a := range exposure {
		raw = append(raw, a...)
	}

	stars := make([]Star, 0)

	return &StarsExtractor{
		Width:     xs,
		Height:    ys,
		Threshold: 0,
		Radius:    radius,
		Raw:       raw,
		Stars:     stars,
	}
}

func (s *StarsExtractor) GetBrightPixels() []Star {
	stars := s.Stars

	for index, value := range s.Raw {
		// If the value of the pixel is above the threshold, append it to the stars array:
		if float32(value) > s.Threshold {
			is := Star{Index: int32(index), Value: value, X: float32(index % s.Width), Y: float32(index / s.Width), Intensity: value, HFR: 1.0}

			// Check if the star is within the radius of any other stars for optimisation:
			if len(stars) > 0 {
				star := stars[len(stars)-1]

				// Replace previous bright star with current if it is perceived brighter:
				if star.Y == is.Y && star.X >= is.X-float32(s.Radius) && star.Value <= is.Value {
					stars[len(stars)-1] = is
				}
			}

			// Append the star to the stars array:
			stars = append(stars, is)
		}
	}

	return stars
}
