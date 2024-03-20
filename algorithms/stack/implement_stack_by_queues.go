package stack

import "fmt"

type MyQueue1 struct {
	data []int
}

func ConstructorMyQueue1() MyQueue1 {
	return MyQueue1{
		data: make([]int, 0),
	}
}

func (q *MyQueue1) Push(x int) {
	q.data = append(q.data, x)
}

func (q *MyQueue1) Pop() int {
	if q.IsEmpty() {
		return -1
	}
	top := q.data[0]
	if len(q.data) == 1 {
		q.data = make([]int, 0)
		return top
	}
	q.data = q.data[1:]
	return top
}
func (q *MyQueue1) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *MyQueue1) Top() int {
	if q.IsEmpty() {
		return -1
	}
	return q.data[0]
}

type MyStack1 struct {
	input  MyQueue1
	output MyQueue1
}

func ConstructorMyStack1() MyStack1 {
	return MyStack1{
		input:  ConstructorMyQueue1(),
		output: ConstructorMyQueue1(),
	}
}

func (this *MyStack1) Push(x int) {
	this.input.Push(x)
}

func (this *MyStack1) Pop() int {
	for !this.input.IsEmpty() {
		this.output.Push(this.input.Pop())
	}
	top := this.output.data[len(this.output.data)-1]
	this.output.data = this.output.data[:len(this.output.data)-1]
	return top
}

func (this *MyStack1) Top() int {
	for !this.input.IsEmpty() {
		this.output.Push(this.input.Pop())
	}
	return this.output.data[len(this.output.data)-1]
}

func (this *MyStack1) Empty() bool {
	return this.input.IsEmpty() && this.output.IsEmpty()
}

func TestMyStack1() {
	myStack1 := ConstructorMyStack1()
	myStack1.Push(1)
	myStack1.Push(2)
	fmt.Println(myStack1.Top())
	myStack1.Push(3)
	fmt.Println(myStack1.Top())
}
