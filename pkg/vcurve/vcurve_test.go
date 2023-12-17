package vcurve

import (
	"testing"
)

var (
	points = []Point{
		{x: 29000, y: 40.5},
		{x: 29100, y: 36.2},
		{x: 29200, y: 31.4},
		{x: 29300, y: 28.6},
		{x: 29400, y: 23.1},
		{x: 29500, y: 21.2},
		{x: 29600, y: 16.6},
		{x: 29700, y: 13.7},
		{x: 29800, y: 6.21},
		{x: 29900, y: 4.21},
		{x: 30000, y: 3.98},
		{x: 30100, y: 4.01},
		{x: 30200, y: 4.85},
		{x: 30300, y: 11.1},
		{x: 30400, y: 15.3},
		{x: 30500, y: 22.1},
		{x: 30600, y: 21.9},
		{x: 30700, y: 27.4},
		{x: 30800, y: 32.1},
		{x: 30900, y: 36.5},
		{x: 31000, y: 39.7},
	}
)

func TestNewHyperbolicVCurve(t *testing.T) {
	v := NewHyperbolicVCurve(VCurve{
		Points: points,
	})

	// Deconstruct the VCurveParams struct
	a, b, c, d := v.A, v.B, v.C, v.D

	// Expect the initial guess for A to be the mean value of the yData:
	if a != 21 {
		t.Errorf("A should be 21, but got %v", a)
	}

	// Expect the initial guess for B to be the min value of the yData:
	if b != 3.98 {
		t.Errorf("B should be 3.98, but got %v", b)
	}

	// Expect the initial guess for C to be the mean value of the xData:
	if c != 30000 {
		t.Errorf("C should be 30000, but got %v", c)
	}

	// Expect the initial guess for D to be 0:
	if d != 0 {
		t.Errorf("D should be 0, but got %v", d)
	}
}

// Test for V-curve fitting with hyperbolic model
func TestHyperbolicVCurveLevenbergMarquardtOptimisation(t *testing.T) {
	v := NewHyperbolicVCurve(VCurve{
		Points: points,
	})

	optimisedParams, err := v.LevenbergMarquardtOptimisation()

	// Expect no error!
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// Deconstruct the VCurveParams struct
	a, b, c, d := optimisedParams.A, optimisedParams.B, optimisedParams.C, optimisedParams.D

	// Expect the optimised parameters to be close to the actual parameters:
	// a = 106.8247
	if a-106.8247 > 0.0001 {
		t.Errorf("A should be close to %v, but got %v", 106.8247, a)
	}

	// b = 4.4771
	if b-4.4771 > 0.0001 {
		t.Errorf("B should be close to %v, but got %v", 4.4771, b)
	}

	// c = 30008.5444
	if c-30008.5444 > 0.0001 {
		t.Errorf("C should be close to %v, but got %v", 30008.5444, c)
	}

	// d = -1.7965
	if d+1.7965 > 0.0001 {
		t.Errorf("D should be close to %v, but got %v", -1.7965, d)
	}
}
