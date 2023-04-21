package array

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func InitListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	rs := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	head := rs

	for i := 1; i < len(arr); i++ {
		node := &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		rs.Next = node
		rs = rs.Next

	}
	return head
}

func PrintListNode(list *ListNode) {
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func MergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	tmp := &ListNode{}
	result := tmp

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = list1
			tmp = tmp.Next
			list1 = list1.Next
		} else {
			tmp.Next = list2
			tmp = tmp.Next
			list2 = list2.Next
		}
	}

	for list1 != nil {
		tmp.Next = list1
		tmp = tmp.Next
		list1 = list1.Next
	}

	for list2 != nil {
		tmp.Next = list2
		tmp = tmp.Next
		list2 = list2.Next
	}
	return result.Next
}
