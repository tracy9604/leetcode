package hash_table

import (
	"math"
)

// FindLongestSubstring Naive approach
// + find all substrings
// + save substring into a map then check duplicate
// + if substring is not duplicated, check length and check whether its length is greater than max
func FindLongestSubstring(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	duplicateMap := make(map[string]bool)
	isDuplicate := false
	maxLength := 0
	for i := 0; i < n-1; i++ {
		for j := i; j < n; j++ {
			duplicateMap = make(map[string]bool)
			isDuplicate = false
			for k := i; k <= j; k++ {
				if _, found := duplicateMap[string(s[k])]; !found {
					duplicateMap[string(s[k])] = true
					continue
				}
				isDuplicate = true
				break
			}
			if !isDuplicate {
				maxLength = int(math.Max(float64(maxLength), float64(j-i+1)))
			}
		}
	}
	return maxLength
}

func FindLongestSubstringBySlidingWindow(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}
	result := 0
	duplicateMap := make(map[string]int)
	start := 0
	for end := 0; end < n; end++ {
		if idx, isDuplicate := duplicateMap[string(s[end])]; isDuplicate {
			result = int(math.Max(float64(result), float64(end-start)))

			for i := start; i <= idx; i++ {
				delete(duplicateMap, string(s[i]))
			}

			start = idx + 1
		}
		duplicateMap[string(s[end])] = end
	}
	result = int(math.Max(float64(result), float64(n-start)))
	return result
}
