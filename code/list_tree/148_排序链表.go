package list_tree

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

// 合并两个有序链表
func mergeTwoList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := &ListNode{}
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

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	mid := slow

	return mergeTwoList(sort(head, mid), sort(mid, tail))
}
