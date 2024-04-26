package linked_list

func OddEvenLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	evenHead := head.Next
	oddNode := head
	evenNode := evenHead

	for oddNode.Next != nil && evenNode.Next != nil {
		oddNode.Next = oddNode.Next.Next
		oddNode = oddNode.Next
		evenNode.Next = evenNode.Next.Next
		evenNode = evenNode.Next
	}
	oddNode.Next = evenHead
	return head
}
