package iris

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
