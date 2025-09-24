package list_tree

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	newTarget := targetSum - root.Val

	return hasPathSum(root.Left, newTarget) || hasPathSum(root.Right, newTarget)
}
