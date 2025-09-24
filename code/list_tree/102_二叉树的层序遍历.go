package list_tree

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	l := []*TreeNode{root}

	var level func(layer []*TreeNode)
	level = func(layer []*TreeNode) {
		if len(layer) == 0 {
			return
		}
		tmp := []int{}
		newLayer := []*TreeNode{}
		for _, node := range layer {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				newLayer = append(newLayer, node.Left)
			}
			if node.Right != nil {
				newLayer = append(newLayer, node.Right)
			}
		}
		ans = append(ans, tmp)
		level(newLayer)
	}
	level(l)
	return ans
}
