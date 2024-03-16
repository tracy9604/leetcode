package stack

func popStack(s []string) ([]string, string) {
	if len(s) == 0 {
		return s, ""
	}
	top := s[len(s)-1]
	s = s[:len(s)-1]
	return s, top
}

func CheckValidParentheses(s string) bool {
	stack := make([]string, 0)

	data := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	for _, item := range s {
		if string(item) == "(" || string(item) == "[" || string(item) == "{" {
			stack = append(stack, string(item))
			continue
		}

		tmpStack, top := popStack(stack)
		stack = tmpStack
		if val, found := data[string(item)]; found && val == top {
			continue
		}

		return false

	}
	return len(stack) == 0
}
