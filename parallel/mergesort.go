package parallel

import (
	"sort"
	"sync"
)

// MergeSort performs the merge sort algorithm taking advantage of multiple processors.
func MergeSort(src []int64) {
	temp := make([]int64, len(src))
	mergesort(src, temp)
}

func mergesort(src, temp []int64) {
	if len(src) <= 10000 {
		sort.Slice(src, func(i int, j int) bool { return src[i] <= src[j] })
		return
	}

	mid := len(src) / 2

	left, lTemp := src[:mid], temp[:mid]
	right, rTemp := src[mid:], temp[mid:]

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		mergesort(left, lTemp)
		wg.Done()
	}()

	mergesort(right, rTemp)

	wg.Wait()

	merge(src, temp, left, right)
}

func merge(src, result, left, right []int64) {
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
	copy(src, result)
}
