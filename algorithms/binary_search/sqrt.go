package binary_search

func SqrtNumber(x int) int {
	if x < 2 {
		return x
	}
	left := 1
	right := x / 2
	mid := (left + right) / 2

	for left <= right {
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}

		mid = (left + right) / 2
	}
	return mid
}
