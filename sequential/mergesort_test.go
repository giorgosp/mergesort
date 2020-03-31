package sequential

import (
	"testing"

	"github.com/giorgosp/mergesort/sorttest"
)

func TestMergesort(t *testing.T) {
	sorttest.Test(MergeSort, 1000, t)
}

func BenchmarkMergeSort(b *testing.B) {
	sorttest.Benchmark(MergeSort, 16_000_000, b)
}
