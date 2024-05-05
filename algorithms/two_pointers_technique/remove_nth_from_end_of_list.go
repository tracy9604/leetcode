package two_pointers_technique

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthNodeFromEndOfList(head *ListNode, n int) *ListNode {
	first, second := head, head

	for i := 0; i < n; i++ {
		if second.Next == nil && i == n-1 {
			// n is the end of list
			head = head.Next
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
