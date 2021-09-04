package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 求长度
	curA, curB := headA, headB
	lenA, lenB := 0, 0
	for ; curA != nil; curA = curA.Next {
		lenA++
	}
	for ; curB != nil; curB = curB.Next {
		lenB++
	}

	// 长度差
	var step int
	var fast, slow *ListNode
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}

	// 长的先走
	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	// 找到相同节点
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
