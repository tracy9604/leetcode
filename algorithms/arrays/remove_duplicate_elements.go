package arrays

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 1 {
		return n
	}

	j := 0
	for i := 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}
	}
	nums[j] = nums[n-1]
	return j + 1
}

func RemoveDuplicatesV2(nums []int) int {
	numsMap := make(map[int]int)
	copiedNums := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if _, found := numsMap[nums[i]]; !found {
			numsMap[nums[i]]++
			copiedNums = append(copiedNums, nums[i])
		}
	}
	for i := 0; i < len(copiedNums); i++ {
		nums[i] = copiedNums[i]
	}
	return len(copiedNums)
}

func RemoveDuplicatedElementsV3(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		for nums[left] != val && left < right {
			left++
		}
		for nums[right] == val && left < right {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	return len(nums) - count
}
