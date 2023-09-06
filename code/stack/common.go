package stack

type Stack []string

func (s *Stack) push(data string) {
	*s = append(*s, data)
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	}
	lastIndex := len(*s) - 1
	element := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return element, true
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) top() string {
	return (*s)[len(*s)-1]
}
