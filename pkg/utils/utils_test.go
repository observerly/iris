package utils

import "testing"

func TestCreateMask100Width5Radius(t *testing.T) {
	mask := CreateRadialPixelMask(100, 5)

	if len(mask) != 81 {
		t.Errorf("Expected 121, got %d", len(mask))
	}
}

func TestCreateMask200Width5Radius(t *testing.T) {
	mask := CreateRadialPixelMask(200, 5)

	if len(mask) != 81 {
		t.Errorf("Expected 121, got %d", len(mask))
	}
}

func TestCreateMask100Width10Radius(t *testing.T) {
	mask := CreateRadialPixelMask(100, 10)

	if len(mask) != 317 {
		t.Errorf("Expected 121, got %d", len(mask))
	}
}

func TestCreateMask200Width10Radius(t *testing.T) {
	mask := CreateRadialPixelMask(200, 10)

	if len(mask) != 317 {
		t.Errorf("Expected 121, got %d", len(mask))
	}
}
