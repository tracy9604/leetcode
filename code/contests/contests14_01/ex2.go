package contests14_01

import (
	"math"
)

func BeautifulIndices(s string, a string, b string, k int) []int {
	n := len(a)
	m := len(b)
	t := len(s)
	rs := make([]int, 0)

	i, j := 0, 0
	for i <= t-n {
		str1, str2 := "", ""
		if i+n >= t {
			str1 = s[i:]
		} else {
			str1 = s[i:(i + n)]
		}
		if j+m >= t {
			str2 = s[j:]
		} else {
			str2 = s[j:(j + m)]
		}
		if str1 == a && str2 == b && int(math.Abs(float64(j-i))) <= k {
			rs = append(rs, i)
			i++
			j = 0
			continue
		}
		j++
		if j > t-m {
			j = 0
			i++
		}
	}
	return rs
}
