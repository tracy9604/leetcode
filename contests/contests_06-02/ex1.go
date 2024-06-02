package contests_06_02

import (
	"fmt"
)

func minimumChairs(s string) int {
	arr := make([]int, len(s))
	idx := -1
	for i := 0; i < len(s); i++ {
		if s[i] == 'E' {
			idx++
			arr[idx] = 1
			continue
		}

		if s[i] == 'L' {
			arr[idx] = 0
			idx--
		}
	}
	return idx
}

func TestMinimumChairs() {
	s := "ELEELEELLL"
	fmt.Println(minimumChairs(s))
}
