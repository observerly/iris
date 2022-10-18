package iris

import (
	"fmt"
	"os"
	"testing"
)

var ex = [][]uint32{}

func TestNewMonochromeExposureWidth(t *testing.T) {
	mono := NewMonochromeExposure(ex, 800, 600)

	var got int = mono.Width

	var want int = 800

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochromeExposureHeight(t *testing.T) {
	mono := NewMonochromeExposure(ex, 800, 600)

	var got int = mono.Height

	var want int = 600

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochromeExposurePixels(t *testing.T) {
	mono := NewMonochromeExposure(ex, 800, 600)

	var got int = mono.Pixels

	var want int = 480000

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewMonochromeExposurePreprocess4x4(t *testing.T) {
	var ex = [][]uint32{
		{6, 6, 6, 6},
		{6, 7, 8, 8},
		{7, 8, 9, 7},
		{6, 7, 8, 6},
	}

	mono := NewMonochromeExposure(ex, 4, 4)

	var x int = mono.Width

	var y int = mono.Height

	if x != 4 {
		t.Errorf("got %q, wanted %q", x, 4)
	}

	if y != 4 {
		t.Errorf("got %q, wanted %q", y, 4)
	}

	buff, err := mono.Preprocess()

	if err != nil {
		t.Errorf("Expected the buffer to be created successfully, but got %q", err)
	}

	f, err := os.Create("4x4image.jpg")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("4x4image.jpg")
	}()

	n, err := f.Write(buff.Bytes())

	if err != nil {
		t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
	}

	fmt.Println(n)

	if n >= 512 {
		t.Errorf("Expected the number of bytes to be approximately less than 128, but got %v", n)
	}
}

func TestNewMonochromeExposurePreprocess16x16(t *testing.T) {
	var ex = [][]uint32{
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{6, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 7, 6},
		{7, 8, 9, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 9, 8, 7},
		{6, 7, 8, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 8, 7, 6},
	}

	mono := NewMonochromeExposure(ex, 16, 16)

	var x int = mono.Width

	var y int = mono.Height

	if x != 16 {
		t.Errorf("got %q, wanted %q", x, 16)
	}

	if y != 16 {
		t.Errorf("got %q, wanted %q", y, 16)
	}

	buff, err := mono.Preprocess()

	if err != nil {
		t.Errorf("Expected the buffer to be created successfully, but got %q", err)
	}

	f, err := os.Create("16x16image.jpg")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
		}

		// Clean up the file after we have finished with the test:
		os.Remove("16x16image.jpg")
	}()

	n, err := f.Write(buff.Bytes())

	if err != nil {
		t.Errorf("Expected the image buffer to be saved successfully, but got %q", err)
	}

	if n >= 1086 {
		t.Errorf("Expected the number of bytes to be approximately less than 1086, but got %q", n)
	}
}
