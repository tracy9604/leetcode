package hash_table

func ConvertIntegerToRoman(num int) string {
	result := ""
	nums := [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	symbols := [13]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	for i := 12; i >= 0; i-- {
		div := num / nums[i]
		num %= nums[i]
		for j := div; j > 0; j-- {
			result += symbols[i]
		}
		if num == 0 {
			break
		}
	}
	return result
}

func ConvertRomanToInteger(s string) int {
	symbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	last := 0
	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		tmp := symbols[string(s[i])]
		if tmp < last {
			res -= tmp
		} else {
			res += tmp
		}

		last = tmp
	}

	return res
}
