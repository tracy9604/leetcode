package stack

type Cell struct {
	row int
	col int
}

func isSafe(grid [][]byte, i, j int) bool {
	return i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == '1'
}

func DFSVertex(grid [][]byte, i, j int, visited map[Cell]bool) {
	if grid[i][j] == '0' {
		return
	}
	row := []int{-1, 0, 1, 0}
	col := []int{0, 1, 0, -1}
	visited[Cell{i, j}] = true

	for k := 0; k < 4; k++ {
		rRow := row[k] + i
		cCol := col[k] + j
		cCell := Cell{
			row: rRow,
			col: cCol,
		}
		if val, found := visited[cCell]; found && !val && isSafe(grid, rRow, cCol) {
			DFSVertex(grid, rRow, cCol, visited)
		}
	}
}

func FindNumberOfIslands(grid [][]byte) int {
	visited := make(map[Cell]bool)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			visited[Cell{row: i, col: j}] = false
		}
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			cCell := Cell{row: i, col: j}
			if val, found := visited[cCell]; found && !val && isSafe(grid, i, j) {
				DFSVertex(grid, i, j, visited)
				res++
			}
		}
	}
	return res
}
