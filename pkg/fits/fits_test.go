package iris

import "testing"

var img = NewFITSImage(16, 2, 600, 800)

func TestNewDefaultFITSImageHeaderEnd(t *testing.T) {
	var got = img.Header.End

	var want bool = false

	if got != want {
		t.Errorf("NewFITSImage() Header.End: got %v, want %v", got, want)
	}
}

func TestNewDefaultFITSImageBScale(t *testing.T) {
	var got = img.Bscale

	var want float32 = 1

	if got != want {
		t.Errorf("NewFITSImage() Bscale: got %v, want %v", got, want)
	}
}

var imgNaxisn = NewFITSImageFromNaxisn([]int32{8, 8}, nil, 16, 2, 600, 800)

func TestNewFromNaxisnFITSImageID(t *testing.T) {
	var got = imgNaxisn.ID

	var want int = 0

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() ID: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageFilename(t *testing.T) {
	var got = imgNaxisn.Filename

	var want string = ""

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Filename: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageBitpix(t *testing.T) {
	var got = imgNaxisn.Bitpix

	var want int32 = -32

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bitpix: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageBzero(t *testing.T) {
	var got = imgNaxisn.Bzero

	var want float32 = 0

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bzero: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImageBscale(t *testing.T) {
	var got = imgNaxisn.Bscale

	var want float32 = 1

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bscale: got %v, want %v", got, want)
	}
}

func TestNewFromNaxisnFITSImagePixels(t *testing.T) {
	var got = imgNaxisn.Pixels

	var want int32 = 64

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Pixels: got %v, want %v", got, want)
	}
}

var imgImage = NewFITSImageFromImage(imgNaxisn)

func TestNewFromImageFITSImageID(t *testing.T) {
	var got = imgImage.ID

	var want int = 0

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() ID: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageFilename(t *testing.T) {
	var got = imgNaxisn.Filename

	var want string = ""

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Filename: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageBitpix(t *testing.T) {
	var got = imgNaxisn.Bitpix

	var want int32 = -32

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bitpix: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageBzero(t *testing.T) {
	var got = imgNaxisn.Bzero

	var want float32 = 0

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bzero: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImageBscale(t *testing.T) {
	var got = imgNaxisn.Bscale

	var want float32 = 1

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Bscale: got %v, want %v", got, want)
	}
}

func TestNewFromImageFITSImagePixels(t *testing.T) {
	var got = imgNaxisn.Pixels

	var want int32 = 64

	if got != want {
		t.Errorf("NewFITSImageFromNaxisn() Pixels: got %v, want %v", got, want)
	}
}
