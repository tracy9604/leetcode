package array

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	n := len(needle)

	for i := 0; i <= len(haystack)-n; i++ {
		tmpStr := haystack[i : i+n]
		if tmpStr == needle {
			return i
		}
	}

	return -1
}
