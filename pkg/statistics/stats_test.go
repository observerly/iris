package stats

import "testing"

func TestCalculateMinMeanMax(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	min, mean, max := calcMinMeanMax(data)

	if min != 1 {
		t.Errorf("min should be 1, but got %v", min)
	}

	if mean != 5.5 {
		t.Errorf("mean should be 5.5, but got %v", mean)
	}

	if max != 10 {
		t.Errorf("max should be 10, but got %v", max)
	}
}

func TestCalculateMeanStdDevVar(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mean, stddev, variance := calcMeanStdDevVar(data)

	if mean != 5.5 {
		t.Errorf("mean should be 5.5, but got %v", mean)
	}

	if stddev != 2.872281323269 {
		t.Errorf("stddev should be 2.872281323269, but got %v", stddev)
	}

	if variance != 8.25 {
		t.Errorf("variance should be 8.25, but got %v", variance)
	}
}

func TestCalculateMinMeanMaxStdDevVar(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	min, mean, max, stddev, variance := calcMinMeanMaxStdDevVar(data)

	if min != 1 {
		t.Errorf("min should be 1, but got %v", min)
	}

	if mean != 5.5 {
		t.Errorf("mean should be 5.5, but got %v", mean)
	}

	if max != 10 {
		t.Errorf("max should be 10, but got %v", max)
	}

	if stddev != 2.872281323269 {
		t.Errorf("stddev should be 2.872281323269, but got %v", stddev)
	}

	if variance != 8.25 {
		t.Errorf("variance should be 8.25, but got %v", variance)
	}
}

func TestCalculateMedianOdd(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	median := calcMedian(data)

	if median != 5.5 {
		t.Errorf("median should be 5.5, but got %v", median)
	}
}

func TestCalculateMedianEven(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	median := calcMedian(data)

	if median != 6 {
		t.Errorf("median should be 6, but got %v", median)
	}
}

func TestCalculateMedianDispersedRandom(t *testing.T) {
	data := []float32{10, 12, 23, 23, 16, 23, 21, 16}

	median := calcMedian(data)

	if median != 18.5 {
		t.Errorf("median should be 18.5, but got %v", median)
	}
}
