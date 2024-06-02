package heap

import (
	"container/heap"
	"fmt"
)

type Pair struct {
	num1 int
	num2 int
	sum  int
}

type PairHeap []Pair

func (h PairHeap) Len() int {
	return len(h)
}

func (h PairHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func (h PairHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	minHeap := &PairHeap{}
	heap.Init(minHeap)
	ans := make([][]int, 0)

	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return ans
	}

	for j := 0; j < len(nums2); j++ {
		heap.Push(minHeap, Pair{num1: 0, num2: j, sum: nums1[0] + nums2[j]})
	}

	for k > 0 && minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(Pair)
		ans = append(ans, []int{nums1[top.num1], nums2[top.num2]})
		if top.num1 < len(nums1)-1 {
			heap.Push(minHeap, Pair{num1: top.num1 + 1, num2: top.num2, sum: nums1[top.num1+1] + nums2[top.num2]})
		}
		k--
	}

	return ans
}

func TestKSmallestPairs() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{1, 1, 3}
	k := 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
