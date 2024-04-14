package binary_search

func FindTwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
			continue
		}
		if numbers[left]+numbers[right] > target {
			right--
			continue
		}
	}
	return []int{-1, -1}
}

func FindTwoSumBS(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left+1 < right {
		mid := left + (right-left)/2

		if numbers[mid]+numbers[right] == target {
			return []int{mid + 1, right + 1}
		}

		if numbers[mid]+numbers[right] > target {
			right = mid
		} else {
			left = mid
		}

	}
	if numbers[left]+numbers[right] == target {
		return []int{left + 1, right + 1}
	}
	return []int{-1, -1}
}
