package linked_list

func FindIntersectionPoint(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pointA := headA
	pointB := headB

	for pointA != pointB {
		pointA = pointA.Next
		pointB = pointB.Next

		if pointA == pointB {
			return pointA
		}

		if pointA == nil {
			pointA = headB
		}

		if pointB == nil {
			pointB = headA
		}
	}

	return pointA
}
