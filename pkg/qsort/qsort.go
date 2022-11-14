package qsort

/*
  Partitions an array of float32 with the middle pivot element, and returns the pivot index.

  Values less than the pivot are moved left of the pivot, those greater are moved right.

  Array must not contain IEEE NaN
*/
func QPartitionFloat32(a []float32) int {
	left, right := 0, len(a)-1

	mid := (left + right) >> 1

	pivot := a[mid]

	l := left - 1
	r := right + 1

	for {
		for {
			l++
			if a[l] >= pivot {
				break
			}
		}
		for {
			r--
			if a[r] <= pivot {
				break
			}
		}
		if l >= r {
			return r
		}
		a[l], a[r] = a[r], a[l]
	}
}

/*
	Sort an array of float32 in ascending order.

	Array must not contain IEEE NaN
*/
func QSortFloat32(a []float32) {
	if len(a) > 1 {
		index := QPartitionFloat32(a)
		QSortFloat32(a[:index+1])
		QSortFloat32(a[index+1:])
	}
}

/*
	Select kth lowest element from an array of float32. Partially reorders the array.

	Array must not contain IEEE NaN
*/
func QSelectFloat32(a []float32, k int) float32 {
	left, right := 0, len(a)-1

	for left < right {
		mid := (left + right) >> 1

		pivot := a[mid]

		l, r := left-1, right+1

		for {
			for {
				l++
				if a[l] >= pivot {
					break
				}
			}
			for {
				r--
				if a[r] <= pivot {
					break
				}
			}
			if l >= r {
				break
			}
			a[l], a[r] = a[r], a[l]
		}

		index := r

		offset := index - left + 1

		if k <= offset {
			right = index
		} else {
			left = index + 1
			k = k - offset
		}
	}

	return a[left]
}
