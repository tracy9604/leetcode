package contest10_12

import (
	"fmt"
	"math"
)

func GetGoodIndices(variables [][]int, target int) []int {
	z := math.Pow(51, 35)
	fmt.Println(z)
	n := len(variables)
	rs := make([]int, 0)
	for i := 0; i < n; i++ {
		//a := variables[i][0]
		//b := variables[i][1]
		//c := variables[i][2]
		//m := variables[i][3]

		// 0 <= i < variables.length
		// ((aibi % 10)ci) % mi == target
		z := math.Pow(51, 35)
		fmt.Println(z)
		x := Pow(51, 35)
		y := Pow(x, 21) % 52
		if y == target {
			rs = append(rs, i)
		}
	}
	return rs
}
func Pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * Pow(base, exp-1)
}
