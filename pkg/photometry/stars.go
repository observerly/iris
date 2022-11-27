package photometry

import (
	"math"

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

type StarLink struct {
	Star *Star
	Next *StarLink
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

func (s *StarsExtractor) FilterOverlappingPixels() []Star {
	return filterOverlappingPixels(s.Stars, int32(s.Width), int32(s.Height), int32(s.Radius))
}

func (s *StarsExtractor) ShiftToCenterOfMass() []Star {
	return shiftToCenterOfMass(s.Stars, s.Data, int32(s.Width), s.Threshold, int32(s.Radius))
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
func findBrightPixels(data []float32, xs int32, threshold float32, radius int32) []Star {
	// Roughly, we'll locate 100 bright stars []Star{}
	stars := make([]Star, 0)

	for i, v := range data {
		if v > threshold {
			is := Star{
				Index:     int32(i),
				Value:     v,
				X:         float32(int32(i) % xs),
				Y:         float32(int32(i) / xs),
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
func rejectBadPixels(stars []Star, data []float32, xs int32, sigma float32, adu int32) []Star {
	// Create a radial mask for local 9-pixel neighborhood:
	mask := utils.CreateRadialPixelMask(int32(xs), 1.5)

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

func filterOverlappingPixels(stars []Star, xs int32, ys int32, radius int32) []Star {
	// To avoid quadratic search effort, we bin the stars into a 2D grid.
	binSize := int32(256)

	xBins := (xs + binSize - 1) / binSize
	yBins := (ys + binSize - 1) / binSize

	// Each bin is a linked list of stars, sorted by descending mass
	bins := make([]*StarLink, int(xBins*yBins))
	slis := make([]StarLink, ((len(stars)+1023)/1024)*1024) // use tiered sizing to help the allocator

	r2 := radius * radius

	remainingStars := 0

	// For all stars, filter list in place
forAllStars:
	for _, s := range stars {
		// Find grid cell of this star
		xCell, yCell := int32(s.X+0.5)/binSize, int32(s.Y+0.5)/binSize

		// For this grid cell and all adjacent cells
		for dy := int32(-1); dy <= 1; dy++ {
			if yCell+dy < 0 || yCell+dy >= yBins {
				continue
			}

			for dx := int32(-1); dx <= 1; dx++ {
				if xCell+dx < 0 || xCell+dx >= xBins {
					continue
				}

				// Cell index:
				ci := (xCell + dx) + (yCell+dy)*xBins

				// For all prior stars in that cell
				for ptr := bins[ci]; ptr != nil; ptr = ptr.Next {
					s2 := ptr.Star
					xDist := s.X - s2.X
					yDist := s.Y - s2.Y
					sqDist := int32(xDist*xDist + yDist*yDist + 0.5)

					// Skip current star if it's close to a prior star
					if sqDist <= r2 {
						continue forAllStars
					}
				}
			}
		}

		// Retain star for output
		stars[remainingStars] = s

		// Insert star into grid cell
		slis[remainingStars] = StarLink{&(stars[remainingStars]), nil}

		// Cell index:
		ci := xCell + yCell*xBins

		ptr := bins[ci]

		if ptr == nil {
			bins[ci] = &(slis[remainingStars])
		} else {
			for ptr.Next != nil {
				ptr = ptr.Next
			}

			ptr.Next = &(slis[remainingStars])
		}

		remainingStars++
	}

	bins = nil
	slis = nil

	return stars[:remainingStars]
}

/**
	shiftToCenterOfMass

	Shift the star positions to the center of mass of the star.

	Modifies the given stars array values in place.
**/
func shiftToCenterOfMass(stars []Star, data []float32, xs int32, threshold float32, radius int32) []Star {
	// For all stars, shift to center of mass:
	for i, s := range stars {
		// until the shifts are below 0.01 pixel (i.e. 0.0001 squared error), or max rounds reached
		shiftSquared := float32(math.MaxFloat32)
		for round := int32(0); shiftSquared > 0.0001 && round < 10; round++ {
			// calculate star mass and first moments from current x,y
			xMoment, yMoment := float32(0), float32(0)

			mass := float32(0)

			for y := -radius; y <= radius; y++ {
				for x := -radius; x <= radius; x++ {
					index := s.Index + y*int32(xs) + x

					value := float32(0)

					if index >= 0 && int(index) < len(data) {
						value = data[index] - threshold
						if value < 0 {
							value = 0
						}
					}

					xMoment += float32(x) * value
					yMoment += float32(y) * value
					mass += value
				}
			}

			// Update x and y from moments over mass
			x := s.Index % int32(xs)
			y := s.Index / int32(xs)

			if mass == 0.0 {
				mass = 1e-8
			}

			deltaX := (xMoment) / mass
			deltaY := (yMoment) / mass

			newX := float32(x) + deltaX
			newY := float32(y) + deltaY

			preciseDeltaX := newX - s.X
			preciseDeltaY := newY - s.Y

			shiftSquared = preciseDeltaX*preciseDeltaX + preciseDeltaY*preciseDeltaY

			index := s.Index + xs*int32(deltaY+0.5) + int32(deltaX+0.5)

			value := float32(0)

			if index >= 0 && int(index) < len(data) {
				value = float32(data[index])
			}

			s = Star{
				Index:     index,
				Value:     value,
				X:         float32(newX),
				Y:         float32(newY),
				Intensity: float32(mass),
			}

			stars[i] = s
		}
	}

	return stars
}
