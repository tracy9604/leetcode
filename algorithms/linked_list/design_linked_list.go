package linked_list

type SinglyLinkedList struct {
	val  int
	next *SinglyLinkedList
}

type MyLinkedList struct {
	head   *SinglyLinkedList
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{length: 0}
}

func (this *MyLinkedList) Get(index int) int {
	curr := this.head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.next
	}
	if curr != nil {
		return curr.val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	curr := &SinglyLinkedList{val: val}
	curr.next = this.head
	this.head = curr
}

func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.AddAtHead(val)
		return
	}
	node := &SinglyLinkedList{val: val}
	curr := this.head
	for curr != nil && curr.next != nil {
		curr = curr.next
	}
	curr.next = node
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	node := &SinglyLinkedList{val: val}
	curr := this.head
	for i := 0; i < index-1 && curr != nil; i++ {
		curr = curr.next
	}
	if curr != nil {
		node.next = curr.next
		curr.next = node
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.head = this.head.next
		return
	}
	curr := this.head
	for i := 0; i < index-1 && curr != nil; i++ {
		curr = curr.next
	}

	if curr != nil && curr.next != nil {
		curr.next = curr.next.next
	}
}

func TestDesignLinkedList() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2)
	obj.Get(1)
	obj.DeleteAtIndex(1)
	obj.Get(1)
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
