package sorting

import (
	"sort"
	"strconv"
)

func FindRelativeRank(score []int) []string {
	result := make([]string, len(score))
	tmpMap := make(map[int]int)
	for i := 0; i < len(score); i++ {
		tmpMap[score[i]] = i
	}
	sort.Ints(score)
	for i := len(score) - 1; i >= 0; i-- {
		if i == len(score)-1 {
			result[tmpMap[score[i]]] = "Gold Medal"
		} else if i == len(score)-2 {
			result[tmpMap[score[i]]] = "Silver Medal"
		} else if i == len(score)-3 {
			result[tmpMap[score[i]]] = "Bronze Medal"
		} else {
			result[tmpMap[score[i]]] = strconv.Itoa(len(score) - i)
		}
	}
	return result
}
