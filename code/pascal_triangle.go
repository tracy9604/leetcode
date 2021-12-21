package code

func GeneratePascalTriangle(numRows int) [][]int {
	rs := make([][]int, 0)

	if numRows <= 0 {
		return rs
	}

	for line := 1; line <= numRows; line++ {
		C := 1
		arr := make([]int, line)
		for i := 1; i <= line; i++ {
			arr[i-1] = C
			C = (C * (line - i)) / i
		}
		rs = append(rs, arr)
	}

	return rs
}
