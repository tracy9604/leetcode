package array

func ValidParentheses(s string) bool {
	arr := make([]string, 0)
	a := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	for i := 0; i < len(s); i++ {
		char := string(s[i])
		if _, check := a[char]; check {
			arr = append(arr, char)
		} else if (char == ")" && arr[len(arr)-1] == "(") || (char == "]" && arr[len(arr)-1] == "[") || char == "}" && arr[len(arr)-1] == "{" {
			arr = arr[:(len(arr) - 1)]
		}
	}
	return len(arr) == 0
}
