package stats

import (
	"math"
	"sort"
)

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
