package iris

import (
	"strings"
	"testing"
)

var header = NewFITSHeader()

func TestNewDefaultFITSHeaderEnd(t *testing.T) {
	var got = header.End

	var want bool = false

	if got != want {
		t.Errorf("NewFITSHeader() Header.End: got %v, want %v", got, want)
	}
}

func TestNewDefaultFITSHeaderWriteBoolean(t *testing.T) {
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
