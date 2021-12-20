package code

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeFromArrayWithRoot(arr []int) *TreeNode {
	root := &TreeNode{}

	i := 0

	for i < len(arr)-1 {
		root.Val = arr[i]
		count := 0

		if i < len(arr)-2 {
			root.Left = &TreeNode{
				Val:   arr[i+1],
				Left:  nil,
				Right: nil,
			}
			count++
		}

		if i < len(arr)-3 {
			root.Right = &TreeNode{
				Val:   arr[i+2],
				Left:  nil,
				Right: nil,
			}
			count++
		}
		i += count
	}

	return root
}
