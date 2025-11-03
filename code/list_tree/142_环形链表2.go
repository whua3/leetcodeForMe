package list_tree

func detectCycle1(head *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]struct{}, 0)
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return head
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}

	return nil
}
