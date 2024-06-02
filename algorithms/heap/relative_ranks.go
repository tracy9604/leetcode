package heap

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Score struct {
	val int
	idx int
}

type ScoreHeap []Score

func (h ScoreHeap) Len() int {
	return len(h)
}

func (h ScoreHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}

func (h ScoreHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ScoreHeap) Push(x interface{}) {
	*h = append(*h, x.(Score))
}

func (h *ScoreHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findRelativeRanks(score []int) []string {
	maxHeap := &ScoreHeap{}
	heap.Init(maxHeap)

	for i := 0; i < len(score); i++ {
		heap.Push(maxHeap, Score{val: score[i], idx: i})
	}

	idx := 1
	ans := make([]string, len(score))
	for maxHeap.Len() > 0 {
		top := heap.Pop(maxHeap).(Score)

		if idx == 1 {
			ans[top.idx] = "Gold Medal"
		} else if idx == 2 {
			ans[top.idx] = "Silver Medal"
		} else if idx == 3 {
			ans[top.idx] = "Bronze Medal"
		} else {
			ans[top.idx] = strconv.Itoa(idx)
		}
		idx++
	}

	return ans
}

func TestFindRelativeRanks() {
	score := []int{10, 3, 8, 9, 4}
	fmt.Println(findRelativeRanks(score))

}
