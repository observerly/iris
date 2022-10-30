package iris

// FITS Image struct:
// @see https://fits.gsfc.nasa.gov/fits_primer.html
// @see https://fits.gsfc.nasa.gov/standard40/fits_standard40aa-le.pdf
type FITSImage struct {
	ID       int        // Sequential ID number, for log output. Counted upwards from 0 for light frames. By convention, dark is -1 and flat is -2
	Filename string     // Original file name, if any, for log output.
	Header   FITSHeader // The FITS Header with all keys, values, comments, history entries etc.
	Bitpix   int32      // Bits per pixel value from the header. Positive values are integral, negative floating.
	Bzero    float32    // Zero offset. (True pixel value is Bzero + Bscale * Data[i]).
	Bscale   float32    // Value scaler. (True pixel value is Bzero + Bscale * Data[i]).
	Naxisn   []int32    // Axis dimensions. Most quickly varying dimension first (i.e. X,Y)
	Pixels   int32      // Number of pixels in the image. Product of Naxisn[]
	Data     []float32  // The image data
	Exposure float32    // Image exposure in seconds
}

// Creates a new instance of FITS image initialized with empty header
func NewFITSImage() *FITSImage {
	return &FITSImage{
		Header: NewFITSHeader(),
		Bscale: 1,
	}
}

// Creates a new instance of FITS image from given naxisn:
// (Data is not copied, allocated if nil. naxisn is deep copied)
func NewFITSImageFromNaxisn(naxisn []int32, data []float32) *FITSImage {
	numPixels := int32(1)

	for _, naxis := range naxisn {
		numPixels *= naxis
	}

	if data == nil {
		data = make([]float32, numPixels)
	}

	return &FITSImage{
		ID:       0,
		Filename: "",
		Header:   NewFITSHeader(),
		Bitpix:   -32,
		Bzero:    0,
		Bscale:   1,
		Naxisn:   append([]int32(nil), naxisn...), // clone slice
		Pixels:   numPixels,
		Data:     data,
		Exposure: 0,
	}
}

// Creates a new instance of FITS image from given image:
// (New data array will be allocated)
func NewFITSImageFromImage(img *FITSImage) *FITSImage {
	data := make([]float32, img.Pixels)

	return &FITSImage{
		ID:       img.ID,
		Filename: img.Filename,
		Header:   img.Header,
		Bitpix:   img.Bitpix,
		Bzero:    img.Bzero,
		Bscale:   img.Bscale,
		Naxisn:   append([]int32(nil), img.Naxisn...), // clone slice
		Pixels:   img.Pixels,
		Data:     data,
		Exposure: img.Exposure,
	}
}
