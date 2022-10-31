package iris

import (
	"io"
)

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

/*
  Writes a FITS header according to the FITS standard
  @see https://fits.gsfc.nasa.gov/standard40/fits_standard40aa-le.pdf
*/
func (h *FITSHeader) Write(w io.Writer) {}
