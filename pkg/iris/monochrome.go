package iris

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"sync"
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

func (m *MonochromeExposure) Preprocess() (bytes.Buffer, error) {
	var wg sync.WaitGroup

	wg.Add(m.Width * m.Height)

	for i := 0; i < m.Width; i++ {
		for j := 0; j < m.Height; j++ {
			go func(i, j int) {
				m.Image.SetGray(i, j, color.Gray{uint8(m.Raw[i][j])})
				wg.Done()
			}(i, j)
		}
	}

	wg.Wait()

	var buff bytes.Buffer

	err := jpeg.Encode(&buff, m.Image, &jpeg.Options{Quality: 100})

	if err != nil {
		return buff, err
	}

	return buff, nil
}
