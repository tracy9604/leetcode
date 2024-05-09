package two_pointers_technique

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}

	// find the middle of the list
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	// reverse the second half
	secondHalf := slow.Next
	slow.Next = nil
	curr := secondHalf
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	secondHalf = prev

	// merge two halves
	curr1 := head
	curr2 := secondHalf
	for curr1 != nil && curr2 != nil {
		next1 := curr1.Next
		next2 := curr2.Next
		curr1.Next = curr2
		curr1 = next1
		curr2.Next = next1
		curr2 = next2
	}
}
