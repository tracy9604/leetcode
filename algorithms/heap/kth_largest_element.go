package heap

import "fmt"

type MinHeap struct {
	data []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{data: make([]int, 0)}
}

func (h *MinHeap) Push(val int) {
	h.data = append(h.data, val)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeap) Pop() int {
	if len(h.data) == 0 {
		return -1
	}

	top := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.heapifyDown(0)
	return top
}

func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}

func (h *MinHeap) heapifyUp(index int) {
	for h.data[parent(index)] > h.data[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.data) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.data[l] < h.data[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.data[index] <= h.data[childToCompare] {
			return
		}

		h.swap(index, childToCompare)
		index = childToCompare
		l, r = left(index), right(index)
	}
}

func findKthLargestElement(nums []int, k int) int {
	minHeap := NewMinHeap()
	for _, num := range nums {
		minHeap.Push(num)
		if len(minHeap.data) > k {
			minHeap.Pop()
		}
	}
	return minHeap.Pop()
}

func TestMinHeap() {
	nums := []int{7, 6, 5, 4, 3, 2, 1}
	k := 5

	fmt.Println(findKthLargestElement(nums, k))
}
