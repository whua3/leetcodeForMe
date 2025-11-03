package list_tree

// 官方题解一：递归
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs1(newHead.Next)
	newHead.Next = head

	return newHead
}

// 官方题解二：迭代
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head} // 重点思路：初始化一个dummyHead

	p := dummyHead

	for p.Next != nil && p.Next.Next != nil {
		q := p.Next
		r := q.Next

		p.Next = r
		q.Next = r.Next
		r.Next = q

		p = q
	}

	return dummyHead.Next
}
