package stats

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/observerly/iris/pkg/iris"
)

func TestCalculateMinMeanMax(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	min, mean, max := calcMinMeanMax(data)

	if min != 1 {
		t.Errorf("min should be 1, but got %v", min)
	}

	if mean != 5.5 {
		t.Errorf("mean should be 5.5, but got %v", mean)
	}

	if max != 10 {
		t.Errorf("max should be 10, but got %v", max)
	}
}

func TestCalculateMeanStdDevVar(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	mean, stddev, variance := calcMeanStdDevVar(data)

	if mean != 5.5 {
		t.Errorf("mean should be 5.5, but got %v", mean)
	}

	if stddev != 2.872281323269 {
		t.Errorf("stddev should be 2.872281323269, but got %v", stddev)
	}

	if variance != 8.25 {
		t.Errorf("variance should be 8.25, but got %v", variance)
	}
}

func TestCalculateMinMeanMaxStdDevVar(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	min, mean, max, stddev, variance := calcMinMeanMaxStdDevVar(data)

	if min != 1 {
		t.Errorf("min should be 1, but got %v", min)
	}

	if mean != 5.5 {
		t.Errorf("mean should be 5.5, but got %v", mean)
	}

	if max != 10 {
		t.Errorf("max should be 10, but got %v", max)
	}

	if stddev != 2.872281323269 {
		t.Errorf("stddev should be 2.872281323269, but got %v", stddev)
	}

	if variance != 8.25 {
		t.Errorf("variance should be 8.25, but got %v", variance)
	}
}

func TestCalculateMedianOdd(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	median := calcMedian(data)

	if median != 5.5 {
		t.Errorf("median should be 5.5, but got %v", median)
	}
}

func TestCalculateMedianEven(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	median := calcMedian(data)

	if median != 6 {
		t.Errorf("median should be 6, but got %v", median)
	}
}

func TestCalculateMedianDispersedRandom(t *testing.T) {
	data := []float32{10, 12, 23, 23, 16, 23, 21, 16}

	median := calcMedian(data)

	if median != 18.5 {
		t.Errorf("median should be 18.5, but got %v", median)
	}
}

func TestNewStats(t *testing.T) {
	data := []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	stats := NewStats(data, 10)

	if stats.Min != 1 {
		t.Errorf("min should be 1, but got %v", stats.Min)
	}

	if stats.Mean != 5.5 {
		t.Errorf("mean should be 5.5, but got %v", stats.Mean)
	}

	if stats.Max != 10 {
		t.Errorf("max should be 10, but got %v", stats.Max)
	}

	if stats.StdDev != 2.872281323269 {
		t.Errorf("stddev should be 2.872281323269, but got %v", stats.StdDev)
	}

	if stats.Variance != 8.25 {
		t.Errorf("variance should be 8.25, but got %v", stats.Variance)
	}
}

func TestNewStatsMonochromeExposure(t *testing.T) {
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

	file, err := ioutil.ReadFile("../../data/m42-800x600-monochrome.json")

	if err != nil {
		t.Errorf("Error opening from JSON data: %s", err)
	}

	data := CameraExposure{}

	_ = json.Unmarshal([]byte(file), &data)

	xs := 800

	ys := 600

	mono := iris.NewMonochrome16Exposure(data.Image, 65535, xs, ys)

	mono.PreprocessImageArray(xs, ys)

	stats := NewStats(mono.Data, xs)

	if stats.Min != 3453 {
		t.Errorf("min should be 3453, but got %v", stats.Min)
	}

	if stats.Mean != 27448.309 {
		t.Errorf("mean should be 27448.309, but got %v", stats.Mean)
	}

	if stats.Max != 65535 {
		t.Errorf("max should be 65535, but got %v", stats.Max)
	}

	if stats.StdDev != 10592.966 {
		t.Errorf("stddev should be 10592.966, but got %v", stats.StdDev)
	}
}
