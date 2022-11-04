package iris

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"

	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/histogram"
	"github.com/observerly/iris/pkg/photometry"
	"github.com/observerly/iris/pkg/utils"
)

type MonochromeExposure struct {
	Width     int
	Height    int
	Raw       [][]uint32
	ADU       int32
	Buffer    bytes.Buffer
	Image     *image.Gray
	Otsu      *image.Gray
	Noise     float64
	Threshold uint8
	Pixels    int
}

func NewMonochromeExposure(exposure [][]uint32, adu int32, xs int, ys int) MonochromeExposure {
	img := image.NewGray(image.Rect(0, 0, xs, ys))

	mono := MonochromeExposure{
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

func (m *MonochromeExposure) GetBuffer(img *image.Gray) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: 100})

	if err != nil {
		return buff, err
	}

	return buff, nil
}

func (m *MonochromeExposure) GetFITSImage() *fits.FITSImage {
	f := fits.NewFITSImageFrom2DData(
		m.Raw,
		2,
		int32(m.Width),
		int32(m.Height),
	)

	f.Header.Strings["SENSOR"] = struct {
		Value   string
		Comment string
	}{
		Value:   "Monochrome",
		Comment: "",
	}

	return f
}

func (m *MonochromeExposure) GetOtsuThresholdValue(img *image.Gray, size image.Point, histogram [256]uint64) uint8 {
	var threshold uint8

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
			threshold = uint8(i)
		}
	}

	return threshold
}

func (m *MonochromeExposure) Preprocess() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	size := bounds.Size()

	gray := image.NewGray(bounds)

	setPixel := func(gray *image.Gray, x int, y int) {
		gray.SetGray(x, y, color.Gray{uint8(m.Raw[x][y])})
	}

	utils.DeferForEachPixel(size, func(x, y int) {
		setPixel(gray, x, y)
	})

	m.Image = gray

	return m.GetBuffer(m.Image)
}

func (m *MonochromeExposure) ApplyNoiseReduction() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	size := bounds.Size()

	gray := image.NewGray(bounds)

	noise := photometry.NewNoiseExtractor(m.Raw, m.Width, m.Height)

	m.Noise = noise.GetGaussianNoise()

	setPixel := func(gray *image.Gray, x int, y int) {
		pixel := m.Raw[x][y]

		if pixel < uint32(m.Noise) {
			gray.SetGray(x, y, color.Gray{Y: 0})
		} else {
			gray.SetGray(x, y, color.Gray{uint8(pixel - uint32(m.Noise))})
		}
	}

	utils.DeferForEachPixel(size, func(x, y int) {
		setPixel(gray, x, y)
	})

	m.Image = gray

	return m.GetBuffer(m.Image)
}

func (m *MonochromeExposure) ApplyOtsuThreshold() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	size := bounds.Size()

	// Get the Otsu Method's threshold value for our image:
	m.Threshold = m.GetOtsuThresholdValue(m.Image, size, histogram.HistogramGray(m.Image))

	gray := image.NewGray(bounds)

	setPixel := func(gray *image.Gray, x int, y int) {
		pixel := m.Image.GrayAt(x, y).Y

		if pixel < m.Threshold {
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
