/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/vcurve
//	@license	Copyright © 2021-2025 observerly

/*****************************************************************************************************************/

package vcurve

/*****************************************************************************************************************/

import (
	"math"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/stat"
)

/*****************************************************************************************************************/

// Point is a data point with x and y coordinates.
type Point struct {
	X float64
	Y float64
}

/*****************************************************************************************************************/

// VCurve is a struct that holds the data points for the V-curve.
type VCurve struct {
	Points []Point
}

/*****************************************************************************************************************/

// VCurveParams is a struct that holds the parameters for the V-curve model optisation and the data points.
type VCurveParams struct {
	A float64
	B float64
	C float64
	D float64
	X []float64
	Y []float64
}

/*****************************************************************************************************************/

// NewHyperbolicVCurve
//
// Creates a new VCurve object ready for applying the Levenberg-Marquardt iterative optimization technique.
// The VCurve object is initialized with the data points, and initial guesses for the parameters are calculated
// from the input data.
func NewHyperbolicVCurve(data VCurve) *VCurveParams {
	// Preallocate slices with the exact required capacity
	dataX := make([]float64, 0, len(data.Points))
	dataY := make([]float64, 0, len(data.Points))

	// A single loop to populate the slices
	for _, point := range data.Points {
		dataX = append(dataX, point.X)
		dataY = append(dataY, point.Y)
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
		X: dataX,
		Y: dataY,
	}
}

/*****************************************************************************************************************/

// This is the hyperbolic function that we want to fit to the data.
func hyperbolicFunction(x float64, params VCurveParams) float64 {
	a, b, c, d := params.A, params.B, params.C, params.D
	return b*math.Sqrt(1+math.Pow((x-c)/a, 2)) + d
}

/*****************************************************************************************************************/

// objectiveFunc is the least squares objective function that accepts dataX and dataY.
func objectiveFunc(dataX, dataY []float64) func(params []float64) float64 {
	// If we do not have the same number of x and y data points, we cannot fit the model.
	if len(dataX) != len(dataY) {
		panic("x and y data points must be the same length")
	}

	// If we do not have at least 1 data point, we cannot fit the model.
	if len(dataX) < 1 {
		panic("x data points must have at least 1 data point")
	}

	// Return the objective function.
	return func(params []float64) float64 {
		var sumSq float64
		for i := range dataX {
			yPredicted := hyperbolicFunction(dataX[i], VCurveParams{
				A: params[0],
				B: params[1],
				C: params[2],
				D: params[3],
			})
			sumSq += math.Pow(dataY[i]-yPredicted, 2)
		}
		return sumSq
	}
}

/*****************************************************************************************************************/

// LevenbergMarquardtOptimisation
// Optimizes the hyperbolic function using the Levenberg-Marquardt algorithm.
func (p *VCurveParams) LevenbergMarquardtOptimisation() (VCurveParams, error) {
	// Setting up the optimizer:
	problem := optimize.Problem{
		Func: objectiveFunc(p.X, p.Y),
	}

	// Create custom settings for the optimization:
	settings := &optimize.Settings{
		GradientThreshold: 1e-6, // Customize as needed
		// Add other settings as needed
	}

	// Initial guess for parameters:
	initialParams := []float64{p.A, p.B, p.C, p.D}

	// Run the optimization:
	result, err := optimize.Minimize(problem, initialParams, settings, nil)

	if err != nil {
		return VCurveParams{}, err
	}

	// Return Optimized parameters
	return VCurveParams{
		A: result.X[0],
		B: result.X[1],
		C: result.X[2],
		D: result.X[3],
	}, nil
}

/*****************************************************************************************************************/
