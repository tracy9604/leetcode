package heap

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListHeap []ListNode

func (h ListHeap) Len() int {
	return len(h)
}

func (h ListHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h ListHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *ListHeap) Push(x interface{}) {
	*h = append(*h, x.(ListNode))
}
func (h *ListHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	minHeap := &ListHeap{}
	heap.Init(minHeap)

	for i := 0; i < len(lists); i++ {
		node := lists[i]
		for node != nil {
			heap.Push(minHeap, *node)
			node = node.Next
		}
	}

	var root, prev *ListNode
	for minHeap.Len() > 0 {
		top := heap.Pop(minHeap).(ListNode)
		if root == nil {
			root = &top
			prev = root
			continue
		}

		prev.Next = &top
		prev = prev.Next
	}

	prev.Next = nil

	return root
}

func TestMergeKLists() {
	node15 := &ListNode{Val: 5}
	node14 := &ListNode{Val: 4, Next: node15}
	list1 := &ListNode{Val: 1, Next: node14}

	node24 := &ListNode{Val: 4}
	node23 := &ListNode{Val: 3, Next: node24}
	list2 := &ListNode{Val: 1, Next: node23}

	node36 := &ListNode{Val: 6}
	list3 := &ListNode{Val: 2, Next: node36}

	fmt.Println(mergeKLists([]*ListNode{list1, list2, list3}))
}
