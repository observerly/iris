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
