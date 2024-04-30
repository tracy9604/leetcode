package linked_list

func pop(st []int) (int, []int) {
	if len(st) == 0 {
		return -1, []int{}
	}
	if len(st) == 1 {
		return st[0], []int{}
	}
	return st[0], st[1:]
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1 == nil && list2 == nil {
		return nil
	}

	head1 := list1
	head2 := list2

	st := make([]int, 0)

	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			st = append(st, head1.Val)
			head1 = head1.Next
		} else {
			st = append(st, head2.Val)
			head2 = head2.Next
		}
	}

	for head1 != nil {
		st = append(st, head1.Val)
		head1 = head1.Next
	}

	for head2 != nil {
		st = append(st, head2.Val)
		head2 = head2.Next
	}

	var head, curr *ListNode
	for len(st) > 0 {
		top, tmp := pop(st)
		st = tmp
		node := &ListNode{Val: top}
		if head == nil {
			head = node
			curr = head
			continue
		}
		curr.Next = node
		curr = curr.Next
	}

	return head
}
