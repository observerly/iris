package utils

import "time"

/*
  RNG is a pseudorandom number generator.

  It is unsafe to call RNG methods from concurrent goroutines.
*/
type RNG struct {
	x uint32
}

func getRandomUint32() uint32 {
	x := time.Now().UnixNano()
	return uint32((x >> 32) ^ x)
}

/*
  Uint32 returns pseudorandom uint32 using an XOrShift algorithm.

  It is unsafe to call this method from concurrent goroutines.
*/
func (r *RNG) Uint32() uint32 {
	for r.x == 0 {
		r.x = getRandomUint32()
	}

	x := r.x
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	r.x = x
	return x
}

/*
  Uint32n returns pseudorandom uint32 in the range [0, ...maxN].

  It is unsafe to call this method from concurrent goroutines.

  See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
*/
func (r *RNG) Uint32n(maxN uint32) uint32 {
	x := r.Uint32()
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}
