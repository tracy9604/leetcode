package BFS

import "math"

type water struct {
	jugA int
	jugB int
}

func MeasureWater(x int, y int, target int) bool {
	if target > x+y {
		return false
	}
	if target == 0 || x == target || y == target || target == x+y {
		return true
	}
	queue := make([][]int, 0)
	visited := make(map[*water]bool)

	queue = append(queue, []int{0, 0})

	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]

		if top[0] == target || top[1] == target || top[0]+top[1] == target {
			return true
		}

		// fill jug1
		key := &water{jugA: x, jugB: top[1]}
		if _, found := visited[key]; !found {
			visited[key] = true
			queue = append(queue, []int{x, top[1]})
		}
		// fill jug2
		key = &water{jugA: top[0], jugB: y}
		if _, found := visited[key]; !found {
			visited[key] = true
			queue = append(queue, []int{top[0], y})
		}
		// empty jug1
		key = &water{jugA: 0, jugB: top[1]}
		if _, found := visited[key]; !found {
			visited[key] = true
			queue = append(queue, []int{0, top[1]})
		}
		// empty jug2
		key = &water{jugA: top[0], jugB: 0}
		if _, found := visited[key]; !found {
			visited[key] = true
			queue = append(queue, []int{top[0], 0})
		}
		// pour jug1 to jug2
		pourToY := int(math.Min(float64(top[0]), float64(y-top[1])))
		key = &water{jugA: top[0] - pourToY, jugB: top[1] + pourToY}
		if _, found := visited[key]; !found {
			visited[key] = true
			queue = append(queue, []int{top[0] - pourToY, top[1] + pourToY})
		}
		// pour jug2 to jug1
		pourToX := int(math.Min(float64(x-top[0]), float64(top[1])))
		key = &water{jugA: top[0] + pourToX, jugB: top[1] - pourToX}
		if _, found := visited[key]; !found {
			visited[key] = true
			queue = append(queue, []int{top[0] + pourToX, top[1] - pourToX})
		}

	}
	return false
}
