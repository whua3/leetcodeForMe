package list_tree

// 双指针，fast先走N步
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for fast != nil && n != 0 {
		fast = fast.Next
		n--
	}

	if fast == nil {
		return head.Next // 注意这里，fast为nil说明要移除head
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return head
}
