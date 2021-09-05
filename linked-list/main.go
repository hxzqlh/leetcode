package main

import "fmt"

// 带头节点的单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	cur := head.Next
	for ; cur != nil; cur = cur.Next {
		fmt.Print(cur.Val)
	}
	fmt.Println()
}

func main() {
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}
	head := &ListNode{-1, n1} // 1->2->3->4->5

	//PrintList(head)
	//head = removeNthFromEnd(head, 6)
	//PrintList(head)

	//n6 := &ListNode{6, n4}
	//n7 := &ListNode{7, n6}
	//head2 := &ListNode{-1, n7} // 7->6->4->5
	//node := getIntersectionNode(head, head2)
	//fmt.Println(node.Val)

	//PrintList(head)
	//head = swapPairs2(head)
	//PrintList(head)

	//             |---------|
	// 1->2->3->4->5->8->9->10
	//n10 := &ListNode{10, n5}
	//n9 := &ListNode{9, n10}
	//n8 := &ListNode{8, n9}
	//n5.Next = n8
	//node := detectCycle(head)
	//fmt.Println(node.Val)

	n6 := &ListNode{6, nil}
	n7 := &ListNode{4, n6}
	n8 := &ListNode{2, n7}
	head3 := &ListNode{-1, n8} // 2->4->6
	//ret := mergeList(head, head3)
	//PrintList(ret)

	n9 := &ListNode{5, nil}
	n10 := &ListNode{3, n9}
	n11 := &ListNode{1, n10}
	head4 := &ListNode{-1, n11} // 1->3->5
	ret := mergeKLists2([]*ListNode{head, head3, head4})
	PrintList(ret)
}
