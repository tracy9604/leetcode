package main

import (
	"fmt"
	"leetcode/sorting"
)

func main() {
	nums := []int{2, -3, -1, 5, -4}
	k := 2
	fmt.Println(sorting.LargestSumAfterKNegations(nums, k))
}
