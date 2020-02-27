package parallel

import (
	"runtime"
	"sync"
)

var semChan chan struct{}

// MergeSort performs the merge sort algorithm taking advantage of multiple processors.
func MergeSort(src []int64) {
	// We subtract 1 goroutine which is the one we are already running in.
	extraGoroutines := runtime.NumCPU() - 1
	semChan = make(chan struct{}, extraGoroutines)
	defer close(semChan)
	mergesort(src)
}

func mergesort(src []int64) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2

	left := make([]int64, mid)
	right := make([]int64, len(src)-mid)
	copy(left, src[:mid])
	copy(right, src[mid:])

	wg := sync.WaitGroup{}

	select {
	case semChan <- struct{}{}:
		wg.Add(1)
		go func() {
			mergesort(left)
			<-semChan
			wg.Done()
		}()
	default:
		// All other goroutines are blocked, let's do the job ourselves
		mergesort(left)
	}

	mergesort(right)

	wg.Wait()

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
