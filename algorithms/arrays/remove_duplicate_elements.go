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
