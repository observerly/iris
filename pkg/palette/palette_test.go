package palette

import "testing"

func TestCombineRedChannelSimple(t *testing.T) {
	red := []PaletteChannel{
		{
			Data:     []float32{200, 200, 200},
			Fraction: 0.6,
		},
		{
			Data:     []float32{100, 100, 100},
			Fraction: 0.4,
		},
	}

	r, err := combinePaletteChannel(red)

	if err != nil {
		t.Errorf("error in constructing palette: %v", err)
	}

	if len(r) != 3 {
		t.Errorf("expected 3 values in red channel, got %v", len(r))
	}

	if r[0] != 80 {
		t.Errorf("expected 80 in red channel, got %v", r[0])
	}

	if r[1] != 80 {
		t.Errorf("expected 80 in red channel, got %v", r[1])
	}

	if r[2] != 80 {
		t.Errorf("expected 80 in red channel, got %v", r[2])
	}
}

func TestCombineGreenChannelSimple(t *testing.T) {
	green := []PaletteChannel{
		{
			Data:     []float32{220, 220, 220},
			Fraction: 0.6,
		},
		{
			Data:     []float32{200, 200, 200},
			Fraction: 0.4,
		},
	}

	g, err := combinePaletteChannel(green)

	if err != nil {
		t.Errorf("error in constructing palette: %v", err)
	}

	if len(g) != 3 {
		t.Errorf("expected 3 values in red channel, got %v", len(g))
	}

	if g[0] != 106 {
		t.Errorf("expected 106 in red channel, got %v", g[0])
	}

	if g[1] != 106 {
		t.Errorf("expected 106 in red channel, got %v", g[1])
	}

	if g[2] != 106 {
		t.Errorf("expected 106 in red channel, got %v", g[2])
	}
}

func TestCombineBlueChannelSimple(t *testing.T) {
	blue := []PaletteChannel{
		{
			Data:     []float32{100, 100, 100},
			Fraction: 0.6,
		},
		{
			Data:     []float32{200, 200, 200},
			Fraction: 0.4,
		},
	}

	b, err := combinePaletteChannel(blue)

	if err != nil {
		t.Errorf("error in constructing palette: %v", err)
	}

	if len(b) != 3 {
		t.Errorf("expected 3 values in red channel, got %v", len(b))
	}

	if b[0] != 70 {
		t.Errorf("expected 70 in red channel, got %v", b[0])
	}

	if b[1] != 70 {
		t.Errorf("expected 70 in red channel, got %v", b[1])
	}

	if b[2] != 70 {
		t.Errorf("expected 70 in red channel, got %v", b[2])
	}
}

func TestFromPaletteSimple(t *testing.T) {
	red := []PaletteChannel{
		{
			Data:     []float32{200, 200, 200},
			Fraction: 0.6,
		},
		{
			Data:     []float32{100, 100, 100},
			Fraction: 0.4,
		},
	}

	green := []PaletteChannel{
		{
			Data:     []float32{200, 200, 200},
			Fraction: 0.6,
		},
		{
			Data:     []float32{100, 100, 100},
			Fraction: 0.4,
		},
	}

	blue := []PaletteChannel{
		{
			Data:     []float32{200, 200, 200},
			Fraction: 0.6,
		},
		{
			Data:     []float32{100, 100, 100},
			Fraction: 0.4,
		},
	}

	r, _, _, err := FromPalette(&Palette{
		Name: "Test Palette",
		R:    red,
		G:    green,
		B:    blue,
	})

	if err != nil {
		t.Errorf("error in constructing palette: %v", err)
	}

	if len(r) != 3 {
		t.Errorf("expected 3 values in red channel, got %v", len(r))
	}

	if r[0] != 80 {
		t.Errorf("expected 80 in red channel, got %v", r[0])
	}

	if r[1] != 80 {
		t.Errorf("expected 80 in red channel, got %v", r[1])
	}

	if r[2] != 80 {
		t.Errorf("expected 80 in red channel, got %v", r[2])
	}
}
