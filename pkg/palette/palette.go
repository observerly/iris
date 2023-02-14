package palette

import (
	"fmt"
	"sync"

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

/*
*

	FromPalette takes a colour palette and returns the red, green and blue
	channels as a slice of float32, which can be constructed into an image
	or into FITS data.

	@param p *Palette - the colour palette to construct the channels from
	@returns []float32, []float32, []float32, error - the red, green and blue channels, and any error

*
*/
func FromPalette(p *Palette) (r, g, b []float32, err error) {
	var wg sync.WaitGroup

	var mu sync.Mutex

	wg.Add(3)

	R := make(chan []float32, 1)

	G := make(chan []float32, 1)

	B := make(chan []float32, 1)

	errors := make(chan error, 3)

	// Perform Combination of the Red Channel:
	go func() {
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()
		// Combine the red channel:
		r, err = combinePaletteChannel(p.R)

		if err != nil {
			errors <- err
		}

		R <- r
	}()

	// Perform Combination of the Green Channel:
	go func() {
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()
		// Combine the green channel:
		g, err = combinePaletteChannel(p.G)

		if err != nil {
			errors <- err
		}

		G <- g
	}()

	// Perform Combination of the Blue Channel:
	go func() {
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()
		// Combine the blue channel:
		b, err = combinePaletteChannel(p.B)

		if err != nil {
			errors <- err
		}

		B <- b
	}()

	go func() {
		wg.Wait()
		close(R)
		close(G)
		close(B)
		close(errors)
	}()

	return <-R, <-G, <-B, <-errors
}

/*
*

	combinePaletteChannel takes a slice of PaletteChannel and combines them
	into a single slice of float32.

	@param c []PaletteChannel - the slice of PaletteChannel to combine
	@returns []float32, error - the combined slice of float32, and any error

*
*/
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
