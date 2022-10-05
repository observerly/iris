package iris

import (
	"bytes"
	"image"
)

type MonochromeExposure struct {
	Width  int
	Height int
	Raw    [][]uint32
	Buffer bytes.Buffer
	Image  *image.Gray
}

func NewMonochromeExposure(exposure [][]uint32, xs int, ys int) MonochromeExposure {
	img := image.NewGray(image.Rect(0, 0, xs, ys))

	mono := MonochromeExposure{
		Width:  xs,
		Height: ys,
		Raw:    exposure,
		Buffer: bytes.Buffer{},
		Image:  img,
	}

	return mono
}
