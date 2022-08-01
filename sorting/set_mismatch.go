package sorting

func FindSetMismatch(nums []int) []int {
	tmpMap := make(map[int]int)
	result := make([]int, 2)
	for _, num := range nums {
		tmpMap[num]++
		if tmpMap[num] > 1 {
			result[0] = num
		}
	}

	for i := 0; i < len(nums); i++ {
		_, exist := tmpMap[i+1]
		if !exist {
			result[1] = i + 1
			break
		}
		if tmpMap[i+1] > 1 {
			result[0] = i + 1
		}
	}

	return result
}
