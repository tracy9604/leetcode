package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func CheckLinkedListCircle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	nodesMap := make(map[*ListNode]bool)
	curr := head
	for curr != nil {
		if _, found := nodesMap[curr]; found {
			return true
		}
		nodesMap[curr] = true
		curr = curr.Next
	}
	return false
}

func CheckLinkedListCircleV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slowPointer, fastPointer := head, head
	for slowPointer != nil && fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if slowPointer == fastPointer {
			return true
		}
	}
	return false
}
