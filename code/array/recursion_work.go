package array

import "fmt"

func CheckRecursionWork(n int) int {
	if n == 1 {
		return 1
	}

	x := n * CheckRecursionWork(n-1)
	fmt.Println("x: ", x)
	return x
}
