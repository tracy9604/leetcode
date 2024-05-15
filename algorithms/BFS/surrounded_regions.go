package BFS

type Cell struct {
	row int
	col int
}

func SurroundedRegions(board [][]byte) {
	rows := len(board)
	cols := len(board[0])

	// mark all no flipped cell O
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (i == 0 || i == rows-1 || j == 0 || j == cols-1) && board[i][j] == 'O' {
				markUnFlippedCellDFS(board, i, j, rows, cols)
			}
		}
	}

	// flip other O cell and convert no flipped cell to O
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '.' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

func markUnFlippedCellDFS(board [][]byte, row, col, rows, cols int) {
	rRow := []int{1, 0, -1, 0}
	cCol := []int{0, 1, 0, -1}
	board[row][col] = '.'

	for i := 0; i < 4; i++ {
		nextRow := row + rRow[i]
		nextCol := col + cCol[i]
		if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && board[nextRow][nextCol] == 'O' {
			markUnFlippedCellDFS(board, nextRow, nextCol, rows, cols)
		}
	}

}

func popCellQueue(queue []Cell) Cell {
	if len(queue) == 0 {
		return Cell{}
	}
	top := queue[0]
	if len(queue) == 1 {
		queue = []Cell{}
		return top
	}
	queue = queue[1:]
	return top
}

func markUnFlippedCellBFS(board [][]byte, row, col, rows, cols int) {
	rRow := []int{1, 0, -1, 0}
	cCol := []int{0, 1, 0, -1}

	queue := make([]Cell, 0)
	queue = append(queue, Cell{row: row, col: col})

	for len(queue) > 0 {
		top := popCellQueue(queue)
		board[top.row][top.col] = '.'
		for i := 0; i < 4; i++ {
			nextRow := top.row + rRow[i]
			nextCol := top.col + cCol[i]

			if nextRow >= 0 && nextRow < rows && nextCol >= 0 && nextCol < cols && board[nextRow][nextCol] == 'O' {
				queue = append(queue, Cell{row: nextRow, col: nextCol})
			}
		}
	}
}
