package code

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	temp := make([]int, 0)
	ans := make([][]int, 0)
	sort.Ints(candidates)
	findSumHelper(candidates, temp, &ans, target, 0)
	return ans
}

func findSumHelper(candidates, temp []int, ans *[][]int, target, index int) {
	if target == 0 {
		ansItem := make([]int, 0)
		ansItem = append(ansItem, temp...)
		*ans = append(*ans, ansItem)
		return
	}

	if target < 0 {
		return
	}

	for i := index; i < len(candidates); i++ {
		if i == index || candidates[i] != candidates[i-1] {
			temp = append(temp, candidates[i])
			findSumHelper(candidates, temp, ans, target-candidates[i], i+1)
			temp = temp[:len(temp)-1]
		}
	}
}
