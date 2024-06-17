package daily

import "fmt"

func appendCharacters(s string, t string) int {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			j++
		}
		i++
	}

	return len(t) - j
}

func TestAppendCharacters() {
	s := "z"
	t := "abcde"
	fmt.Println(appendCharacters(s, t))
}
