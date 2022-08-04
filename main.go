package main

import (
	"fmt"
	"leetcode/sorting"
)

func main() {
	aliceSizes := []int{1, 2}
	bobSizes := []int{2, 3}
	fmt.Println(sorting.FairCandySwap(aliceSizes, bobSizes))
}
