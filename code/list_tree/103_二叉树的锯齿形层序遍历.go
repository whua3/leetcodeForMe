package list_tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)

	if root == nil {
		return ans
	}

	list1 := []*TreeNode{}
	list2 := []*TreeNode{}

	list1 = append(list1, root)
	isLeftToRight := true

	appendVal := func(l []int, a int) []int {
		if isLeftToRight {
			l = append(l, a)
		} else {
			l = append([]int{a}, l...)
		}
		return l
	}

	for len(list1) != 0 || len(list2) != 0 {
		layer1 := make([]int, 0)
		isLeftToRight = true
		for _, item := range list1 {
			layer1 = appendVal(layer1, item.Val)
			if item.Left != nil {
				list2 = append(list2, item.Left)
			}
			if item.Right != nil {
				list2 = append(list2, item.Right)
			}
		}
		if len(layer1) != 0 {
			ans = append(ans, layer1)
		}
		list1 = []*TreeNode{}

		isLeftToRight = false
		layer2 := make([]int, 0)
		for _, item := range list2 {
			layer2 = appendVal(layer2, item.Val)
			if item.Left != nil {
				list1 = append(list1, item.Left)
			}
			if item.Right != nil {
				list1 = append(list1, item.Right)
			}
		}
		if len(layer2) != 0 {
			ans = append(ans, layer2)
		}
		list2 = []*TreeNode{}
	}
	return ans
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	dequeue := make([]*TreeNode, 0)
	ans := make([][]int, 0)
	var layerNum int

	if root == nil {
		return ans
	}

	dequeue = append(dequeue, root)
	for len(dequeue) != 0 {
		layer := make([]int, 0)
		tmp := make([]*TreeNode, 0)
		for i := len(dequeue) - 1; i >= 0; i-- {
			layer = append(layer, dequeue[i].Val)
			if layerNum%2 == 0 {
				//left -> right
				if dequeue[i].Left != nil {
					tmp = append(tmp, dequeue[i].Left)
				}
				if dequeue[i].Right != nil {
					tmp = append(tmp, dequeue[i].Right)
				}
			} else {
				// right -> Left
				if dequeue[i].Right != nil {
					tmp = append(tmp, dequeue[i].Right)
				}
				if dequeue[i].Left != nil {
					tmp = append(tmp, dequeue[i].Left)
				}
			}
		}
		dequeue = tmp
		ans = append(ans, layer)
		layerNum++
	}
	return ans
}
