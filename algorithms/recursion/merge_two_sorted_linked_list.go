package recursion

import "sort"

func MergeTwoSortedLinkedList(list1 *ListNode, list2 *ListNode) *ListNode {
	valArr := make([]int, 0)

	for list1 != nil {
		valArr = append(valArr, list1.Val)
		list1 = list1.Next
	}

	for list2 != nil {
		valArr = append(valArr, list2.Val)
		list2 = list2.Next
	}

	sort.Ints(valArr)

	if len(valArr) == 0 {
		return nil
	}

	res := &ListNode{
		Val: valArr[0],
	}
	tmp := res
	for i := 1; i < len(valArr); i++ {
		res.Next = &ListNode{Val: valArr[i]}
		res = res.Next
	}

	return tmp
}

func MergeTwoSortedLinkedListRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	result := &ListNode{}
	if list1.Val <= list2.Val {
		result = list1
		result.Next = MergeTwoSortedLinkedListRecursion(list1.Next, list2)
	} else {
		result = list2
		result.Next = MergeTwoSortedLinkedListRecursion(list1, list2.Next)
	}
	return result
}
