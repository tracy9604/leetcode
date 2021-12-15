package code

func ClimbStairs(n int) int {
	if n < 3 {
		return n
	}

	sum, num := 2, 1

	for i := 2; i < n; i++ {
		num, sum = sum, sum+num
	}

	return sum

}
