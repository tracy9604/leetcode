package daily

func SubSets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	ans := make([][]int, 0)
	tmpSubSets := make([]int, 0)
	var dfs func(i int)

	dfs = func(i int) {
		if i == len(nums) {
			subsetCopy := make([]int, len(tmpSubSets))
			copy(subsetCopy, tmpSubSets)
			ans = append(ans, subsetCopy)
			return
		}

		dfs(i + 1)

		tmpSubSets = append(tmpSubSets, nums[i])

		dfs(i + 1)

		tmpSubSets = tmpSubSets[:len(tmpSubSets)-1]

	}

	dfs(0)

	return ans
}
