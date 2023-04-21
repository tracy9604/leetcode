package array

import "fmt"

func countSquareNum(n int) int {
	rs := 0
	for n > 0 {
		rs += (n % 10) * (n % 10)
		n /= 10
	}
	return rs
}

func CheckHappyNumber(n int) bool {
	slow, fast := n, n
	for {
		slow = countSquareNum(slow)
		fast = countSquareNum(countSquareNum(fast))

		fmt.Println("slow: ", slow)
		fmt.Println("fast: ", fast)

		if slow == fast {
			break
		}
	}

	return slow == 1
}
