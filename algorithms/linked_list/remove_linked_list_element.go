package linked_list

func RemoveLinkedListElement(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	var prev *ListNode
	for curr != nil {
		if curr.Val != val {
			prev = curr
			curr = curr.Next
			continue
		}

		// head's value is val
		if prev == nil {
			head = curr.Next
			curr = curr.Next
			continue
		}

		// remove curr
		prev.Next = curr.Next
		curr = curr.Next
	}
	return head
}
