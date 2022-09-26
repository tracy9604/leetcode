package main

import (
	"fmt"
	"leetcode/code"
)

func main() {
	arr := [][]int{{1, 0},
		{0, 0},
		{1, 0}}
	k := 2
	fmt.Println(code.FindWeakestRows(arr, k))
}
