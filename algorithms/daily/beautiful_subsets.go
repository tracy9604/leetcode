package daily

func FindBeautifulSubSets(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := -1
	counter := make([]int, 1010)
	var dfs func(i int)

	dfs = func(i int) {
		if i >= len(nums) {
			ans++
			return
		}

		dfs(i + 1)
		isBeautifulWithNext := nums[i]+k >= len(counter) || counter[nums[i]+k] == 0
		isBeautifulWithPrev := nums[i]-k < 0 || counter[nums[i]-k] == 0
		if isBeautifulWithNext && isBeautifulWithPrev {
			counter[nums[i]]++
			dfs(i + 1)
			counter[nums[i]]--
		}
	}

	dfs(0)

	return ans
}
