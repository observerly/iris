package utils

import "testing"

func TestSubtractAB(t *testing.T) {
	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s := SubtractFloat32Array(a, b)

	if len(s) != len(a) {
		t.Errorf("result should be of same length as a")
	}

	for i := range s {
		if s[i] != 1 {
			t.Errorf("result should be %v, but got %v", a[i]-b[i], s[i])
		}
	}
}

func TestSubtractABNotEqualLengthPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	SubtractFloat32Array(a, b)
}
