package vcurve

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
