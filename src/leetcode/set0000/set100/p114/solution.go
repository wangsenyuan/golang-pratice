package p114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	if root.Left != nil {
		tmp := root.Left
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
}
