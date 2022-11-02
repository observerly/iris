package fits

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewDefaultFITSHeaderEnd(t *testing.T) {
	var header = NewFITSHeader(16, 2, 600, 800)

	var got = header.End

	var want bool = false

	if got != want {
		t.Errorf("NewFITSHeader() Header.End: got %v, want %v", got, want)
	}
}

func TestNewDefaultFITSHeaderWriteBoolean(t *testing.T) {
	var header = NewFITSHeader(16, 2, 600, 800)

	sb := strings.Builder{}

	header.Bools["SIMPLE"] = struct {
		Value   bool
		Comment string
	}{Value: true, Comment: FITS_STANDARD}

	header.Write(&sb)

	got := sb.String()

	want := 800

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 800 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteString(t *testing.T) {
	var header = NewFITSHeader(16, 2, 600, 800)

	sb := strings.Builder{}

	header.Strings["SIMPLE"] = struct {
		Value   string
		Comment string
	}{Value: "T", Comment: FITS_STANDARD}

	header.Write(&sb)

	got := sb.String()

	want := 880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteStringContinue(t *testing.T) {
	var header = NewFITSHeader(16, 2, 600, 800)

	sb := strings.Builder{}

	header.Strings["PROGRAM"] = struct {
		Value   string
		Comment string
	}{Value: "observerly Online FITS Exposure Generator", Comment: FITS_STANDARD}

	header.Write(&sb)

	got := sb.String()

	want := 880

	fmt.Println(got)

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteDate(t *testing.T) {
	var header = NewFITSHeader(16, 2, 600, 800)

	sb := strings.Builder{}

	header.Dates["EXPTIME"] = struct {
		Value   string
		Comment string
	}{Value: "2022-11-01T11:30:48.294Z", Comment: "ISO 8601 UTC Datetime YYYY-MM-DDTHH:mm:ss.sssZ"}

	header.Write(&sb)

	got := sb.String()

	want := 960

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 960 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteInt(t *testing.T) {
	var header = NewFITSHeader(16, 2, 600, 800)

	sb := strings.Builder{}

	header.Ints["BITPIX"] = struct {
		Value   int32
		Comment string
	}{
		Value:   32,
		Comment: "Number of bits per data pixel",
	}

	header.Write(&sb)

	got := sb.String()

	want := 800

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 800 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteFloat(t *testing.T) {
	var header = NewFITSHeader(16, 2, 600, 800)

	sb := strings.Builder{}

	header.Floats["HFR"] = struct {
		Value   float32
		Comment string
	}{
		Value:   0.1632387,
		Comment: "Median Half-Flux Radius (HFR) of the detected stars",
	}

	header.Write(&sb)

	got := sb.String()

	want := 880

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 880 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteEnd(t *testing.T) {
	var header = NewFITSHeader(16, 2, 600, 800)

	sb := strings.Builder{}

	header.Strings["PROGRAM"] = struct {
		Value   string
		Comment string
	}{Value: "observerly Ltd", Comment: FITS_STANDARD}

	header.Write(&sb)

	got := sb.String()

	want := 800

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 800 characters: got %v, want %v", len(got), want)
	}

	if !header.End {
		t.Errorf("NewFITSHeader() Header.Write() exopected header.End to be true: got %v, want %v", header.End, true)
	}

	if !strings.Contains(got, "END") {
		t.Errorf("NewFITSHeader() Header.Write() exopected header to contain END: got %v, want %v", got, "END")
	}
}
