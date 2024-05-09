package two_pointers_technique

func ReverseWords(s string) string {
	start, end := 0, 0
	n := len(s)
	ans := ""

	for start < n {
		// skip space until find next character
		for start < n && s[start] == ' ' {
			start++
		}

		if start < n {
			// put space before the next word
			if end != 0 {
				ans += " "
				end++
			}
		}
	}

	return ans
}
