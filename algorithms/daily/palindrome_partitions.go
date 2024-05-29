package daily

func Partition(s string) [][]string {
	stringLength := len(s)
	palindromeTable := make([][]bool, stringLength)

	// Initialize palindrome table
	for i := 0; i < stringLength; i++ {
		palindromeTable[i] = make([]bool, stringLength)
		palindromeTable[i][i] = true // Single character is a palindrome
	}

	// Populate palindrome table with actual information
	for length := 2; length <= stringLength; length++ {
		for i := 0; i <= stringLength-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				if length == 2 {
					palindromeTable[i][j] = true // Two characters are a palindrome if they are equal
				} else {
					palindromeTable[i][j] = palindromeTable[i+1][j-1] // More than two characters
				}
			}
		}
	}

	result := make([][]string, 0)
	tmpSubString := make([]string, 0)

	var dfs func(start int)

	dfs = func(start int) {
		if start == stringLength {
			tmpSubStringCopy := make([]string, len(tmpSubString))
			copy(tmpSubStringCopy, tmpSubString)
			result = append(result, tmpSubStringCopy)
			return
		}

		for end := start; end < stringLength; end++ {
			if palindromeTable[start][end] {
				tmpSubString = append(tmpSubString, s[start:end+1])
				dfs(end + 1)
				tmpSubString = tmpSubString[:len(tmpSubString)-1]
			}
		}
	}

	dfs(0)

	return result

}
