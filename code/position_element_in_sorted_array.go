package code

//func FindPositionOfElementInSortedArray(nums []int, target int) []int {
//	low := 0
//	high := len(nums) - 1
//	mid := (low + high) / 2
//
//}

func Search(nums []int, low, high, target int) {
	arr := make([]int, 0)
	for low < high {
		mid := (low + high) / 2
		if nums[mid] == target {
			arr = append(arr, mid)
			low = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

}
