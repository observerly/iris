package iris

import (
	"bytes"
	"image"
	"image/jpeg"
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

func (b *RGGB64Exposure) GetBuffer(img *image.RGBA64) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: 100})

	if err != nil {
		return buff, err
	}

	return buff, nil
}
