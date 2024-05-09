package two_pointers_technique

import "strings"

func checkAlphanumericLetter(letter byte) bool {
	return (letter >= '0' && letter <= '9') || (letter >= 'a' && letter <= 'z')
}

func CheckValidPalindromePhrase(s string) bool {
	// remove all non-alphanumeric letters

	i, j := 0, len(s)-1
	s = strings.ToLower(s)

	for i <= j {
		if !checkAlphanumericLetter(s[i]) {
			i++
			continue
		}
		if !checkAlphanumericLetter(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
