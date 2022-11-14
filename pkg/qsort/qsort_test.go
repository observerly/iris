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

func TestQSortFloat32(t *testing.T) {
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	QSortFloat32(a)

	for i, v := range a {
		if v != float32(i+1) {
			t.Errorf("Expected %d, got %f", i+1, v)
		}
	}
}

func TestQSortFloat32DispersedRandom(t *testing.T) {
	a := []float32{10, 12, 23, 23, 16, 23, 21, 16}

	QSortFloat32(a)

	if a[0] != 10 {
		t.Errorf("Expected 10, got %f", a[0])
	}

	if a[1] != 12 {
		t.Errorf("Expected 12, got %f", a[1])
	}

	if a[2] != 16 {
		t.Errorf("Expected 16, got %f", a[2])
	}

	if a[3] != 16 {
		t.Errorf("Expected 16, got %f", a[3])
	}

	if a[4] != 21 {
		t.Errorf("Expected 21, got %f", a[4])
	}

	if a[5] != 23 {
		t.Errorf("Expected 23, got %f", a[5])
	}

	if a[6] != 23 {
		t.Errorf("Expected 23, got %f", a[6])
	}

	if a[7] != 23 {
		t.Errorf("Expected 23, got %f", a[7])
	}
}

func TestQSelectFloat32(t *testing.T) {
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	v := QSelectFloat32(a, 5)

	if v != 5 {
		t.Errorf("Expected 5, got %f", v)
	}
}

func TestQSelectFloat32DispersedRandom(t *testing.T) {
	a := []float32{10, 12, 23, 23, 16, 23, 21, 16}

	v := QSelectFloat32(a, 5)

	if v != 21 {
		t.Errorf("Expected 21, got %f", v)
	}
}

func TestQSelectFirstQuartileFloat32(t *testing.T) {
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	v := QSelectFirstQuartileFloat32(a)

	if v != 3 {
		t.Errorf("Expected 2, got %f", v)
	}
}

func TestQSelectFirstQuartileFloat32DispersedRandom(t *testing.T) {
	a := []float32{10, 12, 23, 23, 16, 23, 21, 16}

	v := QSelectFirstQuartileFloat32(a)

	if v != 16 {
		t.Errorf("Expected 12, got %f", v)
	}
}
