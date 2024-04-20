package sorting

func HeapSort(nums []int) []int {
	for i := len(nums)/2 - 1; i >= 0; i-- {
		maxHeapify(nums, len(nums), i)
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapify(nums, i, 0)
	}
	return nums
}

func maxHeapify(nums []int, heapSize, index int) {
	left := 2*index + 1
	right := 2*index + 2
	largest := index

	if left < heapSize && nums[left] > nums[largest] {
		largest = left
	}
	if right < heapSize && nums[right] > nums[largest] {
		largest = right
	}

	if largest != index {
		nums[largest], nums[index] = nums[index], nums[largest]
		maxHeapify(nums, heapSize, largest)
	}

}
