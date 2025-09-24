package list_tree

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// 找到链表的中点
	mid := middleNode(head)

	// reverse 后半 list
	secondHalf := reverseList2(mid.Next)

	// merge
	l1, l2 := head, secondHalf
	mid.Next = nil
	head = mergeList(l1, l2)
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, q, r := head, head.Next, head.Next.Next
	head.Next = nil
	for q != nil {
		q.Next = p
		p = q
		q = r
		if r != nil {
			r = r.Next
		}
	}
	return p
}

func mergeList(l1, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	r := tmp
	p, q := l1, l2
	for p != nil && q != nil {
		r.Next = p
		r = r.Next
		p = p.Next
		r.Next = q
		r = r.Next
		q = q.Next
	}
	if p != nil {
		r.Next = p
	}
	if q != nil {
		r.Next = q
	}

	return tmp.Next
}
