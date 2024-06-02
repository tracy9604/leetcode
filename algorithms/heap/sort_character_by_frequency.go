package heap

import (
	"container/heap"
	"fmt"
)

type FrequencyElement struct {
	char      byte
	frequency int
}

type FrequencyHeap []FrequencyElement

func (h FrequencyHeap) Len() int {
	return len(h)
}

func (h FrequencyHeap) Less(i, j int) bool {
	return h[i].frequency > h[j].frequency
}
func (h FrequencyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *FrequencyHeap) Push(x interface{}) {
	*h = append(*h, x.(FrequencyElement))
}
func (h *FrequencyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func frequencySort(s string) string {
	frequencyMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		frequencyMap[s[i]]++
	}

	maxHeap := &FrequencyHeap{}
	heap.Init(maxHeap)

	ans := ""
	for key, val := range frequencyMap {
		heap.Push(maxHeap, FrequencyElement{char: key, frequency: val})
	}

	for maxHeap.Len() > 0 {
		top := heap.Pop(maxHeap).(FrequencyElement)
		for top.frequency > 0 {
			ans += string(top.char)
			top.frequency--
		}
	}

	return ans
}

func TestFrequencySort() {
	s := "treetthh"
	fmt.Println(frequencySort(s))
}
