package list_tree

import "math"

func isValidBST(root *TreeNode) bool {
	return helper2(root, math.MinInt64, math.MaxInt64)
}

func helper2(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper2(root.Left, lower, root.Val) && helper2(root.Right, root.Val, upper)
}
