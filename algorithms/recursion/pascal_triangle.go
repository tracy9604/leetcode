package recursion

func CalcPascalTriangleRecursion(prev []int, current, target int) []int {
	row := make([]int, current+1)
	row[0], row[len(row)-1] = 1, 1

	for i := 1; i < current; i++ {
		row[i] = prev[i-1] + prev[i]
	}

	if current == target {
		return row
	}

	return CalcPascalTriangleRecursion(row, current+1, target)
}

func CalcPascalTriangle(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	return CalcPascalTriangleRecursion([]int{1, 1}, 2, rowIndex)
}
