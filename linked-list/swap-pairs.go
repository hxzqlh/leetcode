package main

// head->1->2->3->4->5
// head->2->3->4       [head->2]
// head->2->1          [2->1]
// head->2->1->3->4->5 [1->3]
// ...
// head->2->1->4->3->5
func swapPairs(head *ListNode) *ListNode {
	pre := head
	cur := head.Next
	for cur != nil && cur.Next != nil {
		next := cur.Next.Next

		pre.Next = cur.Next // head->2
		cur.Next.Next = cur // 2->1
		cur.Next = next     // 1->3

		pre = cur
		cur = next
	}
	return head
}

// 递归版本
func swapPairs2(head *ListNode) *ListNode {
	head.Next = swapHelp(head.Next)
	return head
}

func swapHelp(cur *ListNode) *ListNode {
	if cur == nil || cur.Next == nil {
		return cur
	}
	next := cur.Next
	cur.Next = swapHelp(next.Next)
	next.Next = cur
	return next
}
