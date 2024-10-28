/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris
//	@license	Copyright © 2021-2024 observerly

/*****************************************************************************************************************/

package utils

/*****************************************************************************************************************/

import (
	"errors"
	"fmt"
	"math"
)

/*****************************************************************************************************************/

/*
Add

Computes the element-wise sum of arrays a and b and stores in array s "sum",
that is, s[i]=a[i]+b[i].
*/
func AddFloat32Array(a, b []float32) ([]float32, error) {
	if len(a) != len(b) {
		return nil, errors.New("to add arrays they must be of same length")
	}

	s := make([]float32, len(a))

	for i := range s {
		s[i] = a[i] + b[i]
	}

	return s, nil
}

/*****************************************************************************************************************/

/*
Subtract

Computes the element-wise difference of arrays a and b
and stores in array d "divide", that is, d[i]=a[i]-b[i].
*/
func SubtractFloat32Array(a, b []float32) ([]float32, error) {
	if len(a) != len(b) {
		return nil, errors.New("to subtract arrays they must be of same length")
	}

	s := make([]float32, len(a))

	for i := range s {
		s[i] = a[i] - b[i]
	}

	return s, nil
}

/*****************************************************************************************************************/

/*
Multiply

Computes the element-wise product of arrays a and b and stores
in array p "product", that is, m[i]=a[i]*b[i].
*/
func MultiplyFloat32Array(a, b []float32) ([]float32, error) {
	if len(a) != len(b) {
		return nil, errors.New("to multiply arrays they must be of same length")
	}

	p := make([]float32, len(a))

	for i := range p {
		p[i] = a[i] * b[i]
	}

	return p, nil
}

/*****************************************************************************************************************/

/*
Divide

Computes the element-wise division of arrays a and b, scaled
with bMean and stores in array d "divide", that is, d[i]=a[i]*bMax/b[i].
*/
func DivideFloat32Array(a, b []float32, bMax float32) ([]float32, error) {
	if len(a) != len(b) {
		return nil, errors.New("to divide arrays they must be of same length")
	}

	d := make([]float32, len(a))

	for i := range d {
		index := b[i]

		if index <= 0 {
			d[i] = a[i]
		} else {
			d[i] = a[i] * bMax / b[i]
		}
	}

	return d, nil
}

/*****************************************************************************************************************/

/*
Average

Computes the average of array a and stores in array m "mean",
that is, m[i]=mean(a). If a is empty, m is nil.
*/
func AverageFloat32Array(a []float32) (float32, error) {
	if len(a) == 0 {
		return 0, errors.New("cannot compute the average of an empty array")
	}

	var sum float32

	for _, i := range a {
		sum += i
	}

	return sum / float32(len(a)), nil
}

/*****************************************************************************************************************/

/*
Mean

Computes the mean of array a and stores in array m "mean",
that is, m[i]=mean(a). If a is empty, m is nil.
*/
func MeanFloat32Arrays(a [][]float32) ([]float32, error) {
	if len(a) == 0 {
		return nil, errors.New("to divide arrays they must be of same length")
	}

	m := make([]float32, len(a[0]))

	for i := range m {
		for j := range a {
			// Ensure that each sub-array has the same length as the first one:
			if len(a[j]) != len(a[0]) {
				return nil, fmt.Errorf("issue at array input %v: to compute the mean of multiple arrays the length of each array must be the same", i)
			}

			if a[j][i] == 0 {
				continue
			}

			m[i] += a[j][i]
		}

		m[i] /= float32(len(a))
	}

	return m, nil
}

/*****************************************************************************************************************/

/*
Flatten2DUInt32Array

Flattens a 2D array of uint32 into a 1D array of float32.
*/
func Flatten2DUInt32Array(a [][]uint32) []float32 {
	f := make([]float32, 0)

	for _, j := range a {
		for _, i := range j {
			f = append(f, float32(i))
		}
	}

	return f
}

/*****************************************************************************************************************/

/*
Flatten2DFloat64Array

Flattens a 2D array of float64 into a 1D array of float64.
*/
func Flatten2DFloat64Array(a [][]float64) []float64 {
	f := make([]float64, 0)

	for _, j := range a {
		f = append(f, j...)
	}

	return f
}

/*****************************************************************************************************************/

/*
BoundsFloat32Array

Computes the minimum and maximum values of array a.
*/
func BoundsFloat32Array(a []float32) (float32, float32) {
	// Set the initial min and max values:
	min, max := float32(math.MaxFloat32-1), float32(0.0)

	for _, j := range a {
		if j < min {
			min = j
		} else if j > max {
			max = j
		}
	}

	return min, max
}

/*****************************************************************************************************************/
