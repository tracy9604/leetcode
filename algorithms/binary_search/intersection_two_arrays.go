package binary_search

import "sort"

func FindIntersectionTwoArrays(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]bool)
	for _, item := range nums1 {
		nums1Map[item] = true
	}

	rs := make([]int, 0)
	rsMap := make(map[int]bool)
	for _, item := range nums2 {
		if _, found := nums1Map[item]; found {
			if _, check := rsMap[item]; !check {
				rs = append(rs, item)
				rsMap[item] = true
			}
		}
	}

	return rs
}

func FindIntersectionTwoArraysV2Search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func FindIntersectionTwoArraysV2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)

	rs := make([]int, 0)
	rsMap := make(map[int]bool)
	for i := 0; i < len(nums2); i++ {
		if FindIntersectionTwoArraysV2Search(nums1, nums2[i]) {
			if _, found := rsMap[nums2[i]]; !found {
				rs = append(rs, nums2[i])
				rsMap[nums2[i]] = true
			}
		}
	}
	return rs
}

func FindIntersection2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return FindIntersection2(nums2, nums1)
	}

	ans := make([]int, 0)
	count := make(map[int]int)

	for _, item := range nums1 {
		count[item]++
	}

	for _, item := range nums2 {
		if val, found := count[item]; found && val > 0 {
			ans = append(ans, item)
			count[item]--
		}
	}
	return ans
}
