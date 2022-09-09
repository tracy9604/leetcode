package main

import (
	"fmt"
	"leetcode/code"
)

func main() {
	//[2,21,43,38,0,42,33,7,24,13,12,27,12,24,5,23,29,48,30,31]
	//[2,42,38,0,43,21]
	arr1 := []int{2, 21, 43, 38, 0, 42, 33, 7, 24, 13, 12, 27, 12, 24, 5, 23, 29, 48, 30, 31}
	arr2 := []int{2, 42, 38, 0, 43, 21}
	fmt.Println(code.RelativeSortArray(arr1, arr2))
}
