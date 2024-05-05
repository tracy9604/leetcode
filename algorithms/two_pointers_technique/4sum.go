package two_pointers_technique

import "sort"

func Find4Sum(nums []int, target int) [][]int {
	sort.Ints(nums)

	ans := make([][]int, 0)
	n := len(nums)
	dupMap := make(map[[4]int]bool)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			left := j + 1
			right := n - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					key := [4]int{nums[i], nums[j], nums[left], nums[right]}
					if _, found := dupMap[key]; !found {
						ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
						dupMap[key] = true
					}

					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return ans
}
