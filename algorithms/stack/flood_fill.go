package stack

func isSafeCell(image [][]int, i, j int) bool {
	return i >= 0 && i < len(image) && j >= 0 && j < len(image[i])
}

func FloodFillRecursion(image [][]int, row, col, currentColor, color int, visited map[Cell]bool) {
	if image[row][col] == color {
		return
	}
	image[row][col] = color
	rRow := []int{-1, 0, 1, 0}
	cCol := []int{0, 1, 0, -1}

	for i := 0; i < 4; i++ {
		r := row + rRow[i]
		c := col + cCol[i]
		_, found := visited[Cell{row: r, col: c}]
		if !found && isSafeCell(image, r, c) && image[r][c] == currentColor && image[r][c] != color {
			visited[Cell{row: r, col: c}] = true
			FloodFillRecursion(image, r, c, currentColor, color, visited)
		}
	}
}

func popStack2(stack []Cell) (Cell, []Cell) {
	if len(stack) == 0 {
		return Cell{}, []Cell{}
	}
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return top, stack
}

func DRSFloodFill(image [][]int, row, col, color int) {
	visited := make(map[Cell]bool)
	st := make([]Cell, 0)
	st = append(st, Cell{row: row, col: col})
	rRow := []int{-1, 0, 1, 0}
	cCol := []int{0, 1, 0, -1}
	currentColor := image[row][col]

	for len(st) > 0 {
		top, tmpStack := popStack2(st)
		st = tmpStack
		image[top.row][top.col] = color

		for i := 0; i < 4; i++ {
			r := top.row + rRow[i]
			c := top.col + cCol[i]
			if _, found := visited[Cell{row: r, col: c}]; !found && isSafeCell(image, r, c) && image[r][c] == currentColor && image[r][c] != color {
				visited[Cell{row: r, col: c}] = true
				st = append(st, Cell{row: r, col: c})
			}
		}
	}
}

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	visited := make(map[Cell]bool)
	currentColor := image[sr][sc]
	FloodFillRecursion(image, sr, sc, currentColor, color, visited)
	return image
}

func FloodFill2(image [][]int, sr int, sc int, color int) [][]int {
	DRSFloodFill(image, sr, sc, color)
	return image
}
