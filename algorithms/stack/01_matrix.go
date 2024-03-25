package stack

import "math"

func UpdateMatrixRecursion(mat [][]int, row, col int, visited map[Cell]bool) int {
	if val, found := visited[Cell{row: row, col: col}]; (found && val) || !isSafeCell(mat, row, col) {
		return 1000
	}
	if mat[row][col] == 0 {
		return 0
	}
	visited[Cell{row: row, col: col}] = true

	adjacentCell := 1 + int(math.Min(float64(UpdateMatrixRecursion(mat, row+1, col, visited)),
		math.Min(float64(UpdateMatrixRecursion(mat, row-1, col, visited)), math.Min(float64(UpdateMatrixRecursion(mat, row, col+1, visited)), float64(UpdateMatrixRecursion(mat, row, col-1, visited))))))

	visited[Cell{row: row, col: col}] = false
	return adjacentCell
}

func BFSUpdateMatrix(mat [][]int) [][]int {
	queue := make([]Cell, 0)

	rRow := []int{-1, 0, 1, 0}
	cCol := []int{0, 1, 0, -1}

	rs := make([][]int, 0)
	for i := 0; i < len(mat); i++ {
		tmpRs := make([]int, 0)
		for j := 0; j < len(mat[i]); j++ {
			tmpRs = append(tmpRs, mat[i][j])
		}
		rs = append(rs, tmpRs)
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			rs[i][j] = math.MaxInt

			if mat[i][j] == 0 {
				rs[i][j] = 0
				queue = append(queue, Cell{row: i, col: j})
			}
		}
	}

	for len(queue) > 0 {
		top := queue[0]
		if len(queue) <= 0 {
			queue = make([]Cell, 0)
		} else {
			queue = queue[1:]
		}

		for i := 0; i < 4; i++ {
			r := top.row + rRow[i]
			c := top.col + cCol[i]
			if isSafeCell(mat, r, c) && rs[r][c] > rs[top.row][top.col]+1 {
				rs[r][c] = rs[top.row][top.col] + 1
				queue = append(queue, Cell{row: r, col: c})
			}
		}
	}
	return rs
}

func UpdateMatrix(mat [][]int) [][]int {
	visited := make(map[Cell]bool)
	rs := mat
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				rs[i][j] = UpdateMatrixRecursion(mat, i, j, visited)
			}
		}
	}
	return rs
}

func UpdateMatrixV2(mat [][]int) [][]int {
	return BFSUpdateMatrix(mat)
}
