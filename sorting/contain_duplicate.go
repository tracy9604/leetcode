package sorting

func FindDuplicate(nums []int) bool {
	tmpMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		tmpMap[nums[i]]++
		val, check := tmpMap[nums[i]]
		if check && val >= 1 {
			return true
		}
	}
	return false
}
