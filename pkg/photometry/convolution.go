/*****************************************************************************************************************/

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/iris/photometry
//	@license	Copyright © 2021-2025 observerly

/*****************************************************************************************************************/

package photometry

/*****************************************************************************************************************/

import "math"

/*****************************************************************************************************************/

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

/*****************************************************************************************************************/

func BiLinearConvolveGreenChannel(raw []uint32, w, h, xOffset, yOffset, x, y uint32) []float32 {
	G := make([]float32, int(x)*int(y))

	sqrt2 := float32(math.Sqrt2)

	for j := uint32(0); j < h; j += 2 {
		for i := uint32(0); i < w; i += 2 {
			// Source Offset:
			so := (j+yOffset)*w + (i + xOffset)

			// Destination Offset:
			do := (j)*x + (i)

			g1 := float32(raw[so+1])

			g2 := float32(raw[so+w])

			g3 := float32(2.0*g1+sqrt2*g2) * (1.0 / (2.0 + sqrt2))

			g4 := float32(sqrt2*g1+2.0*g2) * (1.0 / (2.0 + sqrt2))

			if i+xOffset > 0 {
				g3 = float32(raw[so-1])
			}

			if j+yOffset > 0 {
				g4 = float32(raw[so-w])
			}

			g5 := (2.0*g1 + sqrt2*g2) * (1.0 / (2.0 + sqrt2))

			g6 := (sqrt2*g1 + 2.0*g2) * (1.0 / (2.0 + sqrt2))

			if i+xOffset < w-2 {
				g5 = float32(raw[so+2+w])
			}

			if j+yOffset < y-2 {
				g6 = float32(raw[so+1+2*w])
			}

			G[do] = 0.25 * (g1 + g2 + g3 + g4)
			G[do+1] = g1
			G[do+x] = g2
			G[do+1+x] = 0.25 * (g1 + g2 + g5 + g6)
		}
	}

	return G
}

/*****************************************************************************************************************/

func BiLinearConvolveBlueChannel(raw []uint32, w, h, xOffset, yOffset, x, y uint32) []float32 {
	B := make([]float32, int(x)*int(y))

	for j := uint32(0); j < h; j += 2 {
		for i := uint32(0); i < w; i += 2 {
			// Source Offset:
			so := (j+yOffset)*w + (i + xOffset)

			// Destination Offset:
			do := (j)*x + (i)

			b1 := float32(raw[so+1+x])

			b2, b3, b4 := b1, b1, b1

			if i+xOffset > 0 {
				b2 = float32(raw[so-1+x])
			}

			if j+yOffset > 0 {
				b3 = float32(raw[so+1-x])
			}

			if i+xOffset > 0 && j+yOffset > 0 {
				b3 = float32(raw[so+1-x])
				b4 = float32(raw[so-1-x])
			}

			B[do] = 0.25 * (b1 + b2 + b3 + b4)
			B[do+1] = 0.5 * (b1 + b3)
			B[do+x] = 0.5 * (b1 + b2)
			B[do+1+x] = b1
		}
	}

	return B
}

/*****************************************************************************************************************/
