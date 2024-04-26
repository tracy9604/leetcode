package linked_list

func lengthList(head *ListNode) int {
	if head == nil {
		return 0
	}
	curr := head
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}

func RemoveNthNode(head *ListNode, n int) *ListNode {
	length := lengthList(head)
	idxNode := length - n + 1
	var prev *ListNode
	curr := head
	for i := 1; i < idxNode; i++ {
		prev = curr
		curr = curr.Next
	}
	if prev == nil {
		head = head.Next
	} else if prev.Next != nil {
		prev.Next = prev.Next.Next
	}
	return head
}

func RemoveNthNodeByTwoPointer(head *ListNode, n int) *ListNode {
	first := head
	second := head

	for i := 0; i < n; i++ {
		if second.Next == nil {
			if i == n-1 {
				head = head.Next
			}
			return head
		}
		second = second.Next
	}

	for second.Next != nil {
		first = first.Next
		second = second.Next
	}

	first.Next = first.Next.Next
	return head
}
