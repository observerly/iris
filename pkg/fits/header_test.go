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

	header.Bools = map[string]bool{
		"TEST": true,
	}

	header.Write(&sb)

	got := sb.String()

	want := 80

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 80 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteString(t *testing.T) {
	var header = NewFITSHeader()

	sb := strings.Builder{}

	header.Strings = map[string]string{
		"TEST": "TEST",
	}

	header.Write(&sb)

	got := sb.String()

	want := 80

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 80 characters: got %v, want %v", len(got), want)
	}
}

func TestNewDefaultFITSHeaderWriteStringContinue(t *testing.T) {
	var header = NewFITSHeader()

	sb := strings.Builder{}

	header.Strings = map[string]string{
		"TEST": "LONGER TEST STRING THAT SHOULD BE TRUNCATED ON A NEW LINE",
	}

	header.Write(&sb)

	got := sb.String()

	want := 160

	fmt.Println(got)

	if len(got) != want {
		t.Errorf("NewFITSHeader() Header.Write() exopected length of 80 characters: got %v, want %v", len(got), want)
	}
}
