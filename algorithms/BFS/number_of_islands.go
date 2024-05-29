package BFS

func NumberOfIslands(grid [][]byte) int {
	visited := make(map[Cell]bool)
	count := 0
	R := len(grid)
	C := len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			cell := Cell{row: i, col: j}
			if _, found := visited[cell]; !found && grid[i][j] == '1' {
				numberOfIslandDFS(grid, cell, visited, R, C)
				count++
			}
		}
	}
	return count
}

func numberOfIslandDFS(grid [][]byte, cell Cell, visited map[Cell]bool, R, C int) {
	if _, found := visited[cell]; found {
		return
	}
	visited[cell] = true
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}

	for i := 0; i < 4; i++ {
		rRow := cell.row + row[i]
		cCol := cell.col + col[i]
		cCell := Cell{row: rRow, col: cCol}
		if _, found := visited[cCell]; !found && rRow >= 0 && rRow < R && cCol >= 0 &&
			cCol < C && grid[rRow][cCol] == '1' {
			numberOfIslandDFS(grid, cCell, visited, R, C)
		}
	}
}

func numberOfIslandsBFS(grid [][]byte, cell Cell, visited map[Cell]bool) {
	row := []int{1, 0, -1, 0}
	col := []int{0, 1, 0, -1}

	queue := make([]Cell, 0)
	queue = append(queue, cell)

	R := len(grid)
	C := len(grid[0])

	visited[cell] = true

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			rRow := top.row + row[i]
			cCol := top.col + col[i]
			cCell := Cell{row: rRow, col: cCol}

			if _, found := visited[cCell]; !found && rRow >= 0 && rRow < R && cCol >= 0 &&
				cCol < C && grid[rRow][cCol] == '1' {
				visited[cCell] = true
				queue = append(queue, cCell)
			}
		}
	}
}
