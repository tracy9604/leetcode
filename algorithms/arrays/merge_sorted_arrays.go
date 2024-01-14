package arrays

func MergeSortedArrays(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for t := 0; t < n; t++ {
			nums1[t+m] = nums2[t]
		}
		return
	}
	for i := n - 1; i >= 0; i-- {
		j, last := 0, nums1[m-1]
		for j = m - 2; j >= 0 && nums1[j] > nums2[i]; j-- {
			nums1[j+1] = nums1[j]
		}
		if last > nums2[i] {
			nums1[j+1] = nums2[i]
			nums2[i] = last
			nums1[m+i] = last
		} else {
			nums1[m+i] = nums2[i]
		}
	}

	//for t := 0; t < n; t++ {
	//	nums1[t+m] = nums2[t]
	//}
}
