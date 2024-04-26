package linked_list

func CheckPalindromeLinkedList(head *ListNode) bool {
	if head == nil {
		return false
	}

	st := make([]*ListNode, 0)
	curr := head
	for curr != nil {
		st = append(st, curr)
		curr = curr.Next
	}

	curr = head
	for curr != nil && len(st) > 0 {
		top := st[len(st)-1]
		st = st[:(len(st) - 1)]
		if curr.Val != top.Val {
			return false
		}
		curr = curr.Next
	}

	return true
}

func CheckPalindromeLinkedListV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slowPointer, fastPointer := head, head
	prevSlowPointer := head
	var midNode, secondHalf *ListNode
	for fastPointer != nil && fastPointer.Next != nil {
		fastPointer = fastPointer.Next.Next
		prevSlowPointer = slowPointer
		slowPointer = slowPointer.Next
	}

	if fastPointer != nil {
		midNode = slowPointer
		slowPointer = slowPointer.Next
	}

	secondHalf = slowPointer
	prevSlowPointer.Next = nil

	secondHalf = reverse(secondHalf)

	res := compareLists(head, secondHalf)
	secondHalf = reverse(secondHalf)

	if midNode != nil {
		prevSlowPointer.Next = midNode
		midNode.Next = secondHalf
	} else {
		prevSlowPointer.Next = secondHalf
	}
	return res
}

func compareLists(head1, head2 *ListNode) bool {
	temp1, temp2 := head1, head2

	for temp1 != nil && temp2 != nil {
		if temp1.Val != temp2.Val {
			return false
		}

		temp1 = temp1.Next
		temp2 = temp2.Next
	}

	if temp1 == nil && temp2 == nil {
		return true
	}
	return false
}

func reverse(head *ListNode) *ListNode {
	var prev, next *ListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
