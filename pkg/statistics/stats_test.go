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
