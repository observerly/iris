package utils

import "testing"

func TestRGBUint32n8Bit(t *testing.T) {
	r := RNG{}

	n := r.Uint32n(255)

	if n > 255 {
		t.Errorf("n should be less than 255, but got %v", n)
	}
}

func TestRGBUint32n16Bit(t *testing.T) {
	r := RNG{}

	n := r.Uint32n(65535)

	if n > 65535 {
		t.Errorf("n should be less than 65535, but got %v", n)
	}
}
