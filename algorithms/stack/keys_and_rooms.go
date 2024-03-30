package stack

func DFSVisitAllRoom(rooms [][]int, idx int, visited map[int]bool) {
	visited[idx] = true
	for _, item := range rooms[idx] {
		if _, found := visited[item]; !found {
			DFSVisitAllRoom(rooms, item, visited)
		}
	}
}

func CanVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)
	DFSVisitAllRoom(rooms, 0, visited)
	return len(visited) == len(rooms)
}
