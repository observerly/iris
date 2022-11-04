package fits

import (
	"bytes"
	"image/jpeg"
	"io"
	"os"
	"testing"
)

func TestNewDefaultFITSImageHeaderEnd(t *testing.T) {
	var img = NewFITSImage(2, 600, 800)

	var got = img.Header.End

	var want bool = false

	if got != want {
		t.Errorf("NewFITSImage() Header.End: got %v, want %v", got, want)
	}
}

func TestNewDefaultFITSImageBScale(t *testing.T) {
	var img = NewFITSImage(2, 600, 800)

	var got = img.Bscale

	var want float32 = 1

	if got != want {
		t.Errorf("NewFITSImage() Bscale: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageID(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.ID

	var want int = 0

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() ID: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageFilename(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Filename

	var want string = ""

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Filename: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageBitpix(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Bitpix

	var want int32 = -32

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bitpix: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageBzero(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Bzero

	var want float32 = 0

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bzero: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageBscale(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Bscale

	var want float32 = 1

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bscale: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImagePixels(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Pixels

	var want int32 = 64

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Pixels: got %v, want %v", got, want)
	}
}

func TestNewFITSImageFrom2DDataID(t *testing.T) {
	var ex = [][]uint32{
		{1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
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

	var img2DData = NewFITSImageFrom2DData(ex, 2, 16, 16)

	var got = img2DData.ID

	var want int = 0

	if got != want {
		t.Errorf("NewFITSImageFrom2DData() ID: got %v, want %v", got, want)
	}
}

func TestNewFITSImageFrom2DDataPixels(t *testing.T) {
	var ex = [][]uint32{
		{1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
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

	var img2DData = NewFITSImageFrom2DData(ex, 2, 16, 16)

	var got = img2DData.Data

	var want = img2DData.Pixels

	if len(got) != int(want) {
		t.Errorf("NewFITSImageFrom2DData() Data Length should be 256 pixels: got %v, want %v", got, want)
	}
}

func TestNewFITSImageFrom2DDataData(t *testing.T) {
	var ex = [][]uint32{
		{1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
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

	var img2DData = NewFITSImageFrom2DData(ex, 2, 16, 16)

	var got = img2DData.Data

	var want int = 256

	if len(got) != want {
		t.Errorf("NewFITSImageFrom2DData() Data Length should be 256 pixels: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageID(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var imgImage = NewFITSImageFromImage(imgNaxisn)

	var got = imgImage.ID

	var want int = 0

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() ID: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageFilename(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Filename

	var want string = ""

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Filename: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageBitpix(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Bitpix

	var want int32 = -32

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bitpix: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageBzero(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Bzero

	var want float32 = 0

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bzero: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageBscale(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Bscale

	var want float32 = 1

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bscale: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImagePixels(t *testing.T) {
	var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

	var got = imgNaxisn.Pixels

	var want int32 = 64

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Pixels: got %v, want %v", got, want)
	}
}

func TestNewFITSImageFrom2DDataWriteFloatData(t *testing.T) {
	f, err := os.Open("../../images/noise16.jpeg")

	if err != nil {
		t.Errorf("Error opening image: %s", err)
	}

	defer f.Close()

	img, err := jpeg.Decode(f)

	if err != nil {
		t.Errorf("Error decoding image: %s", err)
	}

	bounds := img.Bounds()

	ex := make([][]uint32, bounds.Dx())

	for x := 0; x < bounds.Dx(); x++ {
		col := make([]uint32, bounds.Dy())
		ex[x] = col
	}

	for j := 0; j < bounds.Dy(); j++ {
		for i := 0; i < bounds.Dx(); i++ {
			r, g, b, _ := img.At(i, j).RGBA()
			lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			ex[i][j] = uint32(lum)
		}
	}

	var fit = NewFITSImageFrom2DData(ex, 2, int32(bounds.Dx()), int32(bounds.Dy()))

	var w io.Writer = os.Stdout

	buf := new(bytes.Buffer)

	buf, err = writeFloat32ArrayToBuffer(buf, fit.Data)

	if err != nil {
		t.Errorf("Error writing float32 array: %s", err)
	}

	if buf == nil {
		t.Errorf("Error writing float32 array: %s", err)
	}

	_, err = w.Write(buf.Bytes())

	if err != nil {
		t.Errorf("Error writing float32 array to standard output: %s", err)
	}
}
