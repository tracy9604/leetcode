package two_pointers_technique

func PartitionList(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}

	var leftHead, rightHead *ListNode
	var currLeft, currRight *ListNode
	curr := head
	for curr != nil {
		if curr.Val < x {
			if leftHead == nil {
				leftHead = curr
				currLeft = curr
			} else {
				currLeft.Next = curr
				currLeft = currLeft.Next
			}

		} else {
			if rightHead == nil {
				rightHead = curr
				currRight = curr
			} else {
				currRight.Next = curr
				currRight = currRight.Next
			}
		}
		curr = curr.Next
	}

	if currLeft != nil {
		currLeft.Next = rightHead
	}
	if currRight != nil {
		currRight.Next = nil
	}
	if leftHead == nil {
		return rightHead
	}
	return leftHead
}
