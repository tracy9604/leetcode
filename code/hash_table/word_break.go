package hash_table

import "fmt"

func CheckWorkBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]int)
	for i := 0; i < len(wordDict); i++ {
		wordMap[wordDict[i]]++
	}

	//dp := make([]bool, len(s)+1)
	//dp[0] = true
	//
	//for i := 0; i <= len(s); i++ {
	//	for j := 0; j < i; j++ {
	//		if _, found := wordMap[s[j:i]]; found && dp[j] {
	//			dp[i] = true
	//		}
	//	}
	//}
	//return dp[len(s)]

	return checkEachWordRecursion(s, wordMap)
}

func checkEachWord(s string, wordMap map[string]int, start int, cache *map[int]bool) bool {
	if val, found := (*cache)[start]; found && val {
		return true
	}
	if len(s) == start {
		return true
	}
	temp := ""
	ret := false
	for i := start; i < len(s); i++ {
		if i+1 >= len(s) {
			temp = s[start:]
		} else {
			temp = s[start : i+1]
		}
		if _, found := wordMap[temp]; found && checkEachWord(s, wordMap, i+1, cache) {
			ret = true
			break
		}
	}
	(*cache)[start] = ret
	return ret
}

func checkEachWordRecursion(s string, wordMap map[string]int) bool {
	fmt.Println(s)
	if len(s) == 0 {
		return true
	}

	for i := 1; i <= len(s); i++ {
		val := s[:i]
		fmt.Println(val)
		if _, found1 := wordMap[s[:i]]; found1 && checkEachWordRecursion(s[i:], wordMap) {
			return true
		}
	}
	return false
}
