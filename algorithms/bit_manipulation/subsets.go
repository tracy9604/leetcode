package bit_manipulation

import "fmt"

func subsets(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	// Iterate through all possible bit masks
	for mask := 0; mask < 1<<n; mask++ {
		subset := []int{}
		for i := 0; i < n; i++ {
			// Check if element is included based on the mask
			if (mask>>i)&1 == 1 {
				subset = append(subset, nums[i])
			}
		}
		result = append(result, subset)
	}
	return result
}

func TestSubSets() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
