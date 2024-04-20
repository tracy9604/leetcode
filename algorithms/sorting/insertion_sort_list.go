package sorting

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyNode := &ListNode{Val: 0}
	prev := dummyNode

	for head != nil {
		next := head.Next
		if prev.Val >= head.Val {
			prev = dummyNode
		}
		for prev.Next != nil && prev.Next.Val < head.Val {
			prev = prev.Next
		}
		head.Next = prev.Next
		prev.Next = head
		head = next
	}
	return dummyNode.Next
}
