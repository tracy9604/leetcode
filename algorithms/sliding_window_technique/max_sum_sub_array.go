package sliding_window_technique

import "math"

func FindMaximumSumOfSubArray(arr []int, n, m int) int {
	// compute sum of the first window
	max := 0
	for i := 0; i < m; i++ {
		max += arr[i]
	}

	// compute remaining windows by removing the first element of the previous window and add the last element of the current window
	windowSum := max
	for i := m; i < n; i++ {
		windowSum += arr[i] - arr[i-m]
		max = int(math.Max(float64(max), float64(windowSum)))
	}
	return max
}
