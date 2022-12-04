package palette

import (
	"fmt"

	"github.com/observerly/iris/pkg/utils"
)

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

/**
	combinePaletteChannel takes a slice of PaletteChannel and combines them
	into a single slice of float32.

	@param c []PaletteChannel - the slice of PaletteChannel to combine
	@returns []float32, error - the combined slice of float32, and any error
**/
func combinePaletteChannel(channel []PaletteChannel) ([]float32, error) {
	// Take each channel of the palette, and their respective constituents, and multiply them by the fraction:
	fraction := float32(0.0)

	data := make([][]float32, 0)

	for _, c := range channel {
		for k, v := range c.Data {
			// Return an error to say that the fraction can not be larger than 1.0 for a given channel:
			if c.Fraction > 1.0 {
				return nil, fmt.Errorf("the fraction for the red channel can not be larger than 1.0")
			}

			c.Data[k] = v * c.Fraction
		}

		// Add the data to the reds slice:
		data = append(data, c.Data)

		fraction += c.Fraction
	}

	// Return an error to say that the combined fraction can not be larger than 1.0 for a given channel:
	if fraction > 1.0 {
		return nil, fmt.Errorf("the fraction for the red channel can not be larger than 1.0")
	}

	d, err := utils.MeanFloat32Arrays(data)

	if err != nil {
		return nil, err
	}

	return d, nil
}
