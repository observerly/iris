package iris

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewDefaultFITSHeaderEnd(t *testing.T) {
	var header = NewFITSHeader()

	var got = header.End

	var want bool = false

	if got != want {
		t.Errorf("NewFITSHeader() Header.End: got %v, want %v", got, want)
	}
}

func TestNewDefaultFITSHeaderWriteBoolean(t *testing.T) {
	var header = NewFITSHeader()

	sb := strings.Builder{}

	header.Bools["SIMPLE"] = struct {
		Value   bool
		Comment string
	}{Value: true, Comment: FITS_STANDARD}

	header.Write(&sb)

	got := sb.String()

	want := 160

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 160 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteString(t *testing.T) {
	var header = NewFITSHeader()

	sb := strings.Builder{}

	header.Strings["SIMPLE"] = struct {
		Value   string
		Comment string
	}{Value: "T", Comment: FITS_STANDARD}

	header.Write(&sb)

	got := sb.String()

	want := 160

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 160 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteStringContinue(t *testing.T) {
	var header = NewFITSHeader()

	sb := strings.Builder{}

	header.Strings["PROGRAM"] = struct {
		Value   string
		Comment string
	}{Value: "observerly Online FITS Exposure Generator", Comment: FITS_STANDARD}

	header.Write(&sb)

	got := sb.String()

	want := 240

	fmt.Println(got)

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 160 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteInt(t *testing.T) {
	var header = NewFITSHeader()

	sb := strings.Builder{}

	header.Ints = map[string]int32{
		"TEST": 1,
	}

	header.Write(&sb)

	got := sb.String()

	want := 160

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 160 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteFloat(t *testing.T) {
	var header = NewFITSHeader()

	sb := strings.Builder{}

	header.Floats = map[string]float32{
		"TEST": 1.0,
	}

	header.Write(&sb)

	got := sb.String()

	want := 160

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 160 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteEnd(t *testing.T) {
	var header = NewFITSHeader()

	sb := strings.Builder{}

	header.Strings["PROGRAM"] = struct {
		Value   string
		Comment string
	}{Value: "observerly Ltd", Comment: FITS_STANDARD}

	header.Write(&sb)

	got := sb.String()

	want := 160

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 160 characters: got %v, want %v", len(got), want)
	}

	if !header.End {
		t.Errorf("NewFITSHeader() Header.Write() exopected header.End to be true: got %v, want %v", header.End, true)
	}

	if !strings.Contains(got, "END") {
		t.Errorf("NewFITSHeader() Header.Write() exopected header to contain END: got %v, want %v", got, "END")
	}
}
