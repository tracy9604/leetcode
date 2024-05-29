package heap

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func nthUglyNumber(n int) int {
	minHeap := &IntHeap{1}
	heap.Init(minHeap)

	visited := map[int]bool{1: true}

	ugly := 0
	factors := []int{2, 3, 5}

	for i := 0; i < n; i++ {
		ugly = heap.Pop(minHeap).(int)

		for _, factor := range factors {
			next := ugly * factor
			if _, found := visited[next]; !found {
				visited[next] = true
				heap.Push(minHeap, next)
			}
		}
	}

	return ugly
}

func TestUglyNumber() {
	n := 10
	fmt.Println(nthUglyNumber(n))
}
