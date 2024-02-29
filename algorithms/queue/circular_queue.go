package queue

import "fmt"

type MyCircularQueue struct {
	queue []int
	head  int
	tail  int
	size  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{queue: make([]int, k), head: -1, tail: -1, size: k}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	if q.IsEmpty() {
		q.head = 0
	}
	q.tail = (q.tail + 1) % q.size
	q.queue[q.tail] = value
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	if q.head == q.tail {
		q.head = -1
		q.tail = -1
		return true
	}
	q.head = (q.head + 1) % q.size
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[q.head]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.queue[q.tail]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.head == -1
}

func (q *MyCircularQueue) IsFull() bool {
	return (q.tail+1)%q.size == q.head
}

func TestCircularQueue() {
	myCircularQueue := Constructor(2)
	myCircularQueue.EnQueue(1)
	myCircularQueue.EnQueue(2)
	myCircularQueue.DeQueue()
	myCircularQueue.EnQueue(3)
	myCircularQueue.DeQueue()
	myCircularQueue.EnQueue(3)
	myCircularQueue.DeQueue()
	myCircularQueue.EnQueue(3)
	myCircularQueue.DeQueue()
	myCircularQueue.Front()
	fmt.Println(myCircularQueue.queue)
}
