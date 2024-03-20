package stack

import "fmt"

type MyQueue struct {
	input  MyStack
	output MyStack
}

type MyStack struct {
	data []int
}

func ConstructorMyStack() MyStack {
	return MyStack{data: make([]int, 0)}
}
func (this *MyStack) Push(val int) {
	this.data = append(this.data, val)
}

func (this *MyStack) Pop() {
	if len(this.data) == 0 {
		return
	}

	this.data = this.data[:len(this.data)-1]
}

func (this *MyStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{
		input:  ConstructorMyStack(),
		output: ConstructorMyStack(),
	}
}

func (this *MyQueue) Push(x int) {
	this.input.Push(x)
}

func (this *MyQueue) Pop() int {
	this.Peek()
	top := this.output.Top()
	this.output.Pop()
	return top
}

func (this *MyQueue) Peek() int {
	if len(this.output.data) == 0 {
		for len(this.input.data) > 0 {
			this.output.Push(this.input.Top())
			this.input.Pop()
		}
	}
	return this.output.Top()
}

func (this *MyQueue) Empty() bool {
	return len(this.input.data) == 0 && len(this.output.data) == 0
}

func TestMyQueue() {
	myQueue := ConstructorMyQueue()
	myQueue.Push(1)
	myQueue.Push(2)
	myQueue.Peek()
	myQueue.Pop()
	myQueue.Empty()
	fmt.Println(myQueue)
}
