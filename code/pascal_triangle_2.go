package code

func GeneratePascalTriangleByIndex(rowIndex int) []int {
	rs := make([]int, 0)

	if rowIndex < 0 {
		return rs
	}

	C := 1
	for i := 1; i <= rowIndex+1; i++ {
		rs = append(rs, C)
		C = (C * (rowIndex + 1 - i)) / i
	}

	return rs
}
