package daily

import (
	"fmt"
	"math"
)

func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	left, right := 0, 0
	currentCost := 0
	maxLength := 0

	for right < n {
		// Calculate the cost to convert s[right] to t[right]
		fmt.Println(byte(s[right] - t[right]))
		currentCost += int(math.Abs(float64(s[right]) - float64(t[right])))

		// While currentCost exceeds maxCost, shrink the window from the left
		for currentCost > maxCost {
			currentCost -= int(math.Abs(float64(s[left]) - float64(t[left])))
			left++
		}

		// Update the maximum length of the valid window
		maxLength = int(math.Max(float64(maxLength), float64(right-left+1)))

		// Expand the window by moving the right pointer
		right++
	}

	return maxLength
}

func TestEqualSubString() {
	s := "abcd"
	t := "bcdf"
	maxCost := 3
	fmt.Println(equalSubstring(s, t, maxCost))
}
