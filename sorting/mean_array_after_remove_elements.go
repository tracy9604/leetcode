package sorting

import (
	"math"
	"sort"
)

func TrimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	sum := 0
	count := 0
	for i := n / 20; i < n*19/20; i++ {
		sum += arr[i]
		count++
	}
	result := float64(sum) / float64(count)
	return result * math.Pow(10, 5) / math.Pow(10, 5)
}
