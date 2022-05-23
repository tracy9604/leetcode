package algorithms

// BubbleSort https://www.geeksforgeeks.org/bubble-sort/
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func RecursiveBubbleSort(arr []int, n int) {
	if n > 1 {
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		RecursiveBubbleSort(arr, n-1)
	}
}
