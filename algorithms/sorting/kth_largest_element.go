package sorting

import "math"

type MaxHeap struct {
	heap []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		heap: make([]int, 0),
	}
}

func (m *MaxHeap) swap(index1, index2 int) {
	m.heap[index1], m.heap[index2] = m.heap[index2], m.heap[index1]
}

func (m *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := int(math.Floor(float64((index - 1) / 2)))
		if m.heap[parentIndex] < m.heap[index] {
			break
		}
		m.swap(parentIndex, index)
		index = parentIndex
	}
}

func (m *MaxHeap) heapifyDown(index int) {
	left := 2*index + 1
	right := 2*index + 2
	largestIndex := index

	if left < len(m.heap) && m.heap[left] > m.heap[largestIndex] {
		largestIndex = left
	}
	if right < len(m.heap) && m.heap[right] > m.heap[largestIndex] {
		largestIndex = right
	}
	if largestIndex != index {
		m.swap(largestIndex, index)
		m.heapifyDown(largestIndex)
	}
}

func (m *MaxHeap) push(val int) {
	m.heap = append(m.heap, val)
	m.heapifyUp(len(m.heap) - 1)
}

func (m *MaxHeap) pop() int {
	if len(m.heap) == 0 {
		return -1
	}
	if len(m.heap) == 1 {
		top := m.heap[0]
		m.heap = make([]int, 0)
		return top
	}
	minIdx := 0
	for i := 1; i < len(m.heap); i++ {
		if m.heap[minIdx] > m.heap[i] {
			minIdx = i
		}
	}
	root := m.heap[0]
	m.heap = append(m.heap[:minIdx], m.heap[minIdx+1:]...)
	m.heapifyDown(0)
	return root
}

func (m *MaxHeap) top() int {
	if len(m.heap) == 0 {
		return -1
	}
	return m.heap[0]
}

func (m *MaxHeap) size() int {
	return len(m.heap)
}

func FindKthLargestElement(nums []int, k int) int {
	pq := NewMaxHeap()

	for i := 0; i < len(nums); i++ {
		pq.push(nums[i])

		if pq.size() > k {
			pq.pop()
		}
	}
	return pq.heap[k-1]
}
