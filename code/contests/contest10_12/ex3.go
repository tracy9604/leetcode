package contest10_12

import (
	"fmt"
	"math"
)

func findMax(arr []int, n int) int {
	if n == 1 {
		return arr[0]
	}

	max := findMax(arr, n-1)

	if arr[n-1] > max {
		return arr[n-1]
	}
	return max
}

func CountSubarrays(nums []int, k int) int64 {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		max = int(math.Max(float64(max), float64(nums[i])))
	}

	i := 0
	j := i
	var rs int64
	arr := make([][]int, 0)
	count := 0

	for i < len(nums) {
		for j < len(nums) && nums[j] != max {
			j++
		}
		if nums[j] == max {
			count++
		}
		if count >= k {
			rs++
			arr = append(arr, nums[i:j])
		}

		if j == len(nums) {
			i++
			j = i
			count = 0
		}

	}

	fmt.Println(arr)
	return rs
}

func CountSubarrays2(nums []int, k int) int64 {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	ans, r, c := 0, 0, 0
	for l := 0; l < len(nums); l++ {
		for r < len(nums) && c < k {
			if nums[r] == max {
				c++
			}
			r++
		}
		if c == k {
			ans += len(nums) - r + 1
		}
		if nums[l] == max {
			c--
		}
	}
	return int64(ans)
}
