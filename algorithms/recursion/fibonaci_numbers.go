package recursion

func FibonacciNumber(n int, memoMap map[int]int) int {
	if val, found := memoMap[n]; found {
		return val
	}

	if n < 2 {
		return n
	}
	res := FibonacciNumber(n-1, memoMap) + FibonacciNumber(n-2, memoMap)
	memoMap[n] = res
	return res
}

func FibonacciWithMemoization(n int) int {
	memoMap := make(map[int]int)
	return FibonacciNumber(n, memoMap)
}
