package heap

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func countNodes(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + countNodes(node.left) + countNodes(node.right)
}

func isComplete(node *Node, index, numNodes int) bool {
	if node == nil {
		return true
	}

	if index >= numNodes {
		return false
	}

	return isComplete(node.left, 2*index+1, numNodes) && isComplete(node.right, 2*index+2, numNodes)
}

func isHeap(node *Node) bool {
	if node.left == nil && node.right == nil {
		return true
	}

	if node.right == nil {
		return node.val >= node.left.val
	}

	if node.val >= node.left.val && node.val >= node.right.val {
		return isHeap(node.left) && isHeap(node.right)
	}

	return false
}

func checkBinaryTreeIsHeap(root *Node) bool {
	numNodes := countNodes(root)
	index := 0

	return isComplete(root, index, numNodes) && isHeap(root)
}

func TestBinaryTreeIsHeap() {
	node2 := &Node{val: 2}
	node4 := &Node{val: 4}
	node12 := &Node{val: 12}
	node3 := &Node{val: 3, left: node2, right: node4}
	node46 := &Node{val: 46, left: node12, right: node3}
	node7 := &Node{val: 7}
	node31 := &Node{val: 31}
	node37 := &Node{val: 37, left: node7, right: node31}
	node97 := &Node{val: 97, left: node46, right: node37}

	fmt.Print(checkBinaryTreeIsHeap(node97))
}
