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

func (m *Monochrome16Exposure) GetOtsuThresholdValue(img *image.Gray16, size image.Point, histogram [65535]uint64) uint16 {
	var threshold uint16

	var sumHistogram float64
	var sumBackground float64
	var weightBackground int
	var weightForeground int

	pixels := size.X * size.Y

	maxVariance := 0.0

	for i, bin := range histogram {
		weightBackground += int(bin)

		if weightBackground == 0 {
			continue
		}

		weightForeground = pixels - weightBackground

		if weightForeground == 0 {
			break
		}

		sumBackground += float64(uint64(i) * bin)

		meanBackground := float64(sumBackground) / float64(weightBackground)
		meanForeground := (sumHistogram - sumBackground) / float64(weightForeground)

		variance := float64(weightBackground) * float64(weightForeground) * (meanBackground - meanForeground) * (meanBackground - meanForeground)

		if variance > maxVariance {
			maxVariance = variance
			threshold = uint16(i)
		}
	}

	return threshold
}
