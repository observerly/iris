package photometry

type Star struct {
	Index     int32   // Index of the star in the data array. int32(x)+width*int32(y)
	Value     uint32  // Value of the star in the data array. data[index]
	X         float32 // Precise star x position
	Y         float32 // Precise star y position
	Intensity uint32  // Intensity of the star at position { X, Y }
	HFR       float32 // Half-Flux Radius of the star, in pixels
}

type StarsExtractor struct {
	Width     int
	Height    int
	Threshold float32
	Radius    float32
	Raw       []uint32
	Stars     []Star
	HRF       float32
}

func NewStarsExtractor(exposure [][]uint32, xs int, ys int, radius float32) *StarsExtractor {
	// Locate the brightest pixels in the data array above the threshold and return them as stars:
	var raw []uint32

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, a := range exposure {
		raw = append(raw, a...)
	}

	stars := make([]Star, 0)

	return &StarsExtractor{
		Width:     xs,
		Height:    ys,
		Threshold: 0,
		Radius:    radius,
		Raw:       raw,
		Stars:     stars,
	}
}
