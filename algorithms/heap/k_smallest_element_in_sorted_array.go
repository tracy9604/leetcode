package heap

import (
	"container/heap"
	"fmt"
)

type Cell struct {
	row, col, val int
}

type CellHeap []Cell

func (h CellHeap) Len() int {
	return len(h)
}

func (h CellHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h CellHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *CellHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *CellHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	minHeap := &CellHeap{}
	heap.Init(minHeap)
	n := len(matrix)

	for j := 0; j < n; j++ {
		heap.Push(minHeap, Cell{row: 0, col: j, val: matrix[0][j]})
	}

	ans := -1
	for k > 0 && minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(Cell)
		ans = top.val

		if top.row < n-1 {
			heap.Push(minHeap, Cell{row: top.row + 1, col: top.col, val: matrix[top.row+1][top.col]})
		}
		k--
	}

	return ans
}

func TestKthSmallest() {
	matrix := [][]int{
		{-5},
	}
	k := 1
	fmt.Println(kthSmallest(matrix, k))
}
