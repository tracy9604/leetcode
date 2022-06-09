package algorithms

import "fmt"

func MergeSort(arr []int) {
	if len(arr) > 1 {
		mid := len(arr) / 2

		left := arr[:mid]
		right := arr[mid:]

		fmt.Println("mid ", mid)

		MergeSort(left)
		MergeSort(right)

		fmt.Println("left ", left)
		fmt.Println("right ", right)

		i, j, k := 0, 0, 0

		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j]
				j++
			}
			k++
		}

		for i < len(left) {
			arr[k] = left[i]
			i++
			k++
		}

		for j < len(right) {
			arr[k] = right[j]
			j++
			k++
		}
	}
}
