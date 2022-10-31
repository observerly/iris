package iris

import (
	"fmt"
	"io"
	"strings"
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
func (h *FITSHeader) Write(w io.Writer) {
	for k, v := range h.Bools {
		writeBool(w, k, v, "")
	}

	for k, v := range h.Strings {
		writeString(w, k, v, "")
	}

	for k, v := range h.Ints {
		writeInt(w, k, v, "")
	}

	for k, v := range h.Floats {
		writeFloat(w, k, v, "")
	}

	for k, v := range h.Dates {
		writeString(w, k, v, "")
	}

	h.End = writeEnd(w)
}

// Writes a FITS header boolean T/F value
func writeBool(w io.Writer, key string, value bool, comment string) {
	if len(key) > 8 {
		key = key[0:8]
	}

	if len(comment) > 47 {
		comment = comment[0:47]
	}

	// Default false values are set to "F"
	v := "F"

	// If boolean value true, set to "T"
	if value {
		v = "T"
	}

	fmt.Fprintf(w, "%-8s= %20s / %-47s", key, v, comment)
}

// Writes a FITS header string value, with escaping and continuations if necessary.
func writeString(w io.Writer, key, value, comment string) {
	if len(key) > 8 {
		key = key[0:8]
	}
	if len(comment) > 47 {
		comment = comment[0:47]
	}

	// escape ' characters
	value = strings.Join(strings.Split(value, "'"), "''")

	if len(value) <= 18 {
		fmt.Fprintf(w, "%-8s= '%s'%s / %-47s", key, value, strings.Repeat(" ", 18-len(value)), comment)
	} else {
		fmt.Fprintf(w, "%-8s= '%s&' / %-47s", key, value[0:17], comment)

		value = value[17:]

		for len(value) > 66 {
			fmt.Fprintf(w, "CONTINUE  '%s&' ", value[0:66])
			value = value[66:]
		}

		fmt.Fprintf(w, "CONTINUE  '%s'%s", value, strings.Repeat(" ", 50+(18-len(value))))
	}
}

// Writes a FITS header integer value
func writeInt(w io.Writer, key string, value int32, comment string) {
	if len(key) > 8 {
		key = key[0:8]
	}

	if len(comment) > 47 {
		comment = comment[0:47]
	}

	fmt.Fprintf(w, "%-8s= %20d / %-47s", key, value, comment)
}

// Writes a FITS header float value
func writeFloat(w io.Writer, key string, value float32, comment string) {
	if len(key) > 8 {
		key = key[0:8]
	}

	if len(comment) > 47 {
		comment = comment[0:47]
	}

	fmt.Fprintf(w, "%-8s= %20g / %-47s", key, value, comment)
}

// Writes a FITS header end record
func writeEnd(w io.Writer) bool {
	n, _ := fmt.Fprintf(w, "END%s", strings.Repeat(" ", 80-3))
	return n > 0
}
