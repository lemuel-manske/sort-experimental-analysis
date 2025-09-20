package mergesort

import (
	"flag"
	"math/rand"
	"testing"
	"time"

	"io.com/sort/utils"
)

var size = flag.Int("size", 0, "Number of elements to sort")

func BenchmarkMergeSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		arr := utils.GenerateRandomSlice(*size)

		_ = MergeSort(arr)
	}
}

