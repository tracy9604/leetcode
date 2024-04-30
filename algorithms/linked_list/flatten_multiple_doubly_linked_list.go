package linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func FlattenMultipleDoublyLinkedList(root *Node) *Node {
	if root == nil {
		return root
	}

	curr := root

	for curr != nil {
		if curr.Child != nil {
			cachedNext := curr.Next
			curr.Next = curr.Child
			curr.Child.Prev = curr
			curr.Child = nil

			tail := curr.Next
			for tail != nil && tail.Next != nil {
				tail = tail.Next
			}
			tail.Next = cachedNext
			if cachedNext != nil {
				cachedNext.Prev = tail
			}
		}
		curr = curr.Next
	}

	return root
}
