package sorting

import "math"

// FindLHS https://massivealgorithms.blogspot.com/search?q=Longest+Harmonious+Subsequence
func FindLHS(nums []int) int {
	tmpMap := make(map[int]int)
	for _, num := range nums {
		tmpMap[num]++
	}

	var result float64
	for key, val1 := range tmpMap {
		if val2, exist := tmpMap[key+1]; exist {
			result = math.Max(result, float64(val2+val1))
		}
	}
	return int(result)
}
