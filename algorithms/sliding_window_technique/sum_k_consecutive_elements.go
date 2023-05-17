package sliding_window_technique

import "math"

func FindSumConsecutiveElementsByBF(arr []int, k int) int {
	max := 0
	for i := 0; i < len(arr)-k+1; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += arr[j]
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

func FindSumConsecutiveElementsBySlidingWindow(arr []int, k int) int {
	// compute sum of first window of size k
	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	// compute sums of remaining windows by removing first element of previous window and adding last element of current window
	windowSum := maxSum
	for i := k; i < len(arr); i++ {
		windowSum += arr[i] - arr[i-k]
		maxSum = int(math.Max(float64(maxSum), float64(windowSum)))
	}
	return maxSum
}
