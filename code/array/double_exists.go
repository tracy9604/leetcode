package array

func CheckIfDoubleExists(arr []int) bool {
	tmpMap := make(map[int]int)

	for _, item := range arr {
		if _, found := tmpMap[item*2]; found {
			return true
		}

		if _, found := tmpMap[item/2]; found && item%2 == 0 {
			return true
		}

		tmpMap[item] = 1
	}

	return false
}
