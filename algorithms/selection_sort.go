package algorithms

func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		m := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[m] {
				tmp := arr[m]
				arr[m] = arr[j]
				arr[j] = tmp
			}
		}
	}

	return arr
}

func SelectionSortV2(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}
