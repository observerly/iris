/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/photometry
//	@license	Copyright © 2021-2025 observerly

/*****************************************************************************************************************/

package photometry

/*****************************************************************************************************************/

import (
	"math"

	stats "github.com/observerly/iris/pkg/statistics"
	"github.com/observerly/iris/pkg/utils"
)

/*****************************************************************************************************************/

type Star struct {
	Index     int32   // Index of the star in the data array. int32(x)+width*int32(y)
	Value     float32 // Value of the star in the data array. data[index]
	X         float32 // Precise star x position
	Y         float32 // Precise star y position
	Intensity float32 // Intensity of the star at position { X, Y }
	HFR       float32 // Half-Flux Radius of the star, in pixels
}

/*****************************************************************************************************************/

type StarLink struct {
	Star *Star
	Next *StarLink
}

/*****************************************************************************************************************/

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

/*****************************************************************************************************************/

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

/*****************************************************************************************************************/

func (s *StarsExtractor) GetBrightPixels() []Star {
	return findBrightPixels(s.Data, int32(s.Width), s.Threshold, int32(s.Radius))
}

/*****************************************************************************************************************/

func (s *StarsExtractor) RejectBadPixels() []Star {
	return rejectBadPixels(s.Stars, s.Data, int32(s.Width), s.Sigma, s.ADU)
}

/*****************************************************************************************************************/

func (s *StarsExtractor) RejectMisalignedPixels() []Star {
	return rejectMisalignedPixels(s.Stars, s.Data, int32(s.Width), int32(s.Height), s.Radius, s.Radius/2.0)
}

/*****************************************************************************************************************/

func (s *StarsExtractor) FilterOverlappingPixels() []Star {
	return filterOverlappingPixels(s.Stars, int32(s.Width), int32(s.Height), int32(s.Radius))
}

/*****************************************************************************************************************/

func (s *StarsExtractor) ShiftToCenterOfMass() []Star {
	return shiftToCenterOfMass(s.Stars, s.Data, int32(s.Width), s.Threshold, int32(s.Radius))
}

/*****************************************************************************************************************/

func (s *StarsExtractor) ExtractAndFilterHalfFluxRadius(location float32, starInOut float32) []Star {
	stars, HFR := extractAndFilterHalfFluxRadius(s.Stars, s.Data, int32(s.Width), s.Radius, location, starInOut)

	s.HFR = HFR

	return stars
}

/*****************************************************************************************************************/

// FindStars finds stars in the data array, using the given threshold, sigma and radius.
func (s *StarsExtractor) FindStars(stats *stats.Stats, sigma float32, starInOut float32) []Star {
	// Set the global sigma:
	s.Sigma = sigma

	// Find the location and scale:
	location, scale := stats.FastApproxSigmaClippedMedianAndQn()

	s.Threshold = location + scale*sigma

	s.Stars = s.GetBrightPixels()

	s.Stars = s.RejectBadPixels()

	s.Stars = s.FilterOverlappingPixels()

	// Shift stars to their nearest center of mass (centroid position):
	s.Stars = s.ShiftToCenterOfMass()

	s.Stars = s.FilterOverlappingPixels()

	// Reject extracted based on proximity to other stars, e.g., if we have two stars in close proximity
	// we discard both as the extraction pixel position is likely to be incorrect:
	s.Stars = s.RejectMisalignedPixels()

	// Remove implausible stars based on HFR, and return stars (updates s.HFR):
	stars := s.ExtractAndFilterHalfFluxRadius(location, starInOut)

	// Return a clone of the final shortlist of stars, so the longer original object can be reclaimed:
	result := make([]Star, len(stars))

	copy(result, stars)

	stars = nil

	return result
}

/*****************************************************************************************************************/

// Applies an element-wise Median filter to the sparse data points provided by the indices, with the local
// neighborhood defined by the mask.
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

/*****************************************************************************************************************/

// Finds bright pixels in the data array, using the given threshold and radius, applying an
// early overlap rejection to reduce allocations. Use a central pixel value as initial mass, and 1 as initial HFR.
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

/*****************************************************************************************************************/

// Rejects bad pixels which differ from the local median by more than sigma times the estimated standard deviation.
// The stars array is modified in place, and the shortened slice is returned.
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

/*****************************************************************************************************************/

// rejectMisalignedPixels filters out stars that are not aligned with the brightest pixel in their vicinity.
// This is likely due to two stars in close proximity in pixel space, causing the centroid to be misaligned.
func rejectMisalignedPixels(stars []Star, data []float32, xs int32, ys int32, radius float32, maxAllowedDistance float32) []Star {
	remainingStars := 0

	for _, s := range stars {
		// Define the search window around the star's centroid position:
		xStart := int32(s.X - radius)
		xEnd := int32(s.X + radius)
		yStart := int32(s.Y - radius)
		yEnd := int32(s.Y + radius)

		if xStart < 0 {
			xStart = 0
		}

		if xEnd >= xs {
			xEnd = xs - 1
		}

		if yStart < 0 {
			yStart = 0
		}

		if yEnd >= ys {
			yEnd = ys - 1
		}

		// Find the brightest pixel within the radius of the current star:
		maxIntensity := float32(0)
		brightestX := s.X
		brightestY := s.Y

		for y := yStart; y <= yEnd; y++ {
			for x := xStart; x <= xEnd; x++ {
				index := y*xs + x
				intensity := data[index]
				if intensity > maxIntensity {
					maxIntensity = intensity
					brightestX = float32(x)
					brightestY = float32(y)
				}
			}
		}

		// Calculate the distance between centroid and brightest pixel in the vicinity:
		dx := s.X - brightestX
		dy := s.Y - brightestY
		distanceSquared := dx*dx + dy*dy

		// If the distance is within the allowed threshold, keep the star (increment remainingStars), otherwise,
		// discard the star (do not increment remainingStars) and continue to the next star:
		if distanceSquared <= maxAllowedDistance*maxAllowedDistance {
			stars[remainingStars] = s
			remainingStars++
		}

	}

	return stars[:remainingStars]
}

/*****************************************************************************************************************/

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

/*****************************************************************************************************************/

// Shifts the star positions to the center of mass of the star.
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

/*****************************************************************************************************************/

// Extracts the half-flux radius (HFR) of each star, and filters out implausible candidates. Returns a new list
// of stars, each enriched with the HFR field and updated mass.
func extractAndFilterHalfFluxRadius(stars []Star, data []float32, xs int32, radius, location, starInOut float32) (res []Star, avgHFR float32) {
	remainingStars := 0

	avgHFR = float32(0)

	for _, s := range stars {
		// Calculate mass, moment and HFR for star:
		moment, mass, pixels := float32(0), float32(0), int32(0)

		rad := int32(math.Ceil(float64(radius)))

		distSqLimit := int32(math.Ceil(float64(radius+1e-8) * float64(radius+1e-8)))

		for y := -rad; y <= rad; y++ {
			for x := -rad; x <= rad; x++ {
				distSq := x*x + y*y
				if distSq > distSqLimit {
					continue
				}
				distance := float32(math.Sqrt(float64(distSq)))

				index := s.Index + y*xs + x

				value := float32(0.0)

				if index >= 0 && index < int32(len(data)) {
					v := data[index] - location
					if v > 0 {
						value = v
					}
				}
				moment += distance * value
				mass += value
				pixels++
			}
		}

		if mass == 0.0 {
			mass = 1e-8
		}

		hfr := float32(moment / mass)

		// Sanity check results to avoid long lockups:
		if hfr > radius {
			continue
		}

		// Calculate mass inside HFR and number of inner pixels:
		innerMass, innerPixels := float32(0), int32(0)

		innerRad := int32(math.Ceil(float64(hfr)))

		distSqLimit = int32(math.Ceil(float64(hfr * hfr)))

		for y := -innerRad; y <= innerRad; y++ {
			for x := -innerRad; x <= innerRad; x++ {
				distSq := x*x + y*y

				if distSq > distSqLimit {
					continue
				}

				index := s.Index + y*xs + x

				value := float32(0.0)

				if index >= 0 && index < int32(len(data)) {
					v := data[index] - location
					if v > 0 {
						value = v
					}
				}
				innerMass += value
				innerPixels++
			}
		}

		// Plausibility check i.e., is the average inner brightness significantly higher than outside?
		outerMass := mass - innerMass

		outerPixels := pixels - innerPixels

		if innerMass*float32(outerPixels) <= starInOut*outerMass*float32(innerPixels) {
			continue
		}

		// Enrich star with HFR and mass information:
		s.HFR = hfr

		s.Intensity = mass

		stars[remainingStars] = s

		remainingStars++

		// Add to the average HFR:
		avgHFR += float32(hfr)
	}

	// Calculate the average HFR:
	avgHFR /= float32(remainingStars)
	// Return the remaining stars and the average HFR:
	return stars[:remainingStars], avgHFR
}

/*****************************************************************************************************************/
