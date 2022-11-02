package fits

import (
	"fmt"
	"io"
	"strings"
)

// FITS Header struct:
type FITSHeader struct {
	Bools map[string]struct {
		Value   bool
		Comment string
	}
	Ints map[string]struct {
		Value   int32
		Comment string
	}
	Floats map[string]struct {
		Value   float32
		Comment string
	}
	Strings map[string]struct {
		Value   string
		Comment string
	}
	Dates map[string]struct {
		Value   string
		Comment string
	}
	Comments []string
	History  []string
	End      bool
	Length   int32
}

// Create a new instance of FITS header:
func NewFITSHeader(bitpix int32, naxis int32, naxis1 int32, naxis2 int32) FITSHeader {
	h := FITSHeader{
		Bools: make(map[string]struct {
			Value   bool
			Comment string
		}),
		Ints: make(map[string]struct {
			Value   int32
			Comment string
		}),
		Floats: make(map[string]struct {
			Value   float32
			Comment string
		}),
		Strings: make(map[string]struct {
			Value   string
			Comment string
		}),
		Dates: make(map[string]struct {
			Value   string
			Comment string
		}),
		Comments: make([]string, 0),
		History:  make([]string, 0),
		End:      false,
	}

	h.Bools["SIMPLE"] = struct {
		Value   bool
		Comment string
	}{true, FITS_STANDARD}

	h.Ints["BITPIX"] = struct {
		Value   int32
		Comment string
	}{
		Value:   bitpix,
		Comment: "Number of bits per data pixel",
	}

	h.Ints["NAXIS"] = struct {
		Value   int32
		Comment string
	}{
		Value:   naxis,
		Comment: "Number of data axes",
	}

	h.Ints["NAXIS1"] = struct {
		Value   int32
		Comment string
	}{
		Value:   naxis1,
		Comment: "Length of data axis 1",
	}

	h.Ints["NAXIS2"] = struct {
		Value   int32
		Comment string
	}{
		Value:   naxis2,
		Comment: "Length of data axis 2",
	}

	h.Strings["XTENSION"] = struct {
		Value   string
		Comment string
	}{
		Value:   "IMAGE␣␣␣",
		Comment: "FITS Image Extension",
	}

	h.Strings["PROGRAM"] = struct {
		Value   string
		Comment string
	}{Value: "@observerly/iris", Comment: "@observerly/iris FITS Exposure Generator"}

	return h
}

/*
  Writes a FITS header according to the FITS standard
  @see https://fits.gsfc.nasa.gov/standard40/fits_standard40aa-le.pdf
*/
func (h *FITSHeader) Write(w io.Writer) {
	for k, v := range h.Bools {
		writeBool(w, k, v.Value, v.Comment)
	}

	for k, v := range h.Strings {
		writeString(w, k, v.Value, v.Comment)
	}

	for k, v := range h.Ints {
		writeInt(w, k, v.Value, v.Comment)
	}

	for k, v := range h.Floats {
		writeFloat(w, k, v.Value, v.Comment)
	}

	for k, v := range h.Dates {
		writeString(w, k, v.Value, v.Comment)
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
