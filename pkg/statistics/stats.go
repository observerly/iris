package stats

import "math"

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
	if len(data)%2 == 0 {
		return (data[len(data)/2-1] + data[len(data)/2]) / 2
	}
	return data[len(data)/2]
}
