package recursion

func ClimbingStairsRecursion(curSum, target int) int {
	if curSum == target {
		return 1
	}

	if curSum > target {
		return 0
	}

	sum1 := ClimbingStairsRecursion(curSum+1, target)
	sum2 := ClimbingStairsRecursion(curSum+2, target)

	return sum1 + sum2
}

func ClimbingStair(n int) int {
	return ClimbingStairsRecursion(0, n)
}
