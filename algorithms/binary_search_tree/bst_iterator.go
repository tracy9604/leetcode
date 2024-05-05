package binary_search_tree

type BSTIterator struct {
	root *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		root: root,
	}
}

func (this *BSTIterator) Next() int {

}

func (this *BSTIterator) HasNext() bool {

}
