package main

import (
	"fmt"
	"leetcode/code"
)

func main() {
	//[2,21,43,38,0,42,33,7,24,13,12,27,12,24,5,23,29,48,30,31]
	//[2,42,38,0,43,21]
	arr1 := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(code.ArrayRankTransform(arr1))
}
