package list

import (
	"fmt"
	"testing"
)

func TestDetectCycle(t *testing.T) {
	//             |---------|
	// 1->2->3->4->5->8->9->10
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	n1 := &ListNode{1, n2}
	n10 := &ListNode{10, n5}
	n9 := &ListNode{9, n10}
	n8 := &ListNode{8, n9}
	n5.Next = n8
	node := detectCycle(n1)
	fmt.Println(node.Val)
}
