package sorting

func FindIntersectionArray(nums1, nums2 []int) []int {
	tmpMap := make(map[int]int)
	rs := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		tmpMap[nums1[i]] = 0
	}

	for i := 0; i < len(nums2); i++ {
		if val, check := tmpMap[nums2[i]]; check && val < 1 {
			rs = append(rs, nums2[i])
			tmpMap[nums2[i]]++
		}
	}
	return rs
}

func FindIntersectionArrayV2(nums1, nums2 []int) []int {
	tmpMap := make(map[int]int)
	tmpMap2 := make(map[int]int)
	rs := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		tmpMap[nums1[i]] = 0
	}

	for i := 0; i < len(nums2); i++ {
		if _, check := tmpMap[nums2[i]]; check {
			if _, exist := tmpMap2[nums2[i]]; !exist {
				rs = append(rs, nums2[i])
				tmpMap2[nums2[i]]++
			}
		}
	}
	return rs
}
