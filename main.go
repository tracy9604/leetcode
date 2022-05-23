package main

import (
	"fmt"
	"leetcode/algorithms"
)

func main() {
	arr := []int{3, 5, 2, 6, 7, 8, 10}
	arr = algorithms.QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
