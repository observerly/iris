package photometry

import (
	"encoding/json"
	"image"
	"image/color"
	"os"
	"testing"
)

func TestBiLinearConvolveRedChannel(t *testing.T) {
	type CameraExposure struct {
		BayerXOffset int32      `json:"bayerXOffset"`
		BayerYOffset int32      `json:"bayerYOffset"`
		CCDXSize     int32      `json:"ccdXSize"`
		CCDYSize     int32      `json:"ccdYSize"`
		Image        [][]uint32 `json:"exposure"`
		MaxADU       int32      `json:"maxADU"`
		Rank         uint32     `json:"rank"`
		SensorType   string     `json:"sensorType"`
	}

	file, err := os.ReadFile("../../data/m42-800x600-rggb.json")

	if err != nil {
		t.Errorf("Error opening from JSON data: %s", err)
	}

	ex := CameraExposure{}

	_ = json.Unmarshal([]byte(file), &ex)

	var raw []uint32

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range ex.Image {
		raw = append(raw, row...)
	}

	w := uint32(1200)

	h := uint32(800)

	xo := uint32(0)

	yo := uint32(0)

	x := w - xo & ^uint32(1)

	y := h - yo & ^uint32(1)

	red := BiLinearConvolveRedChannel(raw, w, h, xo, yo, x, y)

	img := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))

	// Stack The Red channel into a single image:
	for j := 0; j < int(h); j++ {
		for i := 0; i < int(w); i++ {
			img.Set(i, j, color.RGBA{
				R: uint8(red[j*int(w)+i]),
				G: 0,
				B: 0,
				A: 255,
			})
		}
	}
}

func TestBiLinearConvolveGreenChannel(t *testing.T) {
	type CameraExposure struct {
		BayerXOffset int32      `json:"bayerXOffset"`
		BayerYOffset int32      `json:"bayerYOffset"`
		CCDXSize     int32      `json:"ccdXSize"`
		CCDYSize     int32      `json:"ccdYSize"`
		Image        [][]uint32 `json:"exposure"`
		MaxADU       int32      `json:"maxADU"`
		Rank         uint32     `json:"rank"`
		SensorType   string     `json:"sensorType"`
	}

	file, err := os.ReadFile("../../data/m42-800x600-rggb.json")

	if err != nil {
		t.Errorf("Error opening from JSON data: %s", err)
	}

	ex := CameraExposure{}

	_ = json.Unmarshal([]byte(file), &ex)

	var raw []uint32

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range ex.Image {
		raw = append(raw, row...)
	}

	w := uint32(1200)

	h := uint32(800)

	xo := uint32(0)

	yo := uint32(0)

	x := w - xo & ^uint32(1)

	y := h - yo & ^uint32(1)

	green := BiLinearConvolveGreenChannel(raw, w, h, xo, yo, x, y)

	img := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))

	// Stack The Green channel into a single image:
	for j := 0; j < int(h); j++ {
		for i := 0; i < int(w); i++ {
			img.Set(i, j, color.RGBA{
				R: 0,
				G: uint8(green[j*int(w)+i]),
				B: 0,
				A: 255,
			})
		}
	}
}

func TestBiLinearConvolveBlueChannel(t *testing.T) {
	type CameraExposure struct {
		BayerXOffset int32      `json:"bayerXOffset"`
		BayerYOffset int32      `json:"bayerYOffset"`
		CCDXSize     int32      `json:"ccdXSize"`
		CCDYSize     int32      `json:"ccdYSize"`
		Image        [][]uint32 `json:"exposure"`
		MaxADU       int32      `json:"maxADU"`
		Rank         uint32     `json:"rank"`
		SensorType   string     `json:"sensorType"`
	}

	file, err := os.ReadFile("../../data/m42-800x600-rggb.json")

	if err != nil {
		t.Errorf("Error opening from JSON data: %s", err)
	}

	ex := CameraExposure{}

	_ = json.Unmarshal([]byte(file), &ex)

	var raw []uint32

	// Flatten the 2D Colour Filter Array array into a 1D array:
	for _, row := range ex.Image {
		raw = append(raw, row...)
	}

	w := uint32(1200)

	h := uint32(800)

	xo := uint32(0)

	yo := uint32(0)

	x := w - xo & ^uint32(1)

	y := h - yo & ^uint32(1)

	blue := BiLinearConvolveBlueChannel(raw, w, h, xo, yo, x, y)

	img := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))

	// Stack The Blue channel into a single image:
	for j := 0; j < int(h); j++ {
		for i := 0; i < int(w); i++ {
			img.Set(i, j, color.RGBA{
				R: 0,
				G: 0,
				B: uint8(blue[j*int(w)+i]),
				A: 255,
			})
		}
	}
}
