package contests14_01

func MaxFrequencyElements(nums []int) int {
	numsMap := make(map[int]int)
	max := 0
	rs := 0
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
		if val, _ := numsMap[nums[i]]; val > max {
			max = val
		}
	}
	for _, val := range numsMap {
		if val == max {
			rs += val
		}
	}
	return rs
}
