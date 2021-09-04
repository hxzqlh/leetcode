package main

// 		 1->2->3->4->5->null
// null<-1<-2<-3<-4<-5
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 递归
// p := reverseList(head.Next)
// head.Next = p
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// 递归
func reverseList3(head *ListNode) *ListNode {
	return help(nil, head)
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}

	next := head.Next
	head.Next = pre
	return help(head, next)
}
