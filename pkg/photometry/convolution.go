package photometry

func BiLinearConvolveRedChannel(raw []uint32, w, h, xOffset, yOffset, x, y uint32) []float32 {
	R := make([]float32, int(x)*int(y))

	for j := uint32(0); j < h; j += 2 {
		for i := uint32(0); i < w; i += 2 {
			// Source Offset:
			so := (j+yOffset)*w + (i + xOffset)

			// Destination Offset:
			do := (j)*x + (i)

			r := float32(raw[so])

			r2, r3, r4 := r, r, r

			if i+xOffset < w-2 {
				r2 = float32(raw[so+2])
			}

			if i+xOffset < w-2 && j+yOffset < h-2 {
				r3 = float32(raw[so+2*w])
				r4 = float32(raw[so+2+2*w])
			}

			if j+yOffset < h-2 {
				r3 = float32(raw[so+2*w])
			}

			R[do] = r
			R[do+1] = 0.5 * (r + r2)
			R[do+x] = 0.5 * (r + r3)
			R[do+1+x] = 0.25 * (r + r2 + r3 + r4)
		}
	}

	return R
}

