package list_tree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ans int

	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := getDepth(node.Left), getDepth(node.Right)
		ans = max(ans, l+r+1)

		return max(l, r) + 1
	}

	_ = getDepth(root)

	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
