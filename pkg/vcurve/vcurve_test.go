package vcurve

import (
	"testing"
)

var (
	points = []Point{
		{X: 29000, Y: 40.5},
		{X: 29100, Y: 36.2},
		{X: 29200, Y: 31.4},
		{X: 29300, Y: 28.6},
		{X: 29400, Y: 23.1},
		{X: 29500, Y: 21.2},
		{X: 29600, Y: 16.6},
		{X: 29700, Y: 13.7},
		{X: 29800, Y: 6.21},
		{X: 29900, Y: 4.21},
		{X: 30000, Y: 3.98},
		{X: 30100, Y: 4.01},
		{X: 30200, Y: 4.85},
		{X: 30300, Y: 11.1},
		{X: 30400, Y: 15.3},
		{X: 30500, Y: 22.1},
		{X: 30600, Y: 21.9},
		{X: 30700, Y: 27.4},
		{X: 30800, Y: 32.1},
		{X: 30900, Y: 36.5},
		{X: 31000, Y: 39.7},
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
