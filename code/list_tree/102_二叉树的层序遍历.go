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

func levelOrder2(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	layer := []*TreeNode{root}

	for len(layer) != 0 {
		tmp := []*TreeNode{}
		l := []int{}
		for _, item := range layer {
			l = append(l, item.Val)
			if item.Left != nil {
				tmp = append(tmp, item.Left)
			}
			if item.Right != nil {
				tmp = append(tmp, item.Right)
			}
		}
		layer = tmp
		ans = append(ans, l)
	}

	return ans
}
