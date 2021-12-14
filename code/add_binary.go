package code

func AddBinary(a string, b string) string {
	result := ""
	carry := 0

	i := len(a) - 1
	j := len(b) - 1

	for i >= 0 || j >= 0 {
		r := carry
		if i >= 0 && string(a[i]) == "1" {
			r += 1
		}

		if j >= 0 && string(b[j]) == "1" {
			r += 1
		}

		if r%2 == 1 {
			result = "1" + result
		} else {
			result = "0" + result
		}

		if r < 2 {
			carry = 0
		} else {
			carry = 1
		}
		i--
		j--
	}

	if carry != 0 {
		result = "1" + result
	}

	return result
}
