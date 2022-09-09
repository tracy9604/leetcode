package sorting

import "sort"

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	arr2Map := make(map[int]int)
	arr1Map := make(map[int]int)
	miss := make([]int, 0)

	for idx, num := range arr2 {
		arr2Map[num] = idx
	}

	for _, num := range arr1 {
		arr1Map[num]++
		if _, found := arr2Map[num]; !found {
			miss = append(miss, num)
		}
	}

	tmp := make([]int, len(arr1)-len(miss))

	for idx, num := range arr2 {
		for i := idx; i < idx+arr2Map[num]; i++ {
			tmp[i] = num
		}
	}

	for num, val := range arr1Map {
		if idx, found := arr2Map[num]; found {
			for i := idx; i < idx+val; i++ {
				tmp[i] = num
			}
		} else {
			miss = append(miss, num)
		}

	}

	sort.Ints(miss)
	result := make([]int, 0)
	result = append(result, tmp...)
	result = append(result, miss...)

	return result
}
