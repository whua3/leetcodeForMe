package list_tree

// 230 二叉搜索树中第 k 个最小的元素
func kthSmallest(root *TreeNode, k int) int {

	i := 0
	var target int

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			traverse(node.Left)
		}
		i++
		if i == k {
			target = node.Val
			return
		}
		if node.Right != nil {
			traverse(node.Right)
		}
	}

	traverse(root)
	return target
}
