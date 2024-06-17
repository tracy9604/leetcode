package daily

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)
	visitedMap := make(map[int]bool)
	count := 0
	for i := 0; i < len(hand)-1; i++ {
		if hand[i] < hand[i+1] {
			count++
			visitedMap[i] = true
		}
	}
	return false
}
