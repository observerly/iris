package palette

type PaletteChannel struct {
	Data     []float32
	Fraction float32
}

type Palette struct {
	Name string
	R    []PaletteChannel
	G    []PaletteChannel
	B    []PaletteChannel
}
