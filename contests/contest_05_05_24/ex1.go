package contest_05_05_24

import (
	"strconv"
	"strings"
)

func Ex1(word string) bool {
	if len(word) < 3 {
		return false
	}

	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	isValidVowel := false
	isValidConsonant := false
	for _, letter := range word {
		_, err := strconv.Atoi(string(letter))
		if err == nil {
			continue
		}
		if (letter < 'a' || letter > 'z') && (letter < 'A' || letter > 'Z') {
			return false
		}

		if _, found := vowels[strings.ToLower(string(letter))]; found {
			isValidVowel = true
		} else {
			isValidConsonant = true
		}

	}

	return isValidVowel && isValidConsonant
}
