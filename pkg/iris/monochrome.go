package iris

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"

	"github.com/observerly/iris/pkg/utils"
)

type MonochromeExposure struct {
	Width  int
	Height int
	Raw    [][]uint32
	Buffer bytes.Buffer
	Image  *image.Gray
	Otsu   *image.Gray
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

func (m *MonochromeExposure) ApplyOtsuThreshold(threshold uint8) (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	size := bounds.Size()

	gray := image.NewGray(bounds)

	setPixel := func(gray *image.Gray, x int, y int) {
		pixel := m.Image.GrayAt(x, y).Y

		if pixel < threshold {
			gray.SetGray(x, y, color.Gray{Y: 0})
		} else {
			gray.SetGray(x, y, color.Gray{Y: pixel})
		}
	}

	utils.DeferForEachPixel(size, func(x, y int) {
		setPixel(gray, x, y)
	})

	m.Otsu = gray

	return m.GetBuffer(m.Otsu)
}
