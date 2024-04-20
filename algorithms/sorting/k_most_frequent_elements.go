package sorting

import "sort"

type Frequent struct {
	num      int
	frequent int
}

func FindKMostFrequentElements(nums []int, k int) []int {
	numsFrequent := make([]Frequent, 0)
	numsMap := make(map[int]int)
	for _, item := range nums {
		numsMap[item]++
	}
	for key, val := range numsMap {
		numsFrequent = append(numsFrequent, Frequent{num: key, frequent: val})
	}
	sort.Slice(numsFrequent, func(i, j int) bool {
		return numsFrequent[i].frequent > numsFrequent[j].frequent
	})
	ans := make([]int, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, numsFrequent[i].num)
	}
	return ans
}
