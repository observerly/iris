package iris

import (
	"bytes"
	"fmt"
	"image"
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

func NewRGGBExposure(exposure [][]uint32, xs int, ys int, cfa string) RGGBExposure {
	img := image.NewRGBA(image.Rect(0, 0, xs, ys))

	rggb := RGGBExposure{
		Width:             xs,
		Height:            ys,
		ColourFilterArray: cfa,
		Raw:               exposure,
		Buffer:            bytes.Buffer{},
		Image:             img,
	}

	return rggb
}

/**
	Accepts a CFA (Color Filter Array) string, e.g., "RGGB" and returns the Bayering Matrix offset
**/
func (b *RGGBExposure) GetBayerMatrixOffset() (xOffset int32, yOffset int32, err error) {
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
