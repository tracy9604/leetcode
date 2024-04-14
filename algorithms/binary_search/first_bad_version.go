package binary_search

func isBadVersion(version int) bool {
	return version == 4
}

func FirstBadVersion(n int) int {
	left := 1
	right := n

	for left < right {
		mid := left + (right-left)/2
		isBad := isBadVersion(mid)
		if isBad {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if isBadVersion(left) {
		return left
	}
	return -1
}
