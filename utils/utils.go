package utils 

import (
	"math/rand"
)

func GenerateRandomSlice(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size * 10)
	}
	return arr
}

