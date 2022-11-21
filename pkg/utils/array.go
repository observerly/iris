package utils

import (
	"errors"
	"fmt"
)

/**
  Subtract

  Computes the element-wise difference of arrays a and b
  and stores in array d "divide", that is, d[i]=a[i]-b[i].
**/
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

/**
  Divide

  Computes the element-wise division of arrays a and b, scaled
  with bMean and stores in array d "divide", that is, d[i]=a[i]*bMax/b[i].
**/
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

/**
	Mean

	Computes the mean of array a and stores in array m "mean",
	that is, m[i]=mean(a). If a is empty, m is nil.
**/
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
