package algorithms

func QuickSort(arr []int, start, end int) []int {
	if start < end {
		pivot := partition(arr, start, end)
		QuickSort(arr, start, pivot-1)
		QuickSort(arr, pivot+1, end)
	}
	return arr
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	left := start
	right := end - 1

	for {
		for left <= right && arr[left] < pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}

		if left >= right {
			break
		}

		arr[left], arr[right] = arr[right], arr[left]

		left++
		right--
	}
	arr[left], arr[end] = arr[end], arr[left]
	return left
}
