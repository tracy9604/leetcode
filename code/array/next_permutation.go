package array

func FindNextPermutation(nums []int) {
	n := len(nums)
	var i, j int

	// Find for the pivot element.
	// A pivot is the first element from
	// end of sequence which doesn't follow
	// property of non-increasing suffix
	for i = n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	// if pivot is not found
	if i < 0 {
		for t, k := 0, n-1; t < k; t, k = t+1, k-1 {
			nums[t], nums[k] = nums[k], nums[t]
		}
	} else {
		// Find for the successor of pivot in suffix
		for j = n - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}

		nums[i], nums[j] = nums[j], nums[i]
		for t, k := i+1, n-1; t < k; t, k = t+1, k-1 {
			nums[t], nums[k] = nums[k], nums[t]
		}
	}
}
