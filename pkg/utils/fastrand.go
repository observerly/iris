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
