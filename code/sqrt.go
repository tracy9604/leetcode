package code

func FindSqrt(x int) int {

	if x <= 0 {
		return 0
	}

	left := 0
	right := x
	mid := (left + right) / 2

	for left <= right {
		mid = (left + right) / 2
		mud := mid * mid

		if mud < x {
			left = mid + 1
		} else if mud == x {
			return mid
		} else {
			right = mid - 1
		}
	}

	return right
}
