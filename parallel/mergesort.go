package parallel

import (
	"runtime"
	"sync"
)

// MergeSort performs the merge sort algorithm taking advantage of multiple processors.
func MergeSort(src []int64) {
	// We subtract 1 goroutine which is the one we are already running in.
	extraGoroutines := runtime.NumCPU() - 1
	semChan := make(chan struct{}, extraGoroutines)
	defer close(semChan)
	temp := make([]int64, len(src))
	mergesort(src, temp, semChan)
}

func mergesort(src, temp []int64, semChan chan struct{}) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2

	left, lTemp := src[:mid], temp[:mid]
	right, rTemp := src[mid:], temp[mid:]

	wg := sync.WaitGroup{}

	select {
	case semChan <- struct{}{}:
		wg.Add(1)
		go func() {
			mergesort(left, lTemp, semChan)
			<-semChan
			wg.Done()
		}()
	default:
		// Can't create a new goroutine, let's do the job ourselves.
		mergesort(left, lTemp, semChan)
	}

	mergesort(right, rTemp, semChan)

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
