package stack

import (
	"math"
	"strconv"
)

func FindLongestValidParentheses(s string) int {
	st := &Stack{}
	left := -1
	max := 0
	for idx, item := range s {
		if string(item) == "(" {
			st.push(strconv.Itoa(idx))
			continue
		}

		if st.isEmpty() {
			left = idx
			continue
		}

		st.pop()
		if st.isEmpty() {
			max = int(math.Max(float64(max), float64(idx-left)))
			continue
		}

		top, _ := strconv.Atoi(st.top())
		max = int(math.Max(float64(max), float64(idx-top)))
	}

	return max
}
