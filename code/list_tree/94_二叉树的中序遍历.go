package list_tree

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	var inordor func(*TreeNode)
	inordor = func(node *TreeNode) {
		if node == nil {
			return
		}
		inordor(node.Left)
		ans = append(ans, node.Val)
		inordor(node.Right)
	}

	inordor(root)
	return ans
}
