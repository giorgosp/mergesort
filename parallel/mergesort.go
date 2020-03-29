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
	mergesort(src, semChan)
}

func mergesort(src []int64, semChan chan struct{}) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2
	left := src[:mid]
	right := src[mid:]

	wg := sync.WaitGroup{}

	select {
	case semChan <- struct{}{}:
		wg.Add(1)
		go func() {
			mergesort(left, semChan)
			<-semChan
			wg.Done()
		}()
	default:
		// Can't create a new goroutine, let's do the job ourselves.
		mergesort(left, semChan)
	}

	mergesort(right, semChan)

	wg.Wait()

	merge(left, right)
}

// Merge in-place. left and right are slices of the original src array.
func merge(left, right []int64) {
	for i := range left {
		// If right[0] < left[i], swap them.
		if right[0] < left[i] {
			temp := left[i]
			left[i] = right[0]
			right[0] = temp

			// Move the new right[0] to the appropriate place in the right array, so that is sorted.
			first := right[0]
			var place int
			for place = 1; place < len(right) && first > right[place]; place++ {
				right[place-1] = right[place]
			}
			right[place-1] = first
		}
	}
}
