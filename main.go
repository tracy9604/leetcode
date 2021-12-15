package main

import (
	"leetcode/code"
)

func main() {
	list := code.InitListNode([]int{1, 1, 2, 3, 3})
	rs := code.RemoveDupSortedList(list)
	code.PrintListNode(rs)

}
