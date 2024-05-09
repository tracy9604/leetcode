package two_pointers_technique

func FindIntersectionTwoLinkedLists(headA, headB *ListNode) *ListNode {
	if headA == nil && headB == nil {
		return nil
	}

	currA, currB := headA, headB
	for currA != currB {
		if currA != nil {
			currA = currA.Next
		} else {
			currA = headB
		}

		if currB != nil {
			currB = currB.Next
		} else {
			currB = headA
		}
	}

	return currA
}
