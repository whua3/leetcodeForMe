package list_tree

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}

func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val && check(a.Left, b.Right) && check(a.Right, b.Left)
}
