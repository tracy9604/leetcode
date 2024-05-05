package two_pointers_technique

// refference: https://medium.com/@gokulvaradan/longest-palindromic-substring-40f1ecaede55
func FindLongestPalindromeString(s string) string {
	n := len(s)
	result := ""

	for i := 0; i < n; i++ {
		// odd string
		prev := i - 1
		center := i
		next := i + 1
		oddString := string(s[center])
		for prev >= 0 && next < n && s[prev] == s[next] {
			oddString = string(s[prev]) + oddString
			oddString += string(s[next])
			prev--
			next++
		}

		// even string
		center = i
		next = i + 1
		evenString := ""
		for center >= 0 && next < n && s[center] == s[next] {
			evenString = string(s[center]) + evenString
			evenString += string(s[next])
			center--
			next++
		}

		if len(result) < len(oddString) {
			result = oddString
		}
		if len(result) < len(evenString) {
			result = evenString
		}
	}
	return result
}
