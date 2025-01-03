/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris
//	@license	Copyright © 2021-2025 observerly

/*****************************************************************************************************************/

package iris

/*****************************************************************************************************************/

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"

	"github.com/observerly/iris/pkg/fits"
	"github.com/observerly/iris/pkg/histogram"
	"github.com/observerly/iris/pkg/photometry"
)

/*****************************************************************************************************************/

type MonochromeExposure struct {
	Width     int
	Height    int
	Raw       [][]uint32
	Data      []float32 // The underlying data array
	Processed [][]uint32
	ADU       int32
	Buffer    bytes.Buffer
	Image     *image.Gray
	Otsu      *image.Gray
	Noise     float64
	Threshold uint8
	Pixels    int
}

/*****************************************************************************************************************/

func NewMonochromeExposure(exposure [][]uint32, adu int32, xs int, ys int) MonochromeExposure {
	pixels := xs * ys

	img := image.NewGray(image.Rect(0, 0, xs, ys))

	mono := MonochromeExposure{
		Width:     xs,
		Height:    ys,
		Raw:       exposure,
		Data:      make([]float32, pixels),
		Processed: exposure,
		ADU:       adu,
		Buffer:    bytes.Buffer{},
		Image:     img,
		Pixels:    pixels,
	}

	return mono
}

/*****************************************************************************************************************/

func (m *MonochromeExposure) GetBuffer(img *image.Gray) (bytes.Buffer, error) {
	var buff bytes.Buffer

	err := jpeg.Encode(&buff, img, &jpeg.Options{Quality: 100})

	if err != nil {
		return buff, err
	}

	return buff, nil
}

/*****************************************************************************************************************/

func (m *MonochromeExposure) GetFITSImage() *fits.FITSImage {
	f := fits.NewFITSImageFrom2DData(
		m.Raw,
		2,
		int32(m.Width),
		int32(m.Height),
		m.ADU,
	)

	f.Header.Strings["SENSOR"] = struct {
		Value   string
		Comment string
	}{
		Value:   "Monochrome",
		Comment: "ASCOM Alpaca Sensor Type",
	}

	return f
}

/*****************************************************************************************************************/

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

/*****************************************************************************************************************/

// Preprocesses an ASCOM Alpaca Image Array to a m.Raw 2D array of uint32 values. Converts the 2D array of
// uint16 values to a 2D array of uint32 values, returning a bytes.Buffer containing the preprocessed image.
func (m *MonochromeExposure) PreprocessImageArray(xs int, ys int) (bytes.Buffer, error) {
	// Switch the columns and rows in the image:
	ex := make([][]uint32, xs)

	for y := 0; y < ys; y++ {
		row := make([]uint32, xs)
		ex[y] = row
	}

	for i := 0; i < xs; i++ {
		for j := 0; j < ys; j++ {
			ex[j][i] = m.Raw[i][j]
		}
	}

	m.Raw = ex

	m.Processed = ex

	data := make([]float32, 0)

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range ex {
		for _, col := range row {
			data = append(data, float32(col))
		}
	}

	m.Data = data

	return m.Preprocess()
}

/*****************************************************************************************************************/

func (m *MonochromeExposure) Preprocess() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	gray := image.NewGray(bounds)

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			gray.SetGray(i, j, color.Gray{uint8(m.Raw[j][i])})
		}
	}

	m.Image = gray

	return m.GetBuffer(m.Image)
}

/*****************************************************************************************************************/

func (m *MonochromeExposure) ApplyNoiseReduction() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	gray := image.NewGray(bounds)

	noise := photometry.NewNoiseExtractor(m.Data, m.Width, m.Height)

	m.Noise = noise.GetGaussianNoise()

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			pixel := m.Raw[j][i]

			if pixel < uint32(m.Noise) {
				gray.SetGray(i, j, color.Gray{Y: 0})
				m.Processed[j][i] = 0
			} else {
				reduction := pixel - uint32(m.Noise)
				gray.SetGray(i, j, color.Gray{uint8(reduction)})
				m.Processed[j][i] = reduction
			}
		}
	}

	m.Image = gray

	return m.GetBuffer(m.Image)
}

/*****************************************************************************************************************/

func (m *MonochromeExposure) ApplyOtsuThreshold() (bytes.Buffer, error) {
	bounds := m.Image.Bounds()

	size := bounds.Size()

	// Get the Otsu Method's threshold value for our image:
	m.Threshold = m.GetOtsuThresholdValue(m.Image, size, histogram.HistogramGray(m.Image))

	gray := image.NewGray(bounds)

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			pixel := m.Raw[j][i]

			if pixel < uint32(m.Threshold) {
				gray.SetGray(i, j, color.Gray{Y: 0})
				m.Processed[j][i] = 0
			} else {
				gray.SetGray(i, j, color.Gray{Y: uint8(pixel)})
				m.Processed[j][i] = pixel
			}
		}
	}

	m.Otsu = gray

	return m.GetBuffer(m.Otsu)
}

/*****************************************************************************************************************/
