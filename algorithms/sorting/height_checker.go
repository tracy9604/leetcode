package sorting

func HeightChecker(heights []int) int {
	heightsMap := make(map[int]int)
	for i := 0; i < len(heights); i++ {
		heightsMap[i] = heights[i]
	}

	//sort heights by bubble sorting
	hasSwapped := true
	for hasSwapped {
		hasSwapped = false
		for i := 0; i < len(heights)-1; i++ {
			if heights[i] > heights[i+1] {
				hasSwapped = true
				heights[i], heights[i+1] = heights[i+1], heights[i]
			}
		}
	}

	ans := 0
	for i := 0; i < len(heights); i++ {
		if val, found := heightsMap[i]; found && val != heights[i] {
			ans++
		}
	}
	return ans
}
