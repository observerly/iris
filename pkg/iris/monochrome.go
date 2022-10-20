package iris

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
)

type MonochromeExposure struct {
	Width  int
	Height int
	Raw    [][]uint32
	Buffer bytes.Buffer
	Image  *image.Gray
	Pixels int
}

func NewMonochromeExposure(exposure [][]uint32, xs int, ys int) MonochromeExposure {
	img := image.NewGray(image.Rect(0, 0, xs, ys))

	mono := MonochromeExposure{
		Width:  xs,
		Height: ys,
		Raw:    exposure,
		Buffer: bytes.Buffer{},
		Image:  img,
		Pixels: xs * ys,
	}

	return mono
}

func (m *MonochromeExposure) GetBuffer(img *image.Gray) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: 100})

	if err != nil {
		return buff, err
	}

	return buff, nil
}

func (m *MonochromeExposure) Preprocess() (bytes.Buffer, error) {
	for j := 0; j < m.Height; j++ {
		for i := 0; i < m.Width; i++ {
			m.Image.SetGray(i, j, color.Gray{uint8(m.Raw[j][i])})
		}
	}

	return m.GetBuffer(m.Image)
}
