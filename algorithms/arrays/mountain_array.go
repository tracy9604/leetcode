package arrays

func FindMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	max := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}

	if max == len(arr)-1 || max == 0 {
		return false
	}

	for i := 0; i < max; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}

	for j := max; j < len(arr)-1; j++ {
		if arr[j] <= arr[j+1] {
			return false
		}
	}

	return true
}
