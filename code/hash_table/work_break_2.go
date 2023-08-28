package hash_table

func FindWorkBreak2(s string, wordDict []string) []string {
	output := make([]string, 0)
	wordBreak(s, "", wordDict, &output)
	return output
}

func wordBreak(s string, result string, wordDict []string, output *[]string) {
	n := len(s)
	for i := 1; i <= n; i++ {
		prefix := s[:i]
		if validate(prefix, wordDict) {
			if i == n {
				result += prefix
				*output = append(*output, result)
				return
			}
			wordBreak(s[i:], result+prefix+" ", wordDict, output)
		}
	}
}

func validate(word string, wordDict []string) bool {
	for _, item := range wordDict {
		if word == item {
			return true
		}
	}
	return false
}
