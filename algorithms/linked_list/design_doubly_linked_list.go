package linked_list

type DoublyLinkedList struct {
	Val  int
	Next *DoublyLinkedList
	Prev *DoublyLinkedList
}

type MyDoublyLinkedList struct {
	head *DoublyLinkedList
}

func ConstructorMyDoublyLinkedList() MyDoublyLinkedList {
	return MyDoublyLinkedList{}
}

func (this *MyDoublyLinkedList) Get(index int) int {
	curr := this.head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	if curr != nil {
		return curr.Val
	}
	return -1
}

func (this *MyDoublyLinkedList) AddAtHead(val int) {
	curr := &DoublyLinkedList{Val: val}
	curr.Next = this.head
	if this.head != nil {
		this.head.Prev = curr
	}
	this.head = curr
}

func (this *MyDoublyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}
	node := &DoublyLinkedList{Val: val}
	curr := this.head
	for curr != nil && curr.Next != nil {
		curr = curr.Next
	}
	node.Prev = curr
	curr.Next = node
}

func (this *MyDoublyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	node := &DoublyLinkedList{Val: val}
	curr := this.head
	for i := 0; i < index-1 && curr != nil; i++ {
		curr = curr.Next
	}
	if curr != nil {
		node.Next = curr.Next
		node.Prev = curr

		if node.Next != nil {
			node.Next.Prev = node
		}
		curr.Next = node
	}
}

func (this *MyDoublyLinkedList) DeleteAtIndex(index int) {
	if this.head == nil {
		return
	}
	if index == 0 && this.head != nil {
		this.head = this.head.Next

		if this.head != nil {
			this.head.Prev = nil
		}
		return
	}

	curr := this.head
	for i := 0; i < index-1 && curr != nil; i++ {
		curr = curr.Next
	}

	if curr != nil && curr.Next != nil {
		curr.Next = curr.Next.Next
		if curr.Next != nil {
			curr.Next.Prev = curr
		}
	}
}

func TestDoublyLinkedList() {
	obj := ConstructorMyDoublyLinkedList()
	obj.AddAtHead(7)
	obj.AddAtHead(2)
	obj.AddAtHead(1)
	obj.AddAtIndex(3, 0)
	obj.DeleteAtIndex(2)
	obj.AddAtHead(6)
	obj.AddAtTail(4)
	obj.Get(4)
	obj.AddAtHead(4)
	obj.AddAtIndex(5, 0)
	obj.AddAtHead(6)
}
