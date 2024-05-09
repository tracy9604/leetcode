package two_pointers_technique

func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// divide
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	tmp := slow.Next
	slow.Next = nil
	leftHalf, rightHalf := SortList(head), SortList(tmp)

	// merge
	dummyNode := &ListNode{Val: -1}
	curr := dummyNode
	for leftHalf != nil && rightHalf != nil {
		if leftHalf.Val <= rightHalf.Val {
			curr = leftHalf
			leftHalf = leftHalf.Next
		} else {
			curr = rightHalf
			rightHalf = rightHalf.Next
		}
		curr = curr.Next
	}

	if leftHalf != nil {
		curr.Next = leftHalf
	}

	if rightHalf != nil {
		curr.Next = rightHalf
	}

	return dummyNode.Next
}
