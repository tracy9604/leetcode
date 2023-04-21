package array

func LengthOfLastWord(s string) int {
	arr := make([]string, 0)
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 32 {
			str = ""
		} else {
			str += string(s[i])
		}

		if len(str) > 0 {
			arr = append(arr, str)
		}

	}
	if len(arr) > 0 {
		return len(arr[len(arr)-1])
	}
	return 0
}

func LengthOfLastWord2(s string) int {
	count := 0
	for i := len(s) - 1; i > -1; i-- {
		if string(s[i]) != " " {
			count++
		} else if count > 0 {
			return count
		}
	}
	return count
}
