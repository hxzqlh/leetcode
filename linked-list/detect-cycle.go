package main

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head.Next, head.Next

	// 找相遇节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	// 找环形入口节点
	slow = head.Next
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
