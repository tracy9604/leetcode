package recursion

import "fmt"

func ReserveString(s []byte) {
	if len(s) <= 1 {
		return
	}
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reserveStringRecursionHelper(s []byte, length int) string {
	if length < 1 {
		return ""
	}
	if length == 1 {
		return string(s[0])
	}
	return string(s[length-1]) + reserveStringRecursionHelper(s, length-1)
}

func ReserveStringRecursion(s []byte) {
	tmpS := reserveStringRecursionHelper(s, len(s))
	fmt.Println(tmpS)
	for i := 0; i < len(tmpS); i++ {
		s[i] = tmpS[i]
	}
	fmt.Println(string(s))
}
