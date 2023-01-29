package algorithms

import "strconv"

func IsValidSudoku(board [][]byte) bool {
	grid := make([][]int, 9)
	for i := 0; i < len(board); i++ {
		grid[i] = make([]int, 9)
		for j := 0; j < len(board[i]); j++ {
			if string(board[i][j]) == "." {
				grid[i][j] = 0
				continue
			}
			num, _ := strconv.Atoi(string(board[i][j]))
			grid[i][j] = num
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 && !isSafe(grid, i, j, grid[i][j]) {
				return false
			}
		}
	}
	return true
}

func CheckValidSudoku(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 0 && !isSafe(grid, i, j, grid[i][j]) {
				return false
			}
		}
	}
	return true
}

func isSafe(grid [][]int, row, col, num int) bool {
	// check if found same num in row, return false
	for i := 0; i < 9; i++ {
		if i != col && grid[row][i] == num {
			return false
		}
	}

	// check if found same num in col, return false
	for i := 0; i < 9; i++ {
		if i != row && grid[i][col] == num {
			return false
		}
	}

	// check if found same num in subgrid 3x3, return false
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i+startRow != row && j+startCol != col && grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func SolveSudoku(grid [][]int, row, col int) bool {
	if row == 8 && col == 9 {
		return true
	}

	if col == 9 {
		row++
		col = 0
	}

	if grid[row][col] == 0 {
		return true
	}

	if grid[row][col] > 0 {
		return SolveSudoku(grid, row, col+1)
	}

	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num
			if SolveSudoku(grid, row, col+1) {
				return true
			}
		}
		grid[row][col] = 0
	}
	return false
}
