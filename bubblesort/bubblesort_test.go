package bubblesort

import (
	"flag"
	"math/rand"
	"testing"
	"time"

	"io.com/sort/utils"
)

var size = flag.Int("size", 0, "Number of elements to sort")

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		arr := utils.GenerateRandomSlice(*size)

		bubbleSort(arr)
	}
}

