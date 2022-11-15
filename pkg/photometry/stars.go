package photometry

type Star struct {
	Index     int32   // Index of the star in the data array. int32(x)+width*int32(y)
	Value     float32 // Value of the star in the data array. data[index]
	X         float32 // Precise star x position
	Y         float32 // Precise star y position
	Intensity float32 // Intensity of the star at position { X, Y }
	HFR       float32 // Half-Flux Radius of the star, in pixels
}

type StarsExtractor struct {
	Width     int
	Height    int
	Threshold float32
	Radius    float32
	Data      []float32
	Stars     []Star
	HFR       float32
}

func NewStarsExtractor(data []float32, xs int, ys int, radius float32) *StarsExtractor {
	stars := make([]Star, 100)

	return &StarsExtractor{
		Width:     xs,
		Height:    ys,
		Threshold: 0,
		Radius:    radius,
		Data:      data,
		Stars:     stars,
	}
}

func (s *StarsExtractor) GetBrightPixels() []Star {
	return findBrightPixels(s.Data, int32(s.Width), s.Threshold, int32(s.Radius))
}

/**
	findBrightPixels()

	Find pixels above the threshold and return them as stars. Applies early
	overlap rejection based on radius to reduce allocations.

	Uses central pixel value as initial mass, 1 as initial HFR.
**/
func findBrightPixels(data []float32, width int32, threshold float32, radius int32) []Star {
	// Roughly, we'll locate 100 bright stars []Star{}
	stars := make([]Star, len(data)/100)[:0]

	for i, v := range data {
		if v > threshold {
			is := Star{
				Index:     int32(i),
				Value:     v,
				X:         float32(int32(i) % width),
				Y:         float32(int32(i) / width),
				Intensity: v,
				HFR:       1,
			}

			// Check if the star is within the radius of any other stars for optimisation:
			if len(stars) > 0 {
				star := stars[len(stars)-1]

				// Replace previous bright star with current if it is perceived brighter:
				if star.Y == is.Y && star.X >= is.X-float32(radius) && star.Value <= is.Value {
					stars[len(stars)-1] = is
				}
			}

			// Append the star to the stars array:
			stars = append(stars, is)
		}
	}

	return stars
}
