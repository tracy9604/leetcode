package stack

import "strings"

func SolveDecodeString(s string) string {
	st1 := make([]int, 0)
	st2 := make([]string, 0)
	for i := 0; i < len(s); i++ {
		count := 0
		if s[i] >= '0' && s[i] <= '9' {

			for s[i] >= '0' && s[i] <= '9' {
				count = count*10 + int(s[i]-'0')
				i++
			}
			i--
			st1 = append(st1, count)
			continue
		}

		if s[i] == ']' {
			count = st1[len(st1)-1]
			st1 = st1[:len(st1)-1]

			tmpSubStr := ""
			for len(st2) > 0 && string(st2[len(st2)-1]) != "[" {
				tmpSubStr = st2[len(st2)-1] + tmpSubStr
				st2 = st2[:len(st2)-1]
			}
			if len(st2) > 0 && string(st2[len(st2)-1]) == "[" {
				st2 = st2[:len(st2)-1]
			}
			tmpRs := ""
			for j := 0; j < count; j++ {
				tmpRs += tmpSubStr
			}
			st2 = append(st2, tmpRs)
			continue
		}

		st2 = append(st2, string(s[i]))

	}
	return strings.Join(st2, "")
}
