package string_problems

import "math"

// leetcode: https://leetcode.com/problems/longest-substring-without-repeating-characters/

func FindLengthOfLongestSubStringBruteForce(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 1
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			sub := s[i : j+1]
			dupMap := make(map[string]int)
			isValid := true
			for k := 0; k < len(sub); k++ {
				if _, found := dupMap[string(sub[k])]; found {
					isValid = false
				}
				dupMap[string(sub[k])]++
			}
			if isValid && len(sub) > max {
				max = len(sub)
			}
		}
	}
	return max
}

func FindLengthOfLongestSubStringWindowSliding(s string) int {
	// find the first window
	max := 0
	for i := 0; i < len(s); i++ {
		visitedMap := make(map[string]bool)
		for j := i; j < len(s); j++ {
			_, found := visitedMap[string(s[j])]
			if found {
				break
			}
			max = int(math.Max(float64(max), float64(j-i+1)))
			visitedMap[string(s[j])] = true
		}
	}
	return max
}
