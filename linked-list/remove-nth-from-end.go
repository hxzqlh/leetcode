package main

// head->1->2->3->4->5
// head->1->2->3->5
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head.Next
	slow := head

	i := 1
	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
		i++
	}

	if i <= n {
		return head
	}

	slow.Next = slow.Next.Next

	return head
}
