package two_pointers_technique

import "sort"

func AssignCookies(g []int, s []int) int {
	i, j := 0, 0
	sort.Ints(g)
	sort.Ints(s)

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			j++
			i++
			continue
		}
		if g[i] > s[j] {
			j++
		}
	}

	return i
}
