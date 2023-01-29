package sorting

func FindMajorityElement(nums []int) int {
	votes := 0
	major := -1
	for i := 0; i < len(nums); i++ {
		if votes == 0 {
			major = nums[i]
			votes = 1
		} else if major != nums[i] {
			votes--
		} else {
			votes++
		}
	}
	return major
}
