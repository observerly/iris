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

func TestDivideAB(t *testing.T) {
	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	d := DivideFloat32Array(a, b, 10)

	if len(d) != len(a) {
		t.Errorf("result should be of same length as a")
	}

	if d[0] != a[0]*10/b[0] {
		t.Errorf("result should be %v, but got %v", a[0]*10/b[0], d[0])
	}

	if d[5] != a[5]*10/b[5] {
		t.Errorf("result should be %v, but got %v", a[5]*10/b[5], d[5])
	}
}

func TestDivideABDegenerate(t *testing.T) {
	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{-1, -2, 3, 4, 5, -6, 7, 8, 9, 10}

	d := DivideFloat32Array(a, b, 10)

	if len(d) != len(a) {
		t.Errorf("result should be of same length as a")
	}

	if d[0] != a[0] {
		t.Errorf("result should be %v, but got %v", a[0], d[0])
	}

	if d[1] != a[1] {
		t.Errorf("result should be %v, but got %v", a[1], d[1])
	}

	if d[5] != a[5] {
		t.Errorf("result should be %v, but got %v", a[5], d[5])
	}
}

func TestDivideABNotEqualLengthPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	DivideFloat32Array(a, b, 10)
}
