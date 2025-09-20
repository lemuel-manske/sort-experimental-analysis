package quicksort

import (
	"flag"
	"math/rand"
	"testing"
	"time"

	"io.com/sort/utils"
)

var size = flag.Int("size", 0, "Number of elements to sort")

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	pivotIndex := len(arr) / 2

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
}

func BenchmarkQuickSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		arr := utils.GenerateRandomSlice(*size)
		QuickSort(arr)
	}
}
