package photometry

import (
	stats "github.com/observerly/iris/pkg/statistics"
	"github.com/observerly/iris/pkg/utils"
)

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
	Sigma     float32
	Radius    float32
	Data      []float32
	ADU       int32
	Stars     []Star
	HFR       float32
}

func NewStarsExtractor(data []float32, xs int, ys int, radius float32, adu int32) *StarsExtractor {
	stars := make([]Star, 0)

	return &StarsExtractor{
		Width:     xs,
		Height:    ys,
		Threshold: 0,
		Sigma:     2.5,
		Radius:    radius,
		Data:      data,
		ADU:       adu,
		Stars:     stars,
	}
}

func (s *StarsExtractor) GetBrightPixels() []Star {
	return findBrightPixels(s.Data, int32(s.Width), s.Threshold, int32(s.Radius))
}

func (s *StarsExtractor) RejectBadPixels() []Star {
	return rejectBadPixels(s.Stars, s.Data, int32(s.Width), s.Sigma, s.ADU)
}

/**
	gatherNeighbourhoodAndCalcMedian()

	Applies an element-wise Median filter to the sparse data points provided by the indices,
	with the local neighborhood defined by the mask.

	@returns the median of the masked neighborhood:
**/
func gatherNeighbourhoodAndCalcMedian(data []float32, index int32, mask []int32, buffer []float32, adu int32) float32 {
	// Gather the neighborhood of each indexed data point into an array:
	num := 0

	for _, o := range mask {
		// Source Offset:
		so := index + o

		if so >= 0 && so < int32(len(data)) {
			buffer[num] = data[so]
			num++
		}
	}

	s := stats.NewStats(buffer, adu, num)

	return s.FastMedian()
}

/**
	findBrightPixels()

	Find pixels above the threshold and return them as stars. Applies early
	overlap rejection based on radius to reduce allocations.

	Uses central pixel value as initial mass, 1 as initial HFR.
**/
func findBrightPixels(data []float32, width int32, threshold float32, radius int32) []Star {
	// Roughly, we'll locate 100 bright stars []Star{}
	stars := make([]Star, 0)

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

/**
	rejectBadPixels()

	Reject bad pixels which differ from the local median by more than sigma times the
	estimated standard deviation.

	Modifies the given stars array values, and returns shortened slice.
**/
func rejectBadPixels(stars []Star, data []float32, width int32, sigma float32, adu int32) []Star {
	// Create a radial mask for local 9-pixel neighborhood:
	mask := utils.CreateRadialPixelMask(int32(width), 1.5)

	buffer := make([]float32, len(mask))

	// Estimate standard deviation of pixels from local neighborhood median based on random 1% of pixels
	numSamples := len(data) / 100

	sample := make([]float32, numSamples)

	rng := utils.RNG{}

	for i := 0; i < numSamples; i++ {
		index := int32(rng.Uint32n(uint32(len(data))))

		median := gatherNeighbourhoodAndCalcMedian(data, index, mask, buffer, adu)

		sample[i] = data[index] - median
	}

	s := stats.NewStats(sample, adu, numSamples)

	// Filter out star candidates more than sigma standard deviations away from the estimated local median
	threshold := s.StdDev * sigma

	remainingStars := 0

	for _, star := range stars {
		median := gatherNeighbourhoodAndCalcMedian(data, star.Index, mask, buffer, adu)

		diff := data[star.Index] - median

		// Accept star if it is less than sigma standard deviations away from the local median:
		if diff < threshold && -diff < threshold {
			stars[remainingStars] = star
			remainingStars++
		}
	}

	return stars[:remainingStars]
}
