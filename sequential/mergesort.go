package sequential

// MergeSort performs the merge sort algorithm sequentially.
func MergeSort(src []int64) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2

	left := make([]int64, mid)
	right := make([]int64, len(src)-mid)
	copy(left, src[:mid])
	copy(right, src[mid:])

	MergeSort(left)
	MergeSort(right)

	merge(src, left, right)
}

func merge(result, left, right []int64) {
	var l, r, i int

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			result[i] = left[l]
			l++
		} else if r < len(right) {
			result[i] = right[r]
			r++
		}
		i++
	}
}
