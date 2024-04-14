package binary_search

func findCrossoverNumber(arr []int, left, right, x int) int {
	if arr[right] < x { // x is greater than all
		return right
	}
	if arr[left] > x { // x is smaller than all
		return left
	}

	for left+1 < right {
		mid := left + (right-left)/2
		if arr[mid] <= x && arr[mid+1] > x {
			return mid
		}
		if arr[mid] < x {
			left = mid
		} else {
			right = mid
		}
	}
	if arr[left] <= x {
		return left
	}
	if arr[right] >= x {
		return right
	}
	return -1
}

func FindKClosestElements(arr []int, k int, x int) []int {
	left := findCrossoverNumber(arr, 0, len(arr)-1, x)
	right := left + 1

	if arr[left] == x {
		left--
	}

	count := 0
	ans := make([]int, 0)

	for left >= 0 && right < len(arr) && count < k {
		if x-arr[left] < arr[right]-x {
			ans = append(ans, arr[left])
			left--
		} else {
			ans = append(ans, arr[right])
			right++
		}
		count++
	}

	for count < k && left >= 0 {
		ans = append(ans, arr[left])
		left--
		count++
	}

	for count < k && right < len(arr) {
		ans = append(ans, arr[right])
		right++
		count++
	}

	return ans
}
