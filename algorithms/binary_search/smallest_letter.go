package binary_search

func FindSmallestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1

	for left < right {
		mid := left + (right-left)/2

		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}

	}
	if letters[left] > target {
		return letters[left]
	}
	return letters[0]
}
