package array

func MaxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	result := 0
	currMin := prices[0]
	for day := 1; day < len(prices); day++ {
		result = max(prices[day]-currMin, result)
		currMin = min(prices[day], currMin)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
