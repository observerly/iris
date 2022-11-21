package fits

import (
	"bytes"
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
