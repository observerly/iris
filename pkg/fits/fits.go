package iris

// FITS Header struct:
type FITSHeader struct {
	Bools    map[string]bool
	Ints     map[string]int32
	Floats   map[string]float32
	Strings  map[string]string
	Dates    map[string]string
	Comments []string
	History  []string
	End      bool
	Length   int32
}

// Create a new instance of FITS header:
func NewFITSHeader() FITSHeader {
	return FITSHeader{
		Bools:    make(map[string]bool),
		Ints:     make(map[string]int32),
		Floats:   make(map[string]float32),
		Strings:  make(map[string]string),
		Dates:    make(map[string]string),
		Comments: make([]string, 0),
		History:  make([]string, 0),
		End:      false,
	}
}

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
