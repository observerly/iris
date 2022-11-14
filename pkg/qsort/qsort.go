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

/*
	Select first quartile of an array of float32. Partially reorders the array.

	Array must not contain IEEE NaN
*/
func QSelectFirstQuartileFloat32(a []float32) float32 {
	return QSelectFloat32(a, (len(a)>>2)+1)
}

/*
	Select median of an array of float32. Partially reorders the array.

	Array must not contain IEEE NaN
*/
func QSelectMedianFloat32(a []float32) float32 {
	// Quickly  select the midpoint element:
	k := (len(a) >> 1) + 1

	// Get the upper kth element:
	upper := QSelectFloat32(a, k)

	// For odd lengths, the found element is the median:
	if (len(a) & 1) != 0 {
		return upper
	}

	// For even lengths, calculate the maximum of all elements below the pivot:
	lower := a[0]

	for i := 1; i < k-1; i++ {
		if a[i] > lower {
			lower = a[i]
		}
	}

	// Return average of the upper and lower elements:
	return 0.5 * (lower + upper)
}
