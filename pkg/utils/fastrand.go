package utils

/*
  RNG is a pseudorandom number generator.

  It is unsafe to call RNG methods from concurrent goroutines.
*/
type RNG struct {
	x uint32
}
