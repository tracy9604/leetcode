package algorithms

func QuickSort(arr []int, low, high int) []int {
	if low < high {
		pi := partition(arr, low, high)
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}

	return arr
}

/* This function takes last element as pivot, places
   the pivot element at its correct position in sorted
    array, and places all smaller (smaller than pivot)
   to left of pivot and all greater elements to right
   of pivot */
func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1 // to i is 0

	for j := low; j < high; j++ {
		// If current element is smaller than the pivot, increment index of smaller element
		// swap element in i with j
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// i index is element is smaller than pivot => i + 1 is greater than pivot
	// so swap i +1 element and pivot
	// smaller element will before pivot
	// greater element will after pivot
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
