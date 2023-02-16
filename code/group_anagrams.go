package code

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	strHash := make(map[string][]string)
	for _, str := range strs {
		temp := make([]string, 0)
		for _, char := range str {
			temp = append(temp, string(char))
		}
		sort.Strings(temp)
		key := strings.Join(temp, "")
		if _, found := strHash[key]; !found {
			strHash[key] = make([]string, 0)
		}
		strHash[key] = append(strHash[key], str)
	}

	rs := make([][]string, 0)
	for _, val := range strHash {
		rs = append(rs, val)
	}
	return rs
}
