package arrays

func DuplicateZeros(arr []int) {
	i := 0
	for i < len(arr)-1 {
		if arr[i] != 0 {
			i++
			continue
		}
		for j := len(arr) - 1; j > i+1; j-- {
			arr[j] = arr[j-1]
		}
		arr[i+1] = 0
		i += 2
	}
}
