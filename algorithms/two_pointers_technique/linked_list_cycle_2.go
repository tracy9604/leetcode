package two_pointers_technique

func FindLinkedListCycle2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			ans := head

			for ans != slow {
				ans = ans.Next
				slow = slow.Next
			}

			return ans
		}
	}

	return nil
}
