package sorting

func FindThirdMaximumNumber(nums []int) int {
	max, secondMax, thirdMax := -1<<31, -1<<31, -1<<31
	tmpMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		// remove element if it's duplicated
		if _, check := tmpMap[nums[i]]; check {
			continue
		}

		if nums[i] > max {
			thirdMax = secondMax
			secondMax = max
			max = nums[i]
		} else if nums[i] > secondMax {
			thirdMax = secondMax
			secondMax = nums[i]
		} else if nums[i] > thirdMax {
			thirdMax = nums[i]
		}
		tmpMap[nums[i]] = true
	}

	if len(tmpMap) < 3 {
		return max
	}

	return thirdMax
}
