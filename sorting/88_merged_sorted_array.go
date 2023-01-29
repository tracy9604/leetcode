package sorting

import "fmt"

func MergeSortedArray(nums1, nums2 []int, m, n int) {
	i := m - 1
	j := n - 1

	for k := m + n - 1; k >= 0; k-- {
		if (i > 0 && j >= 0 && nums1[i] > nums2[j]) || j < 0 {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	fmt.Println(nums1)
}
