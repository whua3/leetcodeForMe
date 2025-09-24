package list_tree

import "strings"

func reverseList(head *ListNode) *ListNode {
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

	strings.TrimSpace()
	return p
}
