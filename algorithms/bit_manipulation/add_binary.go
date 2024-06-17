package bit_manipulation

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	i, j, carry := len(a)-1, len(b)-1, 0
	ans := make([]string, 0)

	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			carry += int(b[j] - '0')
			j--
		}

		ans = append([]string{strconv.Itoa(carry % 2)}, ans...)
		carry /= 2
	}

	return strings.Join(ans, "")
}

func TestAddBinary() {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}
