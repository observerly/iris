package fits

import (
	"bytes"
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
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 2880 characters: got %v, want %v", len(got), want)
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
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 2880 characters: got %v, want %v", len(got), want)
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
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 2880 characters: got %v, want %v", len(got), want)
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
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 2880 characters: got %v, want %v", len(got), want)
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
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 2880 characters: got %v, want %v", len(got), want)
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
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 2880 characters: got %v, want %v", len(got), want)
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
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 2880 characters: got %v, want %v", len(got), want)
	}

	if !header.End {
		t.Errorf("NewFITSHeader() Header.Write() exopected header.End to be true: got %v, want %v", header.End, true)
	}

	if !strings.Contains(got, "END") {
		t.Errorf("NewFITSHeader() Header.Write() exopected header to contain END: got %v, want %v", got, "END")
	}
}
