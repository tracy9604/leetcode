package bit_manipulation

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	sign := 1
	if dividend*divisor < 0 {
		sign = -1
	}

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	totalQuotient := 0

	for dividend >= divisor {
		count := 0

		for dividend >= divisor<<(count+1) {
			count++
		}

		totalQuotient += 1 << count
		dividend -= divisor << count

	}

	ans := sign * totalQuotient
	INT_MAX := int(math.Pow(2, 31)) - 1
	INT_MIN := -int(math.Pow(2, 31))
	if ans < INT_MIN {
		return INT_MIN
	}

	if ans > INT_MAX {
		return INT_MAX
	}

	return ans

}

func TestDivide() {
	dividend := -2147483648
	divisor := -1
	fmt.Println(divide(dividend, divisor))
}
