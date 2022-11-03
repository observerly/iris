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

func (m *Monochrome16Exposure) GetFITSImage() *fits.FITSImage {
	f := fits.NewFITSImageFrom2DData(
		m.Raw,
		16,
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

func (m *Monochrome16Exposure) Preprocess() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	size := bounds.Size()

	gray := image.NewGray16(bounds)

	setPixel := func(gray *image.Gray16, x int, y int) {
		gray.SetGray16(x, y, color.Gray16{uint16(m.Raw[x][y])})
	}

	utils.DeferForEachPixel(size, func(x, y int) {
		setPixel(gray, x, y)
	})

	m.Image = gray

	return m.GetBuffer(m.Image)
}

func (m *Monochrome16Exposure) ApplyNoiseReduction() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	size := bounds.Size()

	gray := image.NewGray16(bounds)

	noise := photometry.NewNoiseExtractor(m.Raw, m.Width, m.Height)

	m.Noise = noise.GetGaussianNoise()

	setPixel := func(gray *image.Gray16, x int, y int) {
		pixel := m.Raw[x][y]

		if pixel < uint32(m.Noise) {
			gray.SetGray16(x, y, color.Gray16{Y: 0})
		} else {
			gray.SetGray16(x, y, color.Gray16{uint16(pixel - uint32(m.Noise))})
		}
	}

	utils.DeferForEachPixel(size, func(x, y int) {
		setPixel(gray, x, y)
	})

	m.Image = gray

	return m.GetBuffer(m.Image)
}

func (m *Monochrome16Exposure) ApplyOtsuThreshold() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	size := bounds.Size()

	// Get the Otsu Method's threshold value for our image:
	m.Threshold = m.GetOtsuThresholdValue(m.Image, size, histogram.HistogramGray16(m.Image))

	gray := image.NewGray16(bounds)

	setPixel := func(gray *image.Gray16, x int, y int) {
		pixel := m.Image.Gray16At(x, y).Y

		if pixel < m.Threshold {
			gray.SetGray16(x, y, color.Gray16{Y: 0})
		} else {
			gray.SetGray16(x, y, color.Gray16{Y: pixel})
		}
	}

	utils.DeferForEachPixel(size, func(x, y int) {
		setPixel(gray, x, y)
	})

	m.Otsu = gray

	return m.GetBuffer(m.Otsu)
}
