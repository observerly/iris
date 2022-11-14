package qsort

import "testing"

func TestQPartitionFloat32(t *testing.T) {
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	pivot := QPartitionFloat32(a)

	if pivot != 4 {
		t.Errorf("Expected pivot at index 4, got %d", pivot)
	}

	if a[pivot] != 5 {
		t.Errorf("Expected pivot to be 5, got %f", a[pivot])
	}
}

func TestQPartitionFloat32DispersedRandom(t *testing.T) {
	a := []float32{10, 12, 23, 23, 16, 23, 21, 16}

	pivot := QPartitionFloat32(a)

	if pivot != 5 {
		t.Errorf("Expected pivot at index 3, got %d", pivot)
	}

	if a[pivot] != 23 {
		t.Errorf("Expected pivot to be 16, got %f", a[pivot])
	}
}
