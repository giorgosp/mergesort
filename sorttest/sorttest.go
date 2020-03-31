package sorttest

import (
	"math/rand"
	"testing"
	"time"
)

// Test is a helper method to test sort implementations.
func Test(sort func([]int64), elements int, t *testing.T) {
	arr := makeArray(elements)

	sort(arr)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			// arr should be sorted in ascending order
			t.Fail()
		}
	}
}

// Benchmark is a helper method to benchmark a sort.
func Benchmark(sort func([]int64), elements int, b *testing.B) {
	shuffled := makeArray(elements)
	src := make([]int64, elements)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, shuffled)
		b.StartTimer()
		sort(src)
	}
}

// MakeArray creates an array of n random integers.
func makeArray(n int) []int64 {
	arr := make([]int64, n)
	rand.Seed(time.Now().Unix())
	for i := range arr {
		arr[i] = rand.Int63()
	}
	return arr
}
