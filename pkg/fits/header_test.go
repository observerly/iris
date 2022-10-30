package iris

import "testing"

var header = NewFITSHeader()

func TestNewDefaultFITSHeaderEnd(t *testing.T) {
	var got = header.End

	var want bool = false

	if got != want {
		t.Errorf("NewFITSHeader() Header.End: got %v, want %v", got, want)
	}
}
