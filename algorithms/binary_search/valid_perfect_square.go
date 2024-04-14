package binary_search

func FindValidPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	left := 1
	right := num / 2

	for left <= right {
		mid := left + (right-left)/2

		if mid*mid == num {
			return true
		}

		if mid*mid < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
