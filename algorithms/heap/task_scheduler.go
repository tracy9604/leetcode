package heap

import "container/heap"

type Task struct {
	count    int
	task     byte
	cooldown int
}

type TaskHeap []Task

func (h TaskHeap) Len() int {
	return len(h)
}
func (h TaskHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}
func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}
func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
	taskFreq := make(map[byte]int)
	for _, task := range tasks {
		taskFreq[task]++
	}

	h := &TaskHeap{}
	heap.Init(h)
	for task, freq := range taskFreq {
		heap.Push(h, Task{count: freq, task: task})
	}

	time := 0
	queue := make([]Task, 0)

	for h.Len() > 0 || len(queue) > 0 {
		time++

		if h.Len() > 0 {
			current := heap.Pop(h).(Task)
			current.count--
			if current.count > 0 {
				queue = append(queue, Task{count: current.count, task: current.task, cooldown: time + n})
			}
		}

		if len(queue) > 0 && queue[0].cooldown == time {
			heap.Push(h, queue[0])
			queue = queue[1:]
		}
	}
	return time
}
