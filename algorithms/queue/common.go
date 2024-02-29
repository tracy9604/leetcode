package queue

type Queue struct {
	data    []int
	pointer int
}

func NewQueue(k int) Queue {
	return Queue{
		data:    make([]int, k),
		pointer: -1,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.pointer == -1
}

func (q *Queue) Enqueue(value int) {
	if q.pointer >= len(q.data) {
		return
	}
	if q.pointer == -1 {
		q.pointer = 0
	}
	q.data[q.pointer] = value
	q.pointer++
}

func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		return
	}
	for i := 0; i < len(q.data)-1; i++ {
		q.data[i] = q.data[i+1]
	}
	q.pointer--
}

type CircularQueue struct {
	data []int
	head int
	tail int
	size int
}

func NewCircularQueue(k int) CircularQueue {
	return CircularQueue{
		data: make([]int, k),
		head: -1,
		tail: -1,
		size: k,
	}
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == -1
}

func (q *CircularQueue) IsFull() bool {
	return q.head == (q.tail+1)%q.size
}

func (q *CircularQueue) Enqueue(value int) {
	if q.IsFull() {
		return
	}
	if q.IsEmpty() {
		q.head = value
		return
	}
	q.tail = (q.tail + 1) % q.size
	q.data[q.tail] = value
}

func (q *CircularQueue) Dequeue() {
	if q.IsEmpty() {
		return
	}
	if q.head == q.tail {
		q.head = -1
		q.tail = -1
		return
	}
	q.head = (q.head + 1) % q.size
}

type Node struct {
	data      int
	neighbors []*Node
}

type List struct {
	head *Node
}

//
//func (l *List) add(value int) {
//	newNode := &Node{data: value}
//	if l.head == nil {
//		l.head = newNode
//		return
//	}
//	curr := l.head
//	for curr.next != nil {
//		curr = curr.next
//	}
//	curr.next = newNode
//}
//
//func (l *List) remove(value int) {
//	if l.head == nil {
//		return
//	}
//
//	if l.head.data == value {
//		l.head = nil
//		return
//	}
//	curr := l.head
//	for curr.next != nil && curr.next.data != value {
//		curr = curr.next
//	}
//	if curr.next != nil {
//		curr.next = curr.next.next
//	}
//}
