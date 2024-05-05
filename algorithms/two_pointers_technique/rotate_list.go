package two_pointers_technique

func RotateList(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	k %= length

	if k == 0 {
		return head
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil

	return head
}
