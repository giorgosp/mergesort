package mergesort

import (
	"flag"
	"fmt"
	"testing"

	"github.com/giorgosp/mergesort/parallel"
	"github.com/giorgosp/mergesort/sequential"
	"github.com/giorgosp/mergesort/sorttest"
)

// We want the same benchmark name for both impls so that benchstat can compare them
// with each other. That's why we need a flag to specify the implementation instead
// of having BenchmarkSequential and BenchmarkParallel names.
var benchImpl = flag.String("impl", "parallel", "mergesort implementation to use in the benchmark.")

var impls = map[string]func([]int64){
	"sequential": sequential.MergeSort,
	"parallel":   parallel.MergeSort,
}

// BenchmarkMergeSort benchmarks mergesort. Uses the implementation specified
// by the benchImpl flag.
func BenchmarkMergeSort(b *testing.B) {
	fmt.Printf("Benchmarking %s mergesort\n", *benchImpl)

	elements := 16_000_000
	shuffled := sorttest.MakeArray(elements)
	src := make([]int64, elements)

	sort := impls[*benchImpl]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(src, shuffled)
		b.StartTimer()
		sort(src)
	}
}
