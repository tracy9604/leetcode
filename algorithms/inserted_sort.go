package algorithms

import "fmt"

func InsertedSort(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}

	n := len(arr)

	for i := 1; i < n; i++ {
		j := i - 1
		m := arr[i]
		for j >= 0 && m < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = m
		fmt.Println(arr)
	}

	return arr
}
