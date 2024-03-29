package list

import "fmt"

// 不带头节点的单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	cur := head
	for ; cur != nil; cur = cur.Next {
		if cur.Next != nil {
			fmt.Printf("%v->", cur.Val)
		} else {
			fmt.Print(cur.Val)
		}
	}
	fmt.Println()
}

func MakeList(arr []int) *ListNode {
	var head *ListNode
	var next *ListNode
	for i := len(arr) - 1; i >= 0; i-- {
		head = &ListNode{arr[i], nil}
		head.Next = next
		next = head
	}
	return head
}

func IsEqual(a *ListNode, b *ListNode) bool {
	p := a
	q := b

	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}

	if p != nil || q != nil {
		return false
	}

	return true
}

func Len(head *ListNode) int {
	var len int
	for cur := head; cur != nil; cur = cur.Next {
		len++
	}
	return len
}

func FindTail(head *ListNode) *ListNode {
	var p *ListNode
	for cur := head; cur != nil; cur = cur.Next {
		p = cur
	}
	return p
}

// 快慢指针法
func FindMid(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
