package sorting

import (
	"github.com/hxzqlh/leetcode/list"
)

// 单链表快排
func ListQuickSort(head *list.ListNode, end *list.ListNode) {
	if head != end {
		pos := ListPartion(head, end)
		ListQuickSort(head, pos)
		ListQuickSort(pos.Next, end)
	}
}

// 3->1->5->2->4->0
// 3->1->5->2->4->0
// 3->1->2->5->4->0
// 3->1->2->0->4->5
// 0->1->2->3->4->5
func ListPartion(head *list.ListNode, end *list.ListNode) *list.ListNode {
	key := head.Val
	p := head
	q := head.Next
	for q != end {
		if q.Val < key {
			p = p.Next
			p.Val, q.Val = q.Val, p.Val
		}
		q = q.Next
	}
	p.Val, head.Val = head.Val, p.Val
	// list.PrintList(head)
	return p
}
