package stack

import (
	"fmt"
	"math"
)

type MinStack struct {
	data []int
}

func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}

	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return -1
	}
	min := math.MaxInt
	for _, item := range this.data {
		if item < min {
			min = item
		}
	}
	return min
}

func TestMinStack() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	stack.GetMin()
	stack.Pop()
	stack.Top()
	stack.GetMin()
	fmt.Print(stack)
}
