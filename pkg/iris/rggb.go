package iris

import (
	"bytes"
	"image"
)

type RGGBExposure struct {
	Width  int
	Height int
	Raw    [][]uint32
	Buffer bytes.Buffer
	Image  *image.RGBA
}

func NewRGGBExposure(exposure [][]uint32, xs int, ys int) RGGBExposure {
	img := image.NewRGBA(image.Rect(0, 0, xs, ys))

	rggb := RGGBExposure{
		Width:  xs,
		Height: ys,
		Raw:    exposure,
		Buffer: bytes.Buffer{},
		Image:  img,
	}

	return rggb
}
