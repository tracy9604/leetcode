package stack

func popIntStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}
	return stack[:len(stack)-1]
}

func topIntStack(stack []int) int {
	if len(stack) == 0 {
		return -1
	}
	return stack[len(stack)-1]
}

func PredictDailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(temperatures))

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[topIntStack(stack)] < temperatures[i] {
			idx := topIntStack(stack)
			stack = popIntStack(stack)
			ans[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return ans
}
