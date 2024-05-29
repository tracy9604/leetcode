package daily

import "fmt"

func findSpecialArray(nums []int) int {
	for k := 0; k <= len(nums); k++ {
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] >= k {
				count++
			}
		}
		if count == k {
			return k
		}
	}

	return -1

}

func TestFindSpecialArray() {
	nums := []int{3, 5}
	fmt.Println(findSpecialArray(nums))
}
