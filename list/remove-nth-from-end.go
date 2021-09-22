package list

// dummy->1->2->3->4->5 移除倒数第 2 个
// dummy->1->2->3->5
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	fast, slow := head, dummy
	i := 1
	for fast != nil {
		// fast 先走 n+1 步，只有这样同时移动的时候 slow 才能指向删除节点的上一个节点（方便做删除操作）
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}
		i++
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
