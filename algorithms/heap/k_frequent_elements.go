package heap

import (
	"container/heap"
	"fmt"
)

type frequentTuple struct {
	val      int
	frequent int
}

// Element holds the value and its frequency.
type Element struct {
	value     int
	frequency int
}

type FrequentHeap []frequentTuple

func (f FrequentHeap) Len() int {
	return len(f)
}

func (f FrequentHeap) Less(i, j int) bool {
	return f[i].frequent < f[j].frequent
}

func (f FrequentHeap) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f *FrequentHeap) Push(x interface{}) {
	*f = append(*f, x.(frequentTuple))
}

func (f *FrequentHeap) Pop() interface{} {
	old := *f
	n := len(old)
	x := old[n-1]
	*f = old[0 : n-1]
	return x
}

func topKFrequentElements(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	minHeap := &FrequentHeap{}
	heap.Init(minHeap)
	for key, val := range frequencyMap {
		heap.Push(minHeap, frequentTuple{val: key, frequent: val})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	ans := make([]int, 0)
	for minHeap.Len() > 0 {
		top := minHeap.Pop().(frequentTuple)
		ans = append(ans, top.val)
	}

	return ans
}

// MinHeap is a min-heap of Elements.
type MinFrequentHeap []Element

// Implement heap.Interface.
func (h MinFrequentHeap) Len() int           { return len(h) }
func (h MinFrequentHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinFrequentHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds an element to the heap.
func (h *MinFrequentHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

// Pop removes the smallest element from the heap.
func (h *MinFrequentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// topKFrequent finds the k most frequent elements in nums.
func topKFrequent(nums []int, k int) []int {
	// Count the frequency of each element.
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Use a min-heap to keep the top k elements.
	h := &MinFrequentHeap{}
	heap.Init(h)

	for value, frequency := range frequencyMap {
		heap.Push(h, Element{value, frequency})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Extract the elements from the heap.
	result := make([]int, 0, k)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(Element).value)
	}

	return result
}

func TestKFrequentElements() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequentElements(nums, k))
	//fmt.Println(topKFrequent(nums, k))
}
