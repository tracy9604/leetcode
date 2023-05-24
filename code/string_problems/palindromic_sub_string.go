package string_problems

func FindLongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	longest := ""
	n := len(s)
	for i := 0; i < n; i++ {
		var sub string
		for j := n - 1; j >= i; j-- {
			if string(s[i]) == string(s[j]) && len(longest) < j-i+1 {
				sub = s[i : j+1]
				var reservedSub string
				for _, v := range sub {
					reservedSub = string(v) + reservedSub
				}

				if sub == reservedSub {
					longest = sub
				}
			}
		}
	}

	return longest
}

func FindPalindromeSubstringByDP(s string) string {
	n := len(s)
	table := make([][]bool, n)

	// All substrings of length 1 are palindromes
	maxLength := 1
	for i := 0; i < n; i++ {
		table[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if i == j {
				table[i][j] = true
			} else {
				table[i][j] = false
			}

		}
	}

	start := 0
	// check substrings of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			start = i
			table[i][i+1] = true
			maxLength = 2
		}
	}

	// check substrings of length greater than 2
	for k := 3; k <= n; k++ {
		for i := 0; i < n-k+1; i++ {
			// Get the ending index of substring from i and its length is k
			j := i + k - 1

			// Checking for substring from i to j if s[i:j] is a palindrome
			if table[i+1][j-1] && s[i] == s[j] {
				table[i][j] = true
				if k > maxLength {
					start = i
					maxLength = k
				}
			}
		}
	}

	return s[start : start+maxLength]
}
