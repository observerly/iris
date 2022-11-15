package iris

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

type RGGB64Exposure struct {
	Width             int
	Height            int
	Raw               [][]uint32
	R                 []float32
	G                 []float32
	B                 []float32
	ADU               int32
	Buffer            bytes.Buffer
	Image             *image.RGBA64
	ColourFilterArray string
	Pixels            int
}

type RGGB64Color struct {
	Name    string
	Channel []float32
}

func NewRGGB64Exposure(exposure [][]uint32, adu int32, xs int, ys int, cfa string) *RGGB64Exposure {
	img := image.NewRGBA64(image.Rect(0, 0, xs, ys))

	return &RGGB64Exposure{
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

/**
	Accepts a CFA (Color Filter Array) string, e.g., "RGGB" and returns the Bayering Matrix offset
**/
func (b *RGGB64Exposure) GetBayerMatrixOffset() (xOffset int, yOffset int, err error) {
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

func (b *RGGB64Exposure) GetBuffer(img *image.RGBA64) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: 100})

	if err != nil {
		return buff, err
	}

	return buff, nil
}

/**
	Convert an R or G or B channel to a FITS standard image
**/
func (b *RGGB64Exposure) GetFITSImageForChannel(color RGGB64Color) *fits.FITSImage {
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
		Comment: "RGB 16 bit Channel",
	}

	return f
}

/**
	Return each R, G, B FITS standard images
**/
func (b *RGGB64Exposure) GetFITSImages() (*fits.FITSImage, *fits.FITSImage, *fits.FITSImage) {
	var wg sync.WaitGroup

	wg.Add(3)

	R := make(chan *fits.FITSImage, 1)

	G := make(chan *fits.FITSImage, 1)

	B := make(chan *fits.FITSImage, 1)

	go func() {
		defer wg.Done()
		// Get the FITS image for the R channel:
		red := b.GetFITSImageForChannel(RGGB64Color{
			Name:    "Red",
			Channel: b.R,
		})
		R <- red
	}()

	go func() {
		defer wg.Done()
		// Get the FITS image for the G channel:
		green := b.GetFITSImageForChannel(RGGB64Color{
			Name:    "Green",
			Channel: b.G,
		})
		G <- green
	}()

	go func() {
		defer wg.Done()
		// Get the FITS image for the B channel:
		blue := b.GetFITSImageForChannel(RGGB64Color{
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

/**
	Perform Debayering w/ Bilinear Interpolation Technique
**/
func (b *RGGB64Exposure) DebayerBilinearInterpolation() error {
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
			b.Image.Set(i, j, color.RGBA64{
				R: uint16(red[j*b.Width+i]),
				G: uint16(green[j*b.Width+i]),
				B: uint16(blue[j*b.Width+i]),
				A: 255,
			})
		}
	}

	return nil
}

/*
 	PreprocessImageArray()

	Preprocesses an ASCOM Alpaca Image Array to a m.Raw 2D array of uint32 values.
	Converts the 2D array of uint16 values to a 2D array of uint32 values.

	@returns  a bytes.Buffer containing the preprocessed image.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__imagearray

	"... "column-major" order (column changes most rapidly) from the image's row and column
	perspective, while, from the array's perspective, serialisation is actually effected in
	"row-major" order (rightmost index changes most rapidly). This unintuitive outcome arises
	because the ASCOM Camera Interface specification defines the image column dimension as
	the rightmost array dimension."

*/
func (b *RGGB64Exposure) PreprocessImageArray(xs int, ys int) (bytes.Buffer, error) {
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

func (b *RGGB64Exposure) Preprocess() (bytes.Buffer, error) {
	// Perform Debayering w/ Bilinear Interpolation Technique:
	err := b.DebayerBilinearInterpolation()

	if err != nil {
		return b.Buffer, err
	}

	return b.GetBuffer(b.Image)
}
