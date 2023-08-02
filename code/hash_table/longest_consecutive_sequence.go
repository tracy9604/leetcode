package hash_table

import (
	"math"
	"sort"
)

func FindLongestConsecutiveSequence(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dupMap := make(map[int]bool)
	input := make([]int, 0)
	for _, item := range nums {
		if _, found := dupMap[item]; !found {
			input = append(input, item)
			dupMap[item] = true
		}
	}
	sort.Ints(input)
	count := 0
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		count++
		max = int(math.Max(float64(count), float64(max)))
		if nums[i]+1 != nums[i+1] {
			count = 1
		}
	}
	return max
}

func FindLongestConsecutiveSequenceV2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}

	count := 0
	max := 0

	for i := 0; i < len(nums); i++ {
		// nums[i] is the starting point
		count = 0
		if _, found := numsMap[nums[i]-1]; !found {
			j := nums[i]
			for {
				if _, check := numsMap[j]; check {
					count++
					j++
				} else {
					break
				}
			}
			max = int(math.Max(float64(max), float64(count)))
		}
	}
	return max
}
