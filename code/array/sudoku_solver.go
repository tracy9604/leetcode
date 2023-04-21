package array

import "strconv"

func transform(board [][]byte) [][]int {
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
	return grid
}

func findEmptyLocation(grid [][]int, row, col int) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

func usedInRow(grid [][]int, row, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return true
		}
	}
	return false
}

func usedInCol(grid [][]int, col, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return true
		}
	}
	return false
}

func usedInBlock(grid [][]int, row, col, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+row][j+col] == num {
				return true
			}
		}
	}
	return false
}

func isLocationIsSafe(grid [][]int, row, col, num int) bool {
	return !usedInRow(grid, row, num) && usedInCol(grid, col, num) && !usedInBlock(grid, row-row%3, col-col%3, num)
}

func SolveSudoku(board [][]byte) {
	//grid := transform(board)
	//row := 0
	//col := 0
	for i := 0; i < 9; i++ {

	}
}

func isValid(grid [][]int, row, col, num int) bool {
	// check if we find similar value in row
	for i := 0; i < 9; i++ {

	}
	return true
}
