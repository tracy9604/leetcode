package code

import "math"

func Jump(nums []int) int {
	res := 0
	n := len(nums)
	i := 0
	cur := 0
	for cur < n-1 {
		res++
		pre := cur
		for ; i <= pre; i++ {
			cur = int(math.Max(float64(cur), float64(i+nums[i])))
		}
		if pre == cur {
			return -1
		}
	}
	return res
}
