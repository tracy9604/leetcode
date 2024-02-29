package queue

func InitTree() map[int][]int {
	data := map[int][]int{
		1: {2, 3, 4},
		2: {5},
		3: {5, 6},
		4: {7},
		6: {7},
	}

	return data
}

func BFS() int {
	queue := NewQueue(7)
	tree := InitTree()
	queue.Enqueue(1)
	step := 0
	target := 7
	for !queue.IsEmpty() {
		curr := queue.data[0]
		if curr == target {
			return step
		}
		if val, found := tree[curr]; found && len(val) > 0 {
			for _, item := range val {
				queue.Enqueue(item)
			}
		}
		queue.Dequeue()
		step++
	}

	if step > 0 {
		return step
	}
	return -1
}
