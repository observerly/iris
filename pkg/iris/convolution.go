package iris

import "math"

func BiLinearConvolveRedChannel(i, j uint32, raw []uint32, red []float32, w, h, xOffset, yOffset, x, y uint32) []float32 {
	// Source Offset:
	so := (i+yOffset)*w + (j + xOffset)

	// Destination Offset:
	do := (i)*x + (j)

	r1 := float32(raw[so])

	r2, r3, r4 := r1, r1, r1

	if j+xOffset < w-2 && i+yOffset < y-2 {
		r3 = float32(raw[so+2*w])
		r4 = float32(raw[so+2+2*w])
	}

	if j+xOffset < w-2 {
		r2 = float32(raw[so+2])
	}

	if i+yOffset < y-2 {
		r3 = float32(raw[so+2*w])
	}

	red[do] = r1
	red[do+1] = 0.5 * (r1 + r2)
	red[do+x] = 0.5 * (r1 + r3)
	red[do+1+x] = 0.25 * (r1 + r2 + r3 + r4)

	return red
}

func BiLinearConvolveGreenChannel(i, j uint32, raw []uint32, green []float32, w, h, xOffset, yOffset, x, y uint32) []float32 {
	// Source Offset:
	so := (i+yOffset)*w + (j + xOffset)

	// Destination Offset:
	do := (i)*x + (j)

	sqrt2 := float32(math.Sqrt2)

	g1 := float32(raw[so+1])

	g2 := float32(raw[so+w])

	g3 := float32(2.0*g1+sqrt2*g2) * (1.0 / (2.0 + sqrt2))

	g4 := float32(sqrt2*g1+2.0*g2) * (1.0 / (2.0 + sqrt2))

	if j+xOffset > 0 {
		g3 = float32(raw[so-1])
	}

	if i+yOffset > 0 {
		g4 = float32(raw[so-w])
	}

	g5 := (2.0*g1 + sqrt2*g2) * (1.0 / (2.0 + sqrt2))

	g6 := (sqrt2*g1 + 2.0*g2) * (1.0 / (2.0 + sqrt2))

	if j+xOffset < w-2 {
		g5 = float32(raw[so+2+w])
	}

	if i+yOffset < y-2 {
		g6 = float32(raw[so+1+2*w])
	}

	green[do] = 0.25 * (g1 + g2 + g3 + g4)
	green[do+1] = g1
	green[do+x] = g2
	green[do+1+x] = 0.25 * (g1 + g2 + g5 + g6)

	return green
}

func BiLinearConvolveBlueChannel(i, j uint32, raw []uint32, blue []float32, w, h, xOffset, yOffset, x, y uint32) []float32 {
	// Source Offset:
	so := (i+yOffset)*x + (j + xOffset)

	// Destination Offset:
	do := (i)*x + (j)

	b1 := float32(raw[so+1+x])

	b2, b3, b4 := b1, b1, b1

	if j+xOffset > 0 {
		b2 = float32(raw[so-1+x])
	}

	if i+yOffset > 0 {
		b3 = float32(raw[so+1-x])
	}

	if j+xOffset > 0 && i+yOffset > 0 {
		b3 = float32(raw[so+1-x])
		b4 = float32(raw[so-1-x])
	}

	blue[do] = 0.25 * (b1 + b2 + b3 + b4)
	blue[do+1] = 0.5 * (b1 + b3)
	blue[do+x] = 0.5 * (b1 + b2)
	blue[do+1+x] = b1

	return blue
}
