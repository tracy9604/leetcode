package linked_list

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find length
	length := 0
	curr := head

	for curr != nil {
		curr = curr.Next
		length++
	}

	// normalize k to prevent unnecessary rotations
	k %= length
	if k == 0 {
		return head
	}

	fast, slow := head, head
	// move fast pointer k steps ahead
	for k > 0 {
		fast = fast.Next
		k--
	}

	// move both fast and slow until fast reaches the end and slow reaches length - k -1
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	newHead := slow.Next
	slow.Next = nil
	fast.Next = head

	return newHead
}
