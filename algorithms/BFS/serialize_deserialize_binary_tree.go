package BFS

import (
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
	data []string
}

func Constructor() Codec {
	return Codec{
		data: make([]string, 0),
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	this.serializePreOrder(root)
	return strings.Join(this.data, ",")
}

func (this *Codec) serializePreOrder(root *TreeNode) {
	if root == nil {
		this.data = append(this.data, "#")
		return
	}
	this.data = append(this.data, strconv.Itoa(root.Val))
	this.serializePreOrder(root.Left)
	this.serializePreOrder(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	this.data = strings.Split(data, ",")
	return this.deserializePreOrder()
}

func (this *Codec) deserializePreOrder() *TreeNode {
	top := this.data[0]
	this.data = this.data[1:]
	if top == "#" {
		return nil
	}
	topVal, _ := strconv.Atoi(top)
	node := &TreeNode{Val: topVal}
	node.Left = this.deserializePreOrder()
	node.Right = this.deserializePreOrder()
	return node
}

func TestCodeC() {
	obj := Constructor()
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node3 := &TreeNode{Val: 3, Left: node4, Right: node5}
	node2 := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}
	data := obj.serialize(node1)
	fmt.Println(data)
	root := obj.deserialize(data)
	fmt.Println(root)
}
