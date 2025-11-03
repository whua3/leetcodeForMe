package list_tree

// 直接用两个指针遍历两个链表，val相加注意进位
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	carry := 0
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sumTail := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10

		p.Next = &ListNode{Val: sumTail}
		p = p.Next
	}
	if carry > 0 {
		p.Next = &ListNode{Val: carry}
	}
	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	add := 0

	for l1 != nil || l2 != nil || add != 0 {
		tmp := add
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{
			Val: tmp % 10,
		}
		p = p.Next
		add = tmp / 10
	}

	return head.Next
}
