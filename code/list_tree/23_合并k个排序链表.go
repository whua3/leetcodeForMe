package list_tree

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	for _, item := range lists {
		ans = merge(item, ans)
	}
	return ans
}

func merge(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := &ListNode{} // 注意要定义head，用p指向head，方便结果返回
	p := head
	l1, l2 := list1, list2

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}

	if l1 != nil {
		p.Next = l1
	}

	if l2 != nil {
		p.Next = l2
	}

	return head.Next
}
