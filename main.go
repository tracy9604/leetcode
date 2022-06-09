package main

import (
	"fmt"
	"leetcode/algorithms"
)

func main() {
	arr := []int{3, 5, 2, 6}
	algorithms.MergeSort(arr)
	fmt.Println(arr)
}
