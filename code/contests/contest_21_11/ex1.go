package contest_21_11

import "strings"

func FindSubstring(words []string, x byte) []int {
	rs := make([]int, 0)
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], string(x)) {
			rs = append(rs, i)
		}
	}
	return rs
}
