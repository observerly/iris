package utils

/**
  Subtract

  Computes the element-wise difference of arrays a and b
  and stores in array d "divide", that is, d[i]=a[i]-b[i].
**/
func SubtractFloat32Array(a, b []float32) []float32 {
	if len(a) != len(b) {
		panic("to subtract arrays must be of same length")
	}

	s := make([]float32, len(a))

	for i := range s {
		s[i] = a[i] - b[i]
	}

	return s
}

