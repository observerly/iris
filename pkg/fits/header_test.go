package fits

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestNewDefaultFITSHeaderEnd(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	var got = header.End

	var want bool = false

	if got != want {
		t.Errorf("NewFITSHeader() Header.End: got %v, want %v", got, want)
	}
}

func TestNewDefaultFITSHeaderWriteBoolean(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	buf := new(bytes.Buffer)

	header.Bools["SIMPLE"] = struct {
		Value   bool
		Comment string
	}{Value: true, Comment: FITS_STANDARD}

	header.WriteToBuffer(buf)

	got := buf.String()

	want := 2880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() expected length of 2880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteString(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	buf := new(bytes.Buffer)

	header.Strings["SIMPLE"] = struct {
		Value   string
		Comment string
	}{Value: "T", Comment: FITS_STANDARD}

	header.WriteToBuffer(buf)

	got := buf.String()

	want := 2880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() expected length of 2880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteStringContinue(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	buf := new(bytes.Buffer)

	header.Strings["PROGRAM"] = struct {
		Value   string
		Comment string
	}{Value: "observerly Online FITS Exposure Generator", Comment: FITS_STANDARD}

	header.WriteToBuffer(buf)

	got := buf.String()

	want := 2880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() expected length of 2880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteDate(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	buf := new(bytes.Buffer)

	header.Dates["EXPTIME"] = struct {
		Value   string
		Comment string
	}{Value: "2022-11-01T11:30:48.294Z", Comment: "ISO 8601 UTC Datetime YYYY-MM-DDTHH:mm:ss.sssZ"}

	header.WriteToBuffer(buf)

	got := buf.String()

	want := 2880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() expected length of 2880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteInt(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	buf := new(bytes.Buffer)

	header.Ints["BITPIX"] = struct {
		Value   int32
		Comment string
	}{
		Value:   32,
		Comment: "Number of bits per data pixel",
	}

	header.WriteToBuffer(buf)

	got := buf.String()

	want := 2880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() expected length of 2880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteFloat(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	buf := new(bytes.Buffer)

	header.Floats["HFR"] = struct {
		Value   float32
		Comment string
	}{
		Value:   0.1632387,
		Comment: "Median Half-Flux Radius (HFR) of the detected stars",
	}

	header.WriteToBuffer(buf)

	got := buf.String()

	want := 2880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() expected length of 2880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteEnd(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	buf := new(bytes.Buffer)

	header.Strings["PROGRAM"] = struct {
		Value   string
		Comment string
	}{Value: "observerly Ltd", Comment: FITS_STANDARD}

	header.WriteToBuffer(buf)

	got := buf.String()

	want := 2880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() expected length of 2880 characters: got %v, want %v", len(got), want)
	}

	if !header.End {
		t.Errorf("NewFITSHeader() Header.Write() expected header.End to be true: got %v, want %v", header.End, true)
	}

	if !strings.Contains(got, "END") {
		t.Errorf("NewFITSHeader() Header.Write() expected header to contain END: got %v, want %v", got, "END")
	}
}

func GetRegexSubValuesAndSubNames(str []byte) ([][]byte, []string) {
	return re.FindSubmatch(str), re.SubexpNames()
}

func TestCompileFITSHeaderRegExpSIMPLEKey(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("SIMPLE  =                    T / FITS Standard 4.0"))

	want := "SIMPLE"

	got := ""

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('k') {
			got = string(values[i])
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected key to be SIMPLE: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpSIMPLEValue(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("SIMPLE  =                    T / FITS Standard 4.0"))

	want := true

	got := false

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('b') && len(values[i]) > 0 {
			got = values[i][0] == byte('T')
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected value to be true: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpSIMPLEComment(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("SIMPLE  =                    T / FITS Standard 4.0"))

	want := "FITS Standard 4.0"

	got := ""

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('c') && len(values[i]) > 0 {
			got = strings.TrimSpace(string(values[i]))
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected value to be FITS Standard 4.0: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpNAXIS1Key(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("NAXIS1  =                 6000 / Number of pixels in axis 1"))

	want := "NAXIS1"

	got := ""

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('k') {
			got = string(values[i])
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected key to be NAXIS1: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpNAXIS1Value(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("NAXIS1  =                 6000 / Number of pixels in axis 1"))

	want := 6000

	got := 0

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('i') && len(values[i]) > 0 {
			value, err := strconv.ParseInt(string(values[i]), 10, 64)

			if err != nil {
				t.Errorf("CompileFITSHeaderRegExp() expected value to be a number: got %v, want %v", err, nil)
			}

			got = int(value)
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected value to be 6000: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpNAXIS1Comment(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("NAXIS1  =                 6000 / [1] Length of data axis 1"))

	want := "[1] Length of data axis 1"

	got := ""

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('c') && len(values[i]) > 0 {
			got = strings.TrimSpace(string(values[i]))
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected value to be [1] Length of data axis 1: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpSENSORKey(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("SENSOR  = 'Monochrome'         / ASCOM Alpaca Sensor Type"))

	want := "SENSOR"

	got := ""

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('k') {
			got = string(values[i])
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected key to be SENSOR: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpSENSORValue(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("SENSOR  = 'Monochrome'         / ASCOM Alpaca Sensor Type"))

	want := "Monochrome"

	got := ""

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('s') && len(values[i]) > 0 {
			got = strings.TrimSpace(string(values[i]))
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected value to be Monochrome: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpSENSORComment(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("SENSOR  = 'Monochrome'         / ASCOM Alpaca Sensor Type"))

	want := "ASCOM Alpaca Sensor Type"

	got := ""

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('c') && len(values[i]) > 0 {
			got = strings.TrimSpace(string(values[i]))
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected value to be ASCOM Alpaca Sensor Type: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderRegExpEND(t *testing.T) {
	values, names := GetRegexSubValuesAndSubNames([]byte("END"))

	want := "END"

	got := ""

	for i := 1; i < len(names); i++ {

		if values[i] == nil && len(names[i]) != 1 {
			t.Errorf("CompileFITSHeaderRegExp() expected value to be not nil: got %v, want %v", values[i], nil)
		}

		if names[i][0] == byte('E') {
			got = string(values[i])
		}
	}

	if got != want {
		t.Errorf("CompileFITSHeaderRegExp() expected key to be END: got %v, want %v", got, want)
	}
}

func TestCompileFITSHeaderParseLineBool(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	values, names := GetRegexSubValuesAndSubNames([]byte("SIMPLE  =                    T / Standard FITS format"))

	err := header.ParseLine(names, values)

	if err != nil {
		t.Errorf("CompileFITSHeaderParseDate() expected err to be nil: got %v, want %v", err, nil)
	}

	simple := header.Bools["SIMPLE"]

	if !simple.Value {
		t.Errorf("CompileFITSHeaderParseLine() expected SIMPLE to be true: got %v, want %v", simple.Value, true)
	}

	if simple.Comment != "Standard FITS format" {
		t.Errorf("CompileFITSHeaderParseLine() expected SIMPLE comment to be Standard FITS format: but got %v", simple.Comment)
	}
}

func TestCompileFITSHeaderParseLineInt32(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	values, names := GetRegexSubValuesAndSubNames([]byte("NAXIS1  =                  600 / [1] Length of data axis 1"))

	err := header.ParseLine(names, values)

	if err != nil {
		t.Errorf("CompileFITSHeaderParseDate() expected err to be nil: got %v, want %v", err, nil)
	}

	naxis1 := header.Ints["NAXIS1"]

	if naxis1.Value != 600 {
		t.Errorf("CompileFITSHeaderParseLine() expected NAXIS1 to be 600: but got %v", naxis1.Value)
	}

	if naxis1.Comment != "[1] Length of data axis 1" {
		t.Errorf("CompileFITSHeaderParseLine() expected NAXIS1 comment to be [1] Length of data axis 1: but got %v", naxis1.Comment)
	}
}

func TestCompileFITSHeaderParseLineFloat32(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	values, names := GetRegexSubValuesAndSubNames([]byte("EXPOSURE=                0.001 / [s] Exposure time"))

	err := header.ParseLine(names, values)

	if err != nil {
		t.Errorf("CompileFITSHeaderParseDate() expected err to be nil: got %v, want %v", err, nil)
	}

	exposure := header.Floats["EXPOSURE"]

	if exposure.Value != 0.001 {
		t.Errorf("CompileFITSHeaderParseLine() expected EXPOSURE to be 0.001: but got %v", exposure.Value)
	}

	if exposure.Comment != "[s] Exposure time" {
		t.Errorf("CompileFITSHeaderParseLine() expected EXPOSURE comment to be [s] Exposure time: but got %v", exposure.Comment)
	}
}

func TestCompileFITSHeaderParseLineString(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	values, names := GetRegexSubValuesAndSubNames([]byte("SENSOR  = 'Monochrome'         / ASCOM Alpaca Sensor Type"))

	err := header.ParseLine(names, values)

	if err != nil {
		t.Errorf("CompileFITSHeaderParseDate() expected err to be nil: got %v, want %v", err, nil)
	}

	sensor := header.Strings["SENSOR"]

	if sensor.Value != "Monochrome" {
		t.Errorf("CompileFITSHeaderParseLine() expected SENSOR to be true: got %v, want %v", sensor.Value, true)
	}

	if sensor.Comment != "ASCOM Alpaca Sensor Type" {
		t.Errorf("CompileFITSHeaderParseLine() expected SENSOR comment to be ASCOM Alpaca Sensor Type: but got %v", sensor.Comment)
	}
}

func TestCompileFITSHeaderParseDate(t *testing.T) {
	var header = NewFITSHeader(2, 600, 800)

	values, names := GetRegexSubValuesAndSubNames([]byte("DATE-OBS= '2020-01-01T00:00:00.000' / Observation Start Time UTC"))

	err := header.ParseLine(names, values)

	if err != nil {
		t.Errorf("CompileFITSHeaderParseDate() expected err to be nil: got %v, want %v", err, nil)
	}

	date := header.Strings["DATE-OBS"]

	if date.Value != "2020-01-01T00:00:00.000" {
		t.Errorf("CompileFITSHeaderParseLine() expected DATE-OBS to be true: but got %v", date.Value)
	}

	if date.Comment != "Observation Start Time UTC" {
		t.Errorf("CompileFITSHeaderParseLine() expected DATE-OBS comment to be Observation Start Time UTC: but got %v", date.Comment)
	}
}

func TestReadHeaderFromFile(t *testing.T) {
	// Attempt to open the file from the given filename:
	file, err := os.Open("../../samples/noise16.fits")

	if err != nil {
		t.Errorf("ReadHeaderFromFile() expected err to be nil: got %v, want %v", err, nil)
	}

	// Defer closing the file:
	defer file.Close()

	// Create a new FITS header:
	h := NewFITSHeader(2, 1, 1)

	// Read Header:
	err = h.Read(file)

	if err != nil {
		t.Errorf("ReadHeaderFromFile() expected err to be nil: got %v, want %v", err, nil)
	}

	// Check that the mandatory SIMPLE header value exists as per FITS standard:
	if !h.Bools["SIMPLE"].Value {
		t.Errorf("ReadHeaderFromFile() expected SIMPLE to be true: got %v, want %v", h.Bools["SIMPLE"].Value, true)
	}

	// Check that the mandatory BITPIX header value exists as per FITS standard:
	if h.Ints["BITPIX"].Value != -32 {
		t.Errorf("ReadHeaderFromFile() expected BITPIX to be -32: got %v, want %v", h.Ints["BITPIX"].Value, -32)
	}

	// Check that the mandatory NAXIS header value exists as per FITS standard:
	if h.Ints["NAXIS"].Value != 2 {
		t.Errorf("ReadHeaderFromFile() expected NAXIS to be 2: got %v, want %v", h.Ints["NAXIS"].Value, 2)
	}

	// Check that the mandatory NAXIS1 header value exists as per FITS standard:
	if h.Ints["NAXIS1"].Value != 1463 {
		t.Errorf("ReadHeaderFromFile() expected NAXIS1 to be 600: got %v, want %v", h.Ints["NAXIS1"].Value, 600)
	}

	// Check that the mandatory NAXIS2 header value exists as per FITS standard:
	if h.Ints["NAXIS2"].Value != 1168 {
		t.Errorf("ReadHeaderFromFile() expected NAXIS2 to be 800: got %v, want %v", h.Ints["NAXIS2"].Value, 800)
	}

	// Check that we have parsed to the END of the header:
	if !h.End {
		t.Errorf("ReadHeaderFromFile() expected END to be true: but got %v", h.End)
	}

	// Check that what we have parsed is divisble by 2880 bytes:
	if h.Length%2880 != 0 {
		t.Errorf("ReadHeaderFromFile() expected Length to be divisible by 2880: but got %v", h.Length)
	}
}
