package two_pointers_technique

func FindIndexFirstOccurrenceInAString(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	windowSize := len(needle)
	for i := 0; i < len(haystack)-windowSize+1; i++ {
		subStr := haystack[i:(i + windowSize)]
		if subStr == needle {
			return i
		}
	}

	return -1
}

func FindIndexFirstOccurrenceInAStringV2(haystack string, needle string) int {
	i := 0

	for j := 0; j < len(haystack); j++ {
		if haystack[j] == needle[i] {
			i++
		} else {
			j -= i
			i = 0
		}

		if i == len(needle) {
			return j - i + 1
		}
	}
	return -1
}
