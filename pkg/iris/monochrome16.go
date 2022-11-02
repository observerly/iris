package iris

import (
	"bytes"
	"image"
	"image/jpeg"
)

type Monochrome16Exposure struct {
	Width     int
	Height    int
	Raw       [][]uint32
	ADU       int32
	Buffer    bytes.Buffer
	Image     *image.Gray16
	Otsu      *image.Gray16
	Noise     float64
	Threshold uint16
	Pixels    int
}

func NewMonochrome16Exposure(exposure [][]uint32, adu int32, xs int, ys int) Monochrome16Exposure {
	img := image.NewGray16(image.Rect(0, 0, xs, ys))

	mono := Monochrome16Exposure{
		Width:  xs,
		Height: ys,
		Raw:    exposure,
		ADU:    adu,
		Buffer: bytes.Buffer{},
		Image:  img,
		Pixels: xs * ys,
	}

	return mono
}

func (m *Monochrome16Exposure) GetBuffer(img *image.Gray16) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: 100})

	if err != nil {
		return buff, err
	}

	return buff, nil
}
