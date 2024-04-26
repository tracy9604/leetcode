package linked_list

func DetectCircle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slowPointer, fastPointer := head, head
	for slowPointer != nil && fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if slowPointer == fastPointer {
			break
		}
	}

	if slowPointer != fastPointer {
		return nil
	}

	slowPointer = head
	for slowPointer != fastPointer {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
	}

	return slowPointer
}
