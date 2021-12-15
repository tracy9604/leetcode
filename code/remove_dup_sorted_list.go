package code

func RemoveDupSortedList(head *ListNode) *ListNode {
	result := head

	for head != nil && head.Next != nil {

		if head.Val != head.Next.Val {
			head = head.Next
		} else {
			head.Next = head.Next.Next
		}

	}

	return result
}
