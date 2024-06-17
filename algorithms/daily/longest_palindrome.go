package daily

import (
	"fmt"
)

func longestPalindrome(s string) int {
	freqMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freqMap[s[i]]++
	}
	maxLength := 0
	for _, count := range freqMap {
		maxLength += count - (count % 2)

		if maxLength%2 == 0 && count%2 == 1 {
			maxLength++
		}
	}
	return maxLength
}

func TestLongestPalindrome() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
