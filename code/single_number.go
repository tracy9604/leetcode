package code

import "fmt"

func SingleNumber(nums []int) int {
	r := 0
	for _, n := range nums {
		r ^= n
		fmt.Println(r)
	}
	return r
}
