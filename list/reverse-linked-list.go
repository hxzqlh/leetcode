package list

//       1->2->3->4->5->null
// null<-1<-2<-3<-4<-5
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 递归
// 1->2->3->4->5
// 1->[2<-3<-4<-5]
// 1<-[2<-3<-4<-5]
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := ReverseList2(head.Next) // 2<-3<-4<-5(p)
	head.Next.Next = head        // 1<-2
	head.Next = nil              // null<-1
	return p
}

// 递归
func ReverseList3(head *ListNode) *ListNode {
	return reverseHelp(nil, head)
}

func reverseHelp(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}

	next := head.Next
	head.Next = pre
	return reverseHelp(head, next)
}
