package stack

import "strconv"

func pop2ItemsStack(stack []int) (int, int, []int) {
	if len(stack) <= 1 {
		return -1, -1, stack
	}
	top1, top2 := stack[len(stack)-1], stack[len(stack)-2]
	stack = stack[:len(stack)-2]
	return top1, top2, stack
}

func EvaluatePolishNotation(tokens []string) int {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	stack := make([]int, 0)
	for _, item := range tokens {
		if _, found := operators[item]; found {
			top1, top2, tmpStack := pop2ItemsStack(stack)
			stack = tmpStack
			switch item {
			case "+":
				stack = append(stack, top1+top2)
			case "-":
				stack = append(stack, top2-top1)
			case "*":
				stack = append(stack, top1*top2)
			case "/":
				stack = append(stack, top2/top1)
			}
			continue
		}
		numItem, _ := strconv.Atoi(item)
		stack = append(stack, numItem)
	}
	if len(stack) == 0 {
		return -1
	}
	return stack[0]
}
