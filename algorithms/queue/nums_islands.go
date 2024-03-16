package queue

type Cell struct {
	row int
	col int
}

func shift(queue []Cell) []Cell {
	if len(queue) == 0 {
		return queue
	}
	if len(queue) == 1 {
		return []Cell{}
	}
	return queue[1:]
}

func isSafe(grid [][]byte, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == '1'
}

func BFSVertex(grid [][]byte, i, j int, visited map[Cell]bool) {
	// row and col of 8 neighbors of vertex
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}
	queue := make([]Cell, 0)
	queue = append(queue, Cell{row: i, col: j})
	visited[Cell{row: i, col: j}] = true

	for len(queue) > 0 {
		rCell := queue[0].row
		cCell := queue[0].col
		queue = shift(queue)

		// go through all 8 adjacent
		for k := 0; k < 4; k++ {
			r := rCell + row[k]
			c := cCell + col[k]
			if val, _ := visited[Cell{row: r, col: c}]; !val && isSafe(grid, r, c) {
				visited[Cell{row: r, col: c}] = true
				queue = append(queue, Cell{row: r, col: c})
			}
		}
	}
}

func CountNumsIslands(grid [][]byte) int {
	visited := make(map[Cell]bool)

	// mark all cells as not visited
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			visited[Cell{row: i, col: j}] = false
		}
	}

	// call BFS for every unvisited vertex
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if val, found := visited[Cell{row: i, col: j}]; found && !val && isSafe(grid, i, j) {
				BFSVertex(grid, i, j, visited)
				res++
			}
		}
	}

	return res
}
