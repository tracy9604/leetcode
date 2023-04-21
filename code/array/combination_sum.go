package array

import "sort"

func CombinationSum(candidates []int, target int) [][]int {
	temp := make([]int, 0)
	ans := make([][]int, 0)
	sort.Ints(candidates)
	findSum(candidates, temp, &ans, target, 0)
	return ans
}

func findSum(candidates, temp []int, ans *[][]int, target, index int) {
	if target == 0 {
		ansItem := make([]int, 0)
		ansItem = append(ansItem, temp...)
		*ans = append(*ans, ansItem)
		return
	}

	for index < len(candidates) && target-candidates[index] >= 0 {
		temp = append(temp, candidates[index])
		findSum(candidates, temp, ans, target-candidates[index], index)
		index++
		temp = temp[:len(temp)-1]
	}
}
