package daily

import (
	"fmt"
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	freqMap1 := make(map[int]bool)
	for _, item := range arr2 {
		freqMap1[item] = true
	}

	ans := make([]int, len(arr1))
	tmpAns := make([]int, 0)
	freqMap := make(map[int]int)
	for _, item := range arr1 {
		freqMap[item]++
		if _, found := freqMap1[item]; !found {
			tmpAns = append(tmpAns, item)
		}
	}

	index := 0
	for _, item := range arr2 {
		count := freqMap[item]
		for i := 0; i < count; i++ {
			ans[index+i] = item
		}
		index += count
		freqMap[item] = 0
	}
	j := 0
	sort.Ints(tmpAns)
	for index < len(ans) {
		ans[index] = tmpAns[j]
		index++
		j++
	}
	return ans
}

func TestRelativeSortArray() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}
