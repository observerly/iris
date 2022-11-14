package utils

import (
	"image"
	"math"
	"runtime"
	"sync"
)

/*
	Creates a mask of given pixel radius.

	Returns a list of index offsets
*/
func CreateRadialPixelMask(width int32, radius float32) []int32 {
	mask := []int32{}

	rad := int32(radius)

	for y := -rad; y <= rad; y++ {
		for x := -rad; x <= rad; x++ {
			dist := float32(math.Sqrt(float64(y*y + x*x)))

			if dist <= radius+1e-8 {
				offset := y*int32(width) + x
				mask = append(mask, offset)
			}
		}
	}

	return mask
}

/*
  DeferForEachPixel

  Loop through all of the image pixels avaialble and calls a functions for each { x, y } pixel position.

  The image is sub-divided into N * N blocks, where N is the number of available processor threads.

  For each available block, a parallel coroutine is added to the wait group.
*/
func DeferForEachPixel(size image.Point, f func(x int, y int)) {
	procs := runtime.GOMAXPROCS(0)

	var waitGroup sync.WaitGroup

	for i := 0; i < procs; i++ {
		startX := i * int(math.Floor(float64(size.X)/float64(procs)))

		var endX int

		if i < procs-1 {
			endX = (i + 1) * int(math.Floor(float64(size.X)/float64(procs)))
		} else {
			endX = size.X
		}

		for j := 0; j < procs; j++ {
			startY := j * int(math.Floor(float64(size.Y)/float64(procs)))

			var endY int

			if j < procs-1 {
				endY = (j + 1) * int(math.Floor(float64(size.Y)/float64(procs)))
			} else {
				endY = size.Y
			}

			waitGroup.Add(1)

			go func(sX int, eX int, sY int, eY int) {
				defer waitGroup.Done()
				for x := sX; x < eX; x++ {
					for y := sY; y < eY; y++ {
						f(x, y)
					}
				}
			}(startX, endX, startY, endY)

			waitGroup.Wait()
		}
	}
}
