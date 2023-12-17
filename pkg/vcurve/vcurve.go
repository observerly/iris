package vcurve

import (
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

// Point is a data point with x and y coordinates.
type Point struct {
	x float64
	y float64
}

// VCurve is a struct that holds the data points for the V-curve.
type VCurve struct {
	Points []Point
}

// VCurveParams is a struct that holds the parameters for the V-curve model optisation and the data points.
type VCurveParams struct {
	A float64
	B float64
	C float64
	D float64
	x []float64
	y []float64
}

/*
NewHyperbolicVCurve

Creates a new VCurve object ready for applying the Levenberg-Marquardt iterative optimization technique.

The VCurve object is initialized with the data points, and initial guesses for the parameters are calculated from the input data.
*/
func NewHyperbolicVCurve(data VCurve) *VCurveParams {
	// Preallocate slices with the exact required capacity
	dataX := make([]float64, 0, len(data.Points))
	dataY := make([]float64, 0, len(data.Points))

	// A single loop to populate the slices
	for _, point := range data.Points {
		dataX = append(dataX, point.x)
		dataY = append(dataY, point.y)
	}

	// Get the initial guess for the parameter, for A, we take the mean value of the yData:
	A := stat.Mean(dataY, nil)

	// Get the initial guess for B, for B, we take the min value of the yData:
	B := floats.Min(dataY)

	// Get the initial guess for C, for C, we take the mean value of the xData:
	C := stat.Mean(dataX, nil)

	return &VCurveParams{
		A: math.Round(A),
		B: B,
		C: C,
		D: 0,
		x: dataX,
		y: dataY,
	}
}
