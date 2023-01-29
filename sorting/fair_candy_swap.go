package sorting

func FairCandySwap(aliceSizes []int, bobSizes []int) []int {
	var bobSum, aliceSum int
	bobMap := make(map[int]int)
	for i := 0; i < len(aliceSizes); i++ {
		aliceSum += aliceSizes[i]
	}
	for j := 0; j < len(bobSizes); j++ {
		bobMap[bobSizes[j]]++
		bobSum += bobSizes[j]
	}
	result := make([]int, 2)

	for i := 0; i < len(aliceSizes); i++ {
		y := aliceSizes[i] + (bobSum-aliceSum)/2
		if _, exist := bobMap[y]; exist {
			result[0] = aliceSizes[i]
			result[1] = y
			return result
		}
	}

	return result
}
