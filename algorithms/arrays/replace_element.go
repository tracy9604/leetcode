package arrays

func ReplaceElements(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		max := arr[i+1]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > max {
				max = arr[j]
			}
		}
		arr[i] = max
	}
	arr[len(arr)-1] = -1
	return arr
}
