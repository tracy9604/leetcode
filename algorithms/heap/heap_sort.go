package heap

import "fmt"

func heapify(arr []int, n, i int) {
	largest := i
	left := (2 * i) + 1
	right := (2 * i) + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]

		heapify(arr, n, largest)
	}

}
func heapSorting(arr []int) {
	n := len(arr)

	// build heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func TestHeapSorting() {
	arr := []int{12, 11, 13, 5, 6, 7}
	heapSorting(arr)

	fmt.Print(arr)
}
