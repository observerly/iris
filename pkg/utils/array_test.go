/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package utils

/*****************************************************************************************************************/

import "testing"

/*****************************************************************************************************************/

func TestAddAB(t *testing.T) {
	a := []float32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	b := []float32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	s, err := AddFloat32Array(a, b)

	if err != nil {
		t.Errorf("error should be nil, but got %v", err)
	}

	if len(s) != len(a) {
		t.Errorf("result should be of same length as a")
	}

	for i := range s {
		if s[i] != 2 {
			t.Errorf("result should be %v, but got %v", a[i]+b[i], s[i])
		}
	}
}

/*****************************************************************************************************************/

func TestAddABNotEqualLengthPanic(t *testing.T) {
	a := []float32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	b := []float32{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	_, err := AddFloat32Array(a, b)

	if err == nil {
		t.Errorf("error should not be nil for two arrays of unequal length")
	}
}

/*****************************************************************************************************************/

func TestSubtractAB(t *testing.T) {
	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s, err := SubtractFloat32Array(a, b)

	if err != nil {
		t.Errorf("error should be nil, but got %v", err)
	}

	if len(s) != len(a) {
		t.Errorf("result should be of same length as a")
	}

	for i := range s {
		if s[i] != 1 {
			t.Errorf("result should be %v, but got %v", a[i]-b[i], s[i])
		}
	}
}

/*****************************************************************************************************************/

func TestSubtractABNotEqualLengthPanic(t *testing.T) {
	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	_, err := SubtractFloat32Array(a, b)

	if err == nil {
		t.Errorf("error should not be nil for two arrays of unequal length")
	}
}

/*****************************************************************************************************************/

func TestMultiplyAB(t *testing.T) {
	a := []float32{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	b := []float32{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	s, err := MultiplyFloat32Array(a, b)

	if err != nil {
		t.Errorf("error should be nil, but got %v", err)
	}

	if len(s) != len(a) {
		t.Errorf("result should be of same length as a")
	}

	for i := range s {
		if s[i] != 4 {
			t.Errorf("result should be %v, but got %v", a[i]*b[i], s[i])
		}
	}
}

/*****************************************************************************************************************/

func TestMultiplyABNotEqualLengthPanic(t *testing.T) {
	a := []float32{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	b := []float32{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}

	_, err := MultiplyFloat32Array(a, b)

	if err == nil {
		t.Errorf("error should not be nil for two arrays of unequal length")
	}
}

/*****************************************************************************************************************/

func TestDivideAB(t *testing.T) {
	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	d, err := DivideFloat32Array(a, b, 10)

	if err != nil {
		t.Errorf("error should be nil, but got %v", err)
	}

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

/*****************************************************************************************************************/

func TestDivideABDegenerate(t *testing.T) {
	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{-1, -2, 3, 4, 5, -6, 7, 8, 9, 10}

	d, err := DivideFloat32Array(a, b, 10)

	if err != nil {
		t.Errorf("error should be nil, but got %v", err)
	}

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

/*****************************************************************************************************************/

func TestDivideABNotEqualLengthPanic(t *testing.T) {

	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	_, err := DivideFloat32Array(a, b, 10)

	if err == nil {
		t.Errorf("error should not be nil for two arrays of unequal length")
	}
}

/*****************************************************************************************************************/

func TestAverageA(t *testing.T) {
	a := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s, err := AverageFloat32Array(a)

	if err != nil {
		t.Errorf("error should be nil, but got %v", err)
	}

	if s != 5.5 {
		t.Errorf("result should be 5.5, but got %v", s)
	}
}

/*****************************************************************************************************************/

func TestMeanABC(t *testing.T) {
	a := []float32{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	c := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	m, err := MeanFloat32Arrays([][]float32{a, b, c})

	if err != nil {
		t.Errorf("error should be nil, but got %v", err)
	}

	if len(m) != len(a) {
		t.Errorf("result should be of same length as a")
	}

	if m[0] != 4 {
		t.Errorf("result should be 1, but got %v", m[0])
	}

	if m[1] != 4.333333333333333 {
		t.Errorf("result should be 1, but got %v", m[1])
	}

	if m[2] != 4.666666666666667 {
		t.Errorf("result should be 1, but got %v", m[2])
	}

	if m[3] != 5 {
		t.Errorf("result should be 1, but got %v", m[3])
	}

	if m[4] != 5.333333333333333 {
		t.Errorf("result should be 6 but got %v", m[4])
	}

	if m[5] != 5.666666666666667 {
		t.Errorf("result should be 6 but got %v", m[5])
	}

	// Assume here that the mean calculation is correct for all other elements
}

/*****************************************************************************************************************/

func TestMeanABCD(t *testing.T) {
	a := []float32{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	c := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	d := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	m, err := MeanFloat32Arrays([][]float32{a, b, c, d})

	if err != nil {
		t.Errorf("error should be nil, but got %v", err)
	}

	if len(m) != len(a) {
		t.Errorf("result should be of same length as a")
	}

	if m[0] != 3.25 {
		t.Errorf("result should be 1, but got %v", m[0])
	}

	if m[1] != 3.75 {
		t.Errorf("result should be 1, but got %v", m[1])
	}

	if m[2] != 4.25 {
		t.Errorf("result should be 1, but got %v", m[2])
	}

	if m[3] != 4.75 {
		t.Errorf("result should be 1, but got %v", m[3])
	}

	if m[4] != 5.25 {
		t.Errorf("result should be 6 but got %v", m[4])
	}

	if m[5] != 5.75 {
		t.Errorf("result should be 6 but got %v", m[5])
	}

	// Assume here that the mean calculation is correct for all other elements
}

/*****************************************************************************************************************/

func TestMeanABNotEqualLengthPanic(t *testing.T) {
	a := []float32{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	b := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	_, err := MeanFloat32Arrays([][]float32{a, b})

	if err == nil {
		t.Errorf("error should not be nil for two arrays of unequal length")
	}
}

/*****************************************************************************************************************/

func TestFlatten2DUInt32Array2Rows5Columns(t *testing.T) {
	a := [][]uint32{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	f := Flatten2DUInt32Array(a)

	if len(f) != 10 {
		t.Errorf("result should be of length 10, but got %v", len(f))
	}

	if f[0] != 1 {
		t.Errorf("result should be 1, but got %v", f[0])
	}

	if f[1] != 2 {
		t.Errorf("result should be 2, but got %v", f[1])
	}

	if f[2] != 3 {
		t.Errorf("result should be 3, but got %v", f[2])
	}

	if f[3] != 4 {
		t.Errorf("result should be 4, but got %v", f[3])
	}

	if f[4] != 5 {
		t.Errorf("result should be 5, but got %v", f[4])
	}

	if f[5] != 6 {
		t.Errorf("result should be 6, but got %v", f[5])
	}

	if f[6] != 7 {
		t.Errorf("result should be 7, but got %v", f[6])
	}

	if f[7] != 8 {
		t.Errorf("result should be 8, but got %v", f[7])
	}

	if f[8] != 9 {
		t.Errorf("result should be 9, but got %v", f[8])
	}

	if f[9] != 10 {
		t.Errorf("result should be 10, but got %v", f[9])
	}
}

/*****************************************************************************************************************/

func TestFlatten2DUInt32Array1Rows5Columns(t *testing.T) {
	a := [][]uint32{
		{1, 2, 3, 4, 5},
	}

	f := Flatten2DUInt32Array(a)

	if len(f) != 5 {
		t.Errorf("result should be of length 10, but got %v", len(f))
	}

	if f[0] != 1 {
		t.Errorf("result should be 1, but got %v", f[0])
	}

	if f[1] != 2 {
		t.Errorf("result should be 2, but got %v", f[1])
	}

	if f[2] != 3 {
		t.Errorf("result should be 3, but got %v", f[2])
	}

	if f[3] != 4 {
		t.Errorf("result should be 4, but got %v", f[3])
	}

	if f[4] != 5 {
		t.Errorf("result should be 5, but got %v", f[4])
	}
}

/*****************************************************************************************************************/

func TestFlatten2DUFloat64Array2Rows5Columns(t *testing.T) {
	a := [][]float64{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	f := Flatten2DFloat64Array(a)

	if len(f) != 10 {
		t.Errorf("result should be of length 10, but got %v", len(f))
	}

	if f[0] != 1 {
		t.Errorf("result should be 1, but got %v", f[0])
	}

	if f[1] != 2 {
		t.Errorf("result should be 2, but got %v", f[1])
	}

	if f[2] != 3 {
		t.Errorf("result should be 3, but got %v", f[2])
	}

	if f[3] != 4 {
		t.Errorf("result should be 4, but got %v", f[3])
	}

	if f[4] != 5 {
		t.Errorf("result should be 5, but got %v", f[4])
	}

	if f[5] != 6 {
		t.Errorf("result should be 6, but got %v", f[5])
	}

	if f[6] != 7 {
		t.Errorf("result should be 7, but got %v", f[6])
	}

	if f[7] != 8 {
		t.Errorf("result should be 8, but got %v", f[7])
	}

	if f[8] != 9 {
		t.Errorf("result should be 9, but got %v", f[8])
	}

	if f[9] != 10 {
		t.Errorf("result should be 10, but got %v", f[9])
	}
}

/*****************************************************************************************************************/

func TestBoundsFloat32Array(t *testing.T) {
	a := []float32{5, 1, 3, 1, 2, 7, 8, 10, 2, 4, 6, 9}

	min, max := BoundsFloat32Array(a)

	if min != 1 {
		t.Errorf("result should be 1, but got %v", min)
	}

	if max != 10 {
		t.Errorf("result should be 10, but got %v", max)
	}
}

/*****************************************************************************************************************/
