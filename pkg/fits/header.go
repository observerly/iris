package fits

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Regular expression parser for FITS header lines:
var re *regexp.Regexp = compileFITSHeaderRegEx()

type FITSHeaderBool struct {
	Value   bool
	Comment string
}

type FITSHeaderInt struct {
	Value   int32
	Comment string
}

type FITSHeaderFloat struct {
	Value   float32
	Comment string
}

type FITSHeaderString struct {
	Value   string
	Comment string
}

// FITS Header struct:
type FITSHeader struct {
	Bitpix   int32
	Naxis    int32
	Naxis1   int32
	Naxis2   int32
	Bools    map[string]FITSHeaderBool
	Ints     map[string]FITSHeaderInt
	Floats   map[string]FITSHeaderFloat
	Strings  map[string]FITSHeaderString
	Dates    map[string]FITSHeaderString
	Comments []string
	History  []string
	End      bool
	Length   int32
}

// Create a new instance of FITS header:
func NewFITSHeader(naxis int32, naxis1 int32, naxis2 int32) FITSHeader {
	h := FITSHeader{
		Bools:    make(map[string]FITSHeaderBool),
		Ints:     make(map[string]FITSHeaderInt),
		Floats:   make(map[string]FITSHeaderFloat),
		Strings:  make(map[string]FITSHeaderString),
		Dates:    make(map[string]FITSHeaderString),
		Comments: make([]string, 0),
		History:  make([]string, 0),
		End:      false,
	}

	h.Bitpix = -32

	h.Naxis = naxis

	h.Naxis1 = naxis1

	h.Naxis2 = naxis2

	h.Strings["XTENSION"] = struct {
		Value   string
		Comment string
	}{
		Value:   "IMAGE   ",
		Comment: "FITS Image Extension",
	}

	h.Strings["TIMESYS"] = struct {
		Value   string
		Comment string
	}{
		Value:   "UTC",
		Comment: "The temporal reference frame",
	}

	h.Dates["DATE"] = struct {
		Value   string
		Comment string
	}{
		Value:   time.Now().Format("2006-01-02T15:04:05"),
		Comment: "Created Timestamp FITS file was generated",
	}

	h.Strings["ORIGIN"] = struct {
		Value   string
		Comment string
	}{
		Value:   "observerly",
		Comment: "The organization or institution responsible for creating the FITS file",
	}

	h.Strings["PROGRAM"] = struct {
		Value   string
		Comment string
	}{Value: "@observerly/iris", Comment: "@observerly/iris FITS Exposure Generator"}

	return h
}

func (h *FITSHeader) Read(r io.Reader) error {
	block := make([]byte, 2880)

	for h.Length = 0; !h.End; {
		// Read the next 2880 byte block:
		bytesRead, err := io.ReadFull(r, block)

		if err != nil || bytesRead != 2880 {
			return err
		}

		// Increment the header length by the bytes block size:
		h.Length += int32(bytesRead)

		// Parse the header block by block:
		for n := 0; n < 2880/80 && !h.End; n++ {
			line := block[n*80 : (n+1)*80]

			values := re.FindSubmatch(line)

			if len(values) == 0 || values == nil {
				continue
			}

			names := re.SubexpNames()

			h.ParseLine(names, values)
		}
	}

	return nil
}

/*
Writes a FITS header according to the FITS standard to output bytes buffer
@see https://fits.gsfc.nasa.gov/standard40/fits_standard40aa-le.pdf
*/
func (h *FITSHeader) WriteToBuffer(buf *bytes.Buffer) (*bytes.Buffer, error) {
	// SIMPLE needs to be the leading HDR value:
	writeBool(buf, "SIMPLE", true, FITS_STANDARD)
	// BITPIX needs to be the seconda leading HDR value:
	writeInt(buf, "BITPIX", -32, "Number of bits per data pixel")
	// NAXIS header:
	writeInt(buf, "NAXIS", h.Naxis, "[1] Number of array dimensions")
	// NAXIS1 header:
	writeInt(buf, "NAXIS1", h.Naxis1, "[1] Length of data axis 1")
	// NAXIS2 header:
	writeInt(buf, "NAXIS2", h.Naxis2, "[1] Length of data axis 2")
	// BSCALE Header:
	writeInt(buf, "BSCALE", 1, "")
	// BZERO Header:
	writeInt(buf, "BZERO", 0, "")

	// Write the rest of the header values:
	for k, v := range h.Bools {
		writeBool(buf, k, v.Value, v.Comment)
	}

	for k, v := range h.Strings {
		writeString(buf, k, v.Value, v.Comment)
	}

	for k, v := range h.Ints {
		writeInt(buf, k, v.Value, v.Comment)
	}

	for k, v := range h.Floats {
		writeFloat(buf, k, v.Value, v.Comment)
	}

	for k, v := range h.Dates {
		writeString(buf, k, v.Value, v.Comment)
	}

	h.End = writeEnd(buf)

	// Pad current header block with spaces if necessary:
	bytesInHeaderBlock := (buf.Len() % 2880)

	if bytesInHeaderBlock > 0 {
		for i := bytesInHeaderBlock; i < 2880; i++ {
			buf.WriteRune(' ')
		}
	}

	return buf, nil
}

// Reads a FITS header line by line and returns a FITSHeader struct
func (h *FITSHeader) ParseLine(subNames []string, subValues [][]byte) error {
	// The KEY will always be a string of maximum 8 characters:
	key := ""

	// The COMMENT will always be a string of maximum 47 characters:
	comment := ""

	value := interface{}(nil)

	// Ignore index 0 which is the whole line:
	for i := 1; i < len(subNames); i++ {
		if subValues[i] != nil && len(subNames[i]) == 1 {
			switch c := subNames[i][0]; c {

			// End of header line:
			case byte('E'):
				h.End = true

			// Comment line:
			case byte('C'):
				h.Comments = append(h.Comments, string(subValues[i]))

			// History line:
			case byte('H'):
				h.History = append(h.History, string(subValues[i]))

			// Keyword line:
			case byte('k'): // Keyword line
				key = strings.TrimSpace(string(subValues[i]))

			// Boolean value line:
			case byte('b'):
				if len(subValues[i]) > 0 {
					v := subValues[i][0]
					value = v == byte('t') || v == byte('T')
				}

			// Integer value line:
			case byte('i'):
				v, err := strconv.ParseInt(string(subValues[i]), 10, 64)
				if err != nil {
					return err
				}
				value = v

			// Float value line:
			case byte('f'):
				v, err := strconv.ParseFloat(string(subValues[i]), 64)
				if err != nil {
					return err
				}
				value = v

			// String value line:
			case byte('s'):
				value = strings.TrimSpace(string(subValues[i]))

			// Date-like string value line:
			case byte('d'): // date
				d, err := time.Parse(time.RFC3339, strings.TrimSpace(string(subValues[i])))

				if err == nil {
					value = d
				}

			// Comment
			case byte('c'):
				comment = strings.TrimSpace(strings.TrimSpace(string(subValues[i])))

			// The defauly case where we can't parse the line:
			default:
				return fmt.Errorf("FITSHeader.ParseLine: unknown line type: %s", string(subNames[i]))
			}
		}
	}

	// Check if value is a boolean:
	if v, ok := value.(bool); ok {
		h.Bools[key] = struct {
			Value   bool
			Comment string
		}{
			Value:   v,
			Comment: comment,
		}
	}

	// Check if value is an integer:
	if v, ok := value.(int64); ok {
		h.Ints[key] = struct {
			Value   int32
			Comment string
		}{
			Value:   int32(v),
			Comment: comment,
		}
	}

	// Check if value is a float:
	if v, ok := value.(float64); ok {
		h.Floats[key] = struct {
			Value   float32
			Comment string
		}{
			Value:   float32(v),
			Comment: comment,
		}
	}

	// Check if value is a string:
	if v, ok := value.(string); ok {
		h.Strings[key] = struct {
			Value   string
			Comment string
		}{
			Value:   v,
			Comment: comment,
		}
	}

	// Check if value is a date:
	if v, ok := value.(time.Time); ok {
		h.Dates[key] = struct {
			Value   string
			Comment string
		}{
			Value:   v.Format(time.RFC3339),
			Comment: comment,
		}
	}

	return nil
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

// Build regexp parser for FITS header lines
func compileFITSHeaderRegEx() *regexp.Regexp {
	white := "\\s+"
	whiteOpt := "\\s*"
	whiteLine := white

	hist := "HISTORY"
	rest := ".*"
	histLine := hist + white + "(?P<H>" + rest + ")"

	commKey := "COMMENT"
	commLine := commKey + white + "(?P<C>" + rest + ")"

	end := "(?P<E>END)"
	endLine := end + whiteOpt

	key := "(?P<k>[A-Z0-9_-]+)"
	equals := "="

	b := "(?P<b>[TF])"
	i := "(?P<i>[+-]?[0-9]+)"
	f := "(?P<f>[+-]?[0-9]*\\.[0-9]*(?:[ED][-+]?[0-9]+)?)"
	s := "'(?P<s>[^']*)'"
	// [TBI]: Ensure all ISO-8601 dates are parsed correctly:
	d := "(?P<d>[0-9]{1,4}-?[012][0-9]-?[0123][0-9]T[012][0-9]:?[0-5][0-9]:?[0-5][0-9].?[0-9]*)"

	val := "(?:" + b + "|" + i + "|" + f + "|" + s + "|" + d + ")"

	// [TBI]: CONTINUE for strings
	// [TBI]: Complex int: (nr, nr)
	// [TBI]: Complex float: (nr, nr)

	commOpt := "(?:/(?P<c>.*))?"
	keyLine := key + whiteOpt + equals + whiteOpt + val + whiteOpt + commOpt

	lineRe := "^(?:" + whiteLine + "|" + histLine + "|" + commLine + "|" + keyLine + "|" + endLine + ")$"

	return regexp.MustCompile(lineRe)
}
