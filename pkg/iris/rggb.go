/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris
//	@license	Copyright © 2021-2025 observerly

/*****************************************************************************************************************/

package iris

/*****************************************************************************************************************/

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"strings"
	"sync"

	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/photometry"
)

/*****************************************************************************************************************/

type RGGBExposure struct {
	Width             int
	Height            int
	Raw               [][]uint32
	R                 []float32
	G                 []float32
	B                 []float32
	ADU               int32
	Buffer            bytes.Buffer
	Image             *image.RGBA
	ColourFilterArray string
	Pixels            int
}

/*****************************************************************************************************************/

type RGGBColor struct {
	Name    string
	Channel []float32
}

/*****************************************************************************************************************/

func NewRGGBExposure(exposure [][]uint32, adu int32, xs int, ys int, cfa string) *RGGBExposure {
	img := image.NewRGBA(image.Rect(0, 0, xs, ys))

	return &RGGBExposure{
		Width:             xs,
		Height:            ys,
		Raw:               exposure,
		ADU:               adu,
		Buffer:            bytes.Buffer{},
		Image:             img,
		ColourFilterArray: cfa,
		Pixels:            xs * ys,
	}
}

/*****************************************************************************************************************/

// Accepts a CFA (Color Filter Array) string, e.g., "RGGB" and returns the Bayering Matrix offset
func (b *RGGBExposure) GetBayerMatrixOffset() (xOffset int, yOffset int, err error) {
	switch strings.ToLower(b.ColourFilterArray) {
	case "rggb":
		return 0, 0, nil
	case "grbg":
		return 1, 0, nil
	case "gbrg":
		return 0, 1, nil
	case "bggr":
		return 1, 1, nil
	default:
		return 0, 0, fmt.Errorf("unknown color filter array string: %v", b.ColourFilterArray)
	}
}

/*****************************************************************************************************************/

func (b *RGGBExposure) GetBuffer(img *image.RGBA) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: 100})

	if err != nil {
		return buff, err
	}

	return buff, nil
}

/*****************************************************************************************************************/

// Converts an R, G, or B channel to a Monochrome 16 bit exposure
func (b *RGGBExposure) GetMonochrome() MonochromeExposure {
	// Create a 2D array of the specific RGB channel from flattened 1D color channel array:
	raw := make([][]uint32, b.Height)

	for j := 0; j < b.Height; j++ {
		row := make([]uint32, b.Width)
		for i := 0; i < b.Width; i++ {
			// Destination Offset:
			do := j*b.Width + i

			// The RGB to Monochrome natural luminance component Y (from the CIE XYZ system)
			// captures what is most perceived by humans as color in one channel.
			lum := 0.299*float64(b.R[do]) + 0.587*float64(b.G[do]) + 0.114*float64(b.B[do])

			row[i] = uint32(lum)
		}

		raw[j] = row
	}

	m := NewMonochromeExposure(
		raw,
		b.ADU,
		b.Width,
		b.Height,
	)

	return m
}

/*****************************************************************************************************************/

// Converts a R, G, or B channel to a FITS standard image
func (b *RGGBExposure) GetFITSImageForChannel(color RGGBColor) *fits.FITSImage {
	// Create a 2D array of the specific RGB channel from flattened 1D color channel array:
	raw := make([][]uint32, b.Height)

	for j := 0; j < b.Height; j++ {
		row := make([]uint32, b.Width)
		for i := 0; i < b.Width; i++ {
			row[i] = uint32(color.Channel[j*b.Width+i])
		}
		raw[j] = row
	}

	f := fits.NewFITSImageFrom2DData(
		raw,
		2,
		int32(b.Width),
		int32(b.Height),
		b.ADU,
	)

	f.Header.Strings["SENSOR"] = struct {
		Value   string
		Comment string
	}{
		Value:   "RGGB",
		Comment: "ASCOM Alpaca Sensor Type",
	}

	f.Header.Strings["CHANNEL"] = struct {
		Value   string
		Comment string
	}{
		Value:   color.Name,
		Comment: "RGB Channel",
	}

	return f
}

/*****************************************************************************************************************/

// Returns each of the R, G and B channels as FITS images
func (b *RGGBExposure) GetFITSImages() (*fits.FITSImage, *fits.FITSImage, *fits.FITSImage) {
	var wg sync.WaitGroup

	wg.Add(3)

	R := make(chan *fits.FITSImage, 1)

	G := make(chan *fits.FITSImage, 1)

	B := make(chan *fits.FITSImage, 1)

	go func() {
		defer wg.Done()
		// Get the FITS image for the R channel:
		red := b.GetFITSImageForChannel(RGGBColor{
			Name:    "Red",
			Channel: b.R,
		})
		R <- red
	}()

	go func() {
		defer wg.Done()
		// Get the FITS image for the G channel:
		green := b.GetFITSImageForChannel(RGGBColor{
			Name:    "Green",
			Channel: b.G,
		})
		G <- green
	}()

	go func() {
		defer wg.Done()
		// Get the FITS image for the B channel:
		blue := b.GetFITSImageForChannel(RGGBColor{
			Name:    "Blue",
			Channel: b.B,
		})
		B <- blue
	}()

	go func() {
		wg.Wait()
		close(R)
		close(G)
		close(B)
	}()

	return <-R, <-G, <-B
}

/*****************************************************************************************************************/

// Performs a Debayering with a bilinear interpolation technique.
func (b *RGGBExposure) DebayerBilinearInterpolation() error {
	var wg sync.WaitGroup

	wg.Add(3)

	R := make(chan []float32, b.Pixels)

	G := make(chan []float32, b.Pixels)

	B := make(chan []float32, b.Pixels)

	errors := make(chan error, 3)

	var raw []uint32

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range b.Raw {
		raw = append(raw, row...)
	}

	w := uint32(b.Width)

	h := uint32(b.Height)

	xOffset, yOffset, err := b.GetBayerMatrixOffset()

	if err != nil {
		return err
	}

	xo := uint32(xOffset)

	yo := uint32(yOffset)

	// We need to ensure our images are of even pixel dimensions:
	// Effectively, we're ignoring the last row and column of pixels in odd sized images:
	x := w - xo & ^uint32(1)

	y := h - yo & ^uint32(1)

	// Perform Bi-Linear Interpolation on the Colour Filter Array:
	go func() {
		defer wg.Done()
		// Obtain a Convolution in the Red channel:
		red := photometry.BiLinearConvolveRedChannel(raw, w, h, xo, yo, x, y)
		R <- red
	}()

	go func() {
		defer wg.Done()
		// Obtain a Convolution in the Green channel:
		green := photometry.BiLinearConvolveGreenChannel(raw, w, h, xo, yo, x, y)
		G <- green
	}()

	go func() {
		defer wg.Done()
		// Obtain a Convolution in the Blue channel:
		blue := photometry.BiLinearConvolveBlueChannel(raw, w, h, xo, yo, x, y)
		B <- blue
	}()

	go func() {
		wg.Wait()
		close(R)
		close(G)
		close(B)
		close(errors)
	}()

	red, green, blue := <-R, <-G, <-B

	b.R = red

	b.G = green

	b.B = blue

	// Stack The RGB channels into a single image:
	for j := 0; j < b.Height; j++ {
		for i := 0; i < b.Width; i++ {
			b.Image.Set(i, j, color.RGBA{
				R: uint8(red[j*b.Width+i]),
				G: uint8(green[j*b.Width+i]),
				B: uint8(blue[j*b.Width+i]),
				A: 255,
			})
		}
	}

	return nil
}

/*****************************************************************************************************************/

// Preprocesses an ASCOM Alpaca Image Array to a m.Raw 2D array of uint32 values. Converts the 2D array of uint16
// values to a 2D array of uint32 values, returning a bytes.Buffer containing the preprocessed image.
func (b *RGGBExposure) PreprocessImageArray(xs int, ys int) (bytes.Buffer, error) {
	// Switch the columns and rows in the image:
	ex := make([][]uint32, xs)

	for y := 0; y < ys; y++ {
		row := make([]uint32, xs)
		ex[y] = row
	}

	for i := 0; i < xs; i++ {
		for j := 0; j < ys; j++ {
			ex[j][i] = b.Raw[i][j]
		}
	}

	b.Raw = ex

	return b.Preprocess()
}

func (b *RGGBExposure) Preprocess() (bytes.Buffer, error) {
	// Perform Debayering w/ Bilinear Interpolation Technique:
	err := b.DebayerBilinearInterpolation()

	if err != nil {
		return b.Buffer, err
	}

	return b.GetBuffer(b.Image)
}

/*****************************************************************************************************************/
