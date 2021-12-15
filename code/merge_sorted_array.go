package code

import "fmt"

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, m+n)

	copy(nums3, nums1)

	i := 0
	j := 0
	k := 0

	for i < m && j < n {
		if nums3[i] < nums2[j] {
			nums1[k] = nums3[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		nums1[k] = nums3[i]
		i++
		k++
	}

	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}

	fmt.Println(nums1)
}

func MergeSortedArray2(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	for k := m + n - 1; k >= 0; k-- {
		if (i >= 0 && j >= 0 && nums1[i] > nums2[j]) || (j < 0) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
	fmt.Println(nums1)
}
