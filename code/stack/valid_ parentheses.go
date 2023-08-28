package stack

func ValidParentheses(s string) bool {
	st := &Stack{}

	dataMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	for _, item := range s {
		if string(item) == "(" || string(item) == "{" || string(item) == "[" {
			st.push(string(item))
			continue
		}

		element, check := st.pop()
		if !check {
			st.push(string(item))
			continue
		}
		data, found := dataMap[element]
		if found && string(item) == data {
			continue
		} else {
			return false
		}
	}

	return st.isEmpty()
}
