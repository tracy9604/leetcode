package two_pointers_technique

func FindIntersectionTwoArrays2(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]int)

	for _, item := range nums1 {
		nums1Map[item]++
	}

	ans := make([]int, 0)
	for _, item := range nums2 {
		if val, found := nums1Map[item]; found && val > 0 {
			ans = append(ans, item)
			nums1Map[item]--
		}
	}

	return ans
}
