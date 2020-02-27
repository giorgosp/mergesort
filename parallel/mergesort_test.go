package parallel

import (
	"testing"

	"github.com/giorgosp/mergesort/sorttest"
)

func TestMergesort(t *testing.T) {
	sorttest.Test(MergeSort, 1000, t)
}
