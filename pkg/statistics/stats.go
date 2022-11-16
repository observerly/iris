package stats

import (
	"math"
	"sort"

	"github.com/observerly/iris/pkg/qsort"
	"github.com/observerly/iris/pkg/utils"
)

// Statistics on data arrays, calculated on demand
type Stats struct {
	Width    int       // Width of a line in the underlying data array (for noise)
	Data     []float32 // The underlying data array
	Min      float32   // Minimum
	Max      float32   // Maximum
	Mean     float32   // Mean (average)
	StdDev   float32   // Standard Deviation (norm 2, sigma)
	Variance float32   // Variance (sigma^2)
	Location float32   // Selected location indicator (standard: randomized sigma-clipped median using randomized Qn)
	Scale    float32   // Selected scale indicator (standard: randomized Qn)
	Noise    float32   // Noise Estimation
}

func NewStats(data []float32, xs int) *Stats {
	min, mean, max, stddev, variance := calcMinMeanMaxStdDevVar(data)

	return &Stats{
		Width:    xs,
		Data:     data,
		Min:      min,
		Max:      max,
		Mean:     mean,
		StdDev:   stddev,
		Variance: variance,
		Noise:    0,
	}
}

func calcMinMeanMax(data []float32) (min float32, mean float32, max float32) {
	mmin, mmean, mmax := float32(data[0]), float32(0), float32(data[0])

	for _, v := range data {
		if v < mmin {
			mmin = v
		}
		if v > mmax {
			mmax = v
		}
		mmean += v
	}

	return mmin, float32(mmean / float32(len(data))), mmax
}

func calcMeanStdDevVar(data []float32) (mean float32, stddev float32, variance float32) {
	xvar, mmean := float32(0), float32(0)

	for _, v := range data {
		mmean += v
	}

	mmean /= float32(len(data))

	for _, x := range data {
		diff := x - mmean
		xvar += diff * diff
	}

	xvar /= float32(len(data))

	stddev = float32(math.Sqrt(float64(xvar)))

	return mmean, stddev, xvar
}

func calcMinMeanMaxStdDevVar(data []float32) (min float32, mean float32, max float32, stddev float32, variance float32) {
	mmin, mmean, mmax, xvar := float32(data[0]), float32(0), float32(data[0]), float32(0)

	for _, v := range data {
		if v < mmin {
			mmin = v
		}

		if v > mmax {
			mmax = v
		}

		mmean += v
	}

	mmean /= float32(len(data))

	for _, x := range data {
		diff := x - mmean
		xvar += diff * diff
	}

	xvar /= float32(len(data))

	stddev = float32(math.Sqrt(float64(xvar)))

	return mmin, mmean, mmax, stddev, xvar
}

func calcMedian(data []float32) float32 {
	p := make([]float64, len(data))

	for i, v := range data {
		p[i] = float64(v)
	}

	sort.Float64Slice(p).Sort()

	if len(p)%2 == 0 {
		return float32((p[len(p)/2] + p[len(p)/2-1]) / 2)
	}

	return float32(p[len(p)/2])
}

/*
	FastApproxMedian

	Calculates fast approximate median of the (presumably large) data by
	sub-sampling the given number of values and taking the median of that.

	Note: this is not a statistically correct median, but it is fast and
	should be good enough for most purposes.	The sub-sampling is done
	by randomly selecting sub-values from the data array using a random
	number generator pinned to the maximum of the data array.
*/
func (s *Stats) FastApproxMedian(data []float32, sample []float32) float32 {
	rng := utils.RNG{}

	// Obtain the maximum value of the random number generator:
	max := uint32(len(data))

	for i := range sample {
		index := rng.Uint32n(max)
		// Take a sub-sample of the data array:
		sample[i] = data[index]
	}

	median := qsort.QSelectMedianFloat32(sample)

	return median
}

/*
	FastApproxQn

	Calculates fast approximate Qn scale estimate of the (presumably large) data by
	sub-sampling the given number of pairs and taking the first quartile of that.

	Note: this is not a statistically correct median, but it is fast and
	should be good enough for most purposes. The sub-sampling is done
	by randomly selecting sub-values from the data array using a random
	number generator pinned to the maximum of the data array.

	@see http://web.ipac.caltech.edu/staff/fmasci/home/astro_refs/BetterThanMAD.pdf
*/
func (s *Stats) FastApproxQn(data []float32, sample []float32) float32 {
	rng := utils.RNG{}

	// Obtain the maximum value of the random number generator:
	max := uint32(len(data))

	for i := range sample {
		index := 1 + rng.Uint32n(max-1)
		nindex := rng.Uint32n(index)
		// Take a sub-sample of the data array:
		sample[i] = float32(math.Abs(float64(data[index] - data[nindex])))
	}

	// Normalize to the Gaussian standard deviation, for larger samples >> 1000
	// Source for corrected constant https://rdrr.io/cran/robustbase/man/Qn.html
	qn := qsort.QSelectFirstQuartileFloat32(sample) * 2.21914

	return qn
}

/*
	FastApproxBoundedMedian

	Calculates fast approximate median of the (presumably large) data by
	sub-sampling the given number of values and taking the median of that.

	Note: this is not a statistically correct median, but it is fast and
	should be good enough for most purposes. The sub-sampling is done
	by randomly selecting sub-values from the data array using a random
	number generator pinned to the maximum of the data array.
*/
func (s *Stats) FastApproxBoundedMedian(data []float32, sample []float32, lowerBound, higherBound float32) float32 {
	rng := utils.RNG{}

	// Obtain the maximum value of the random number generator:
	max := uint32(len(data))

	for i := range sample {
		var d float32
		for {
			d = data[rng.Uint32n(max)]

			if d >= lowerBound && d <= higherBound {
				break
			}
		}
		// Take a sub-sample of the data array:
		sample[i] = d
	}

	median := qsort.QSelectMedianFloat32(sample)

	return median
}
