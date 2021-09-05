package list

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	// 找相遇节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	// 从头结点和相遇结点，各走一步，直到相遇，相遇点即为环入口点
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
