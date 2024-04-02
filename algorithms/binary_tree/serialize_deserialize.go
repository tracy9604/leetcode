package binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) popQueue(queue []*TreeNode) (*TreeNode, []*TreeNode) {
	if len(queue) == 0 {
		return nil, queue
	}
	if len(queue) == 1 {
		return queue[0], []*TreeNode{}
	}
	front := queue[0]
	queue = queue[1:]
	return front, queue
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	queue := make([]*TreeNode, 0)
	res := make([]string, 0)

	queue = append(queue, root)

	// preorder

	for len(queue) > 0 {
		front, tmpQueue := this.popQueue(queue)
		queue = tmpQueue

		if front == nil {
			res = append(res, "#")
			continue
		}

		res = append(res, strconv.Itoa(front.Val))
		queue = append(queue, front.Left)
		queue = append(queue, front.Right)
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	dataArr := strings.Split(data, ",")
	val, _ := strconv.Atoi(dataArr[0])
	root := &TreeNode{Val: val}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for i := 1; i < len(dataArr); i += 2 {
		node, tmpQueue := this.popQueue(queue)
		queue = tmpQueue

		if dataArr[i] != "#" {
			val, _ := strconv.Atoi(dataArr[i])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		if dataArr[i+1] != "#" {
			val, _ := strconv.Atoi(dataArr[i+1])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
	}
	return root
}

func TestCodec() {
	codeC := Constructor()

	node6 := &TreeNode{Val: 6}
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4, Left: node6, Right: node7}
	node5 := &TreeNode{Val: 5}
	node3 := &TreeNode{
		Val:   3,
		Left:  node4,
		Right: node5,
	}
	node2 := &TreeNode{Val: 2}
	node1 := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}
	data := codeC.serialize(node1)
	root := codeC.deserialize(data)
	fmt.Println(root)
}
