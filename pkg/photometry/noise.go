package photometry

type NoiseExtractor struct {
	Width  int
	Height int
	Noise  float64
	Raw    []uint32
}

func NewNoiseExtractor(exposure [][]uint32, xs int, ys int) *NoiseExtractor {
	// Locate the brightest pixels in the data array above the threshold and return them as stars:
	var raw []uint32

	// Flatten the 2D Exposure Array array into a 1D array:
	for _, a := range exposure {
		raw = append(raw, a...)
	}

	return &NoiseExtractor{
		Width:  xs,
		Height: ys,
		Noise:  0,
		Raw:    raw,
	}
}
