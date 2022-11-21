package utils

import "errors"

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
