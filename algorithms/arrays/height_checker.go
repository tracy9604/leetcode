package arrays

import "sort"

func HeightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)
	count := 0
	for i := 0; i < len(heights); i++ {
		if expected[i] != heights[i] {
			count++
		}
	}
	return count
}

func HeightCheckerV2(heights []int) int {
	cnt := make([]int, 101)
	for i := 0; i < len(heights); i++ {
		cnt[heights[i]]++
	}

	ans, k := 0, 0
	for i := 1; i < 101; i++ {
		for cnt[i] > 0 {
			cnt[i]--
			if heights[k] != i {
				ans++
			}
			k++
		}
	}
	return ans
}
