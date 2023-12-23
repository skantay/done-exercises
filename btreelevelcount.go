package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ldep int = BTreeLevelCount(root.Left)
	var rdep int = BTreeLevelCount(root.Right)

	if ldep > rdep {
		return ldep + 1
	} else {
		return rdep + 1
	}
}
