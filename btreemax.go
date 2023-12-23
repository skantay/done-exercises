package piscine

var max string

var trees []*TreeNode

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{}
	}

	trees = append(trees, root)
	BTreeMax(root.Left)
	BTreeMax(root.Right)

	return trees[len(trees)-2]
}
