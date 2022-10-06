package iris

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"strings"
)

type RGGBExposure struct {
	Width             int
	Height            int
	ColourFilterArray string
	Raw               [][]uint32
	Buffer            bytes.Buffer
	Image             *image.RGBA
}

func NewRGGBExposure(exposure [][]uint32, xs int, ys int, cfa string) *RGGBExposure {
	img := image.NewRGBA(image.Rect(0, 0, xs, ys))

	return &RGGBExposure{
		Width:             xs,
		Height:            ys,
		ColourFilterArray: cfa,
		Raw:               exposure,
		Buffer:            bytes.Buffer{},
		Image:             img,
	}
}

/**
	Accepts a CFA (Color Filter Array) string, e.g., "RGGB" and returns the Bayering Matrix offset
**/
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

/**
	Perform Debayering w/ Bilinear Interpolation Technique
**/
func (b *RGGBExposure) DebayerBilinearInterpolation(xOffset int, yOffset int) error {
	var raw []uint32

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, a := range b.Raw {
		raw = append(raw, a...)
	}

	w := uint32(b.Width)

	h := uint32(b.Height)

	xo := uint32(xOffset)

	yo := uint32(yOffset)

	// We need to ensure our images are of even pixel dimensions:
	// Effectively, we're ignoring the last row and column of pixels in odd sized images:
	x := w - xo & ^uint32(1)

	y := h - yo & ^uint32(1)

	R := make([]float32, int(x)*int(y))

	G := make([]float32, int(x)*int(y))

	B := make([]float32, int(x)*int(y))

	// Perform Bi-Linear Interpolation on the Colour Filter Array:
	for i := uint32(0); i < y; i += 2 {
		for j := uint32(0); j < x; j += 2 {
			// Obtain a Convolution in the Red channel:
			R = BiLinearConvolveRedChannel(i, j, raw, R, w, h, xo, yo, x, y)
			// Obtain a Convolution in the Green channel:
			G = BiLinearConvolveGreenChannel(i, j, raw, G, w, h, xo, yo, x, y)
			// Obtain a Convolution in the Blue channel:
			B = BiLinearConvolveBlueChannel(i, j, raw, B, w, h, xo, yo, x, y)
		}
	}

	// Stack The RGB channels into a single image:
	for i := 0; i < b.Height; i++ {
		for j := 0; j < b.Width; j++ {
			b.Image.Set(i, j, color.RGBA{
				R: uint8(R[i*b.Height+j]),
				G: uint8(G[i*b.Height+j]),
				B: uint8(B[i*b.Height+j]),
				A: 255,
			})
		}
	}

	return nil
}

func (b *RGGBExposure) Preprocess() (bytes.Buffer, error) {
	xOffset, yOffset, err := b.GetBayerMatrixOffset()

	if err != nil {
		return b.Buffer, err
	}

	err = b.DebayerBilinearInterpolation(xOffset, yOffset)

	if err != nil {
		return b.Buffer, err
	}

	// Encode the image as a JPEG:
	err = jpeg.Encode(&b.Buffer, b.Image, &jpeg.Options{Quality: 100})

	if err != nil {
		return b.Buffer, err
	}

	return b.Buffer, nil
}
