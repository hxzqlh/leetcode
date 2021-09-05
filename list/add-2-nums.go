package list

// 反向排序
// 输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
// 输出：2 -> 1 -> 9，即912
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode) // 引入头节点

	var carry int // 进位
	var val int   // 相加之和，>=10则对10取模

	pre := ret
	p := l1
	q := l2
	for p != nil || q != nil || carry > 0 {
		val = carry

		if p != nil {
			val += p.Val
			p = p.Next
		}
		if q != nil {
			val += q.Val
			q = q.Next
		}

		carry = 0
		if val >= 10 {
			val = val % 10
			carry = 1
		}

		pre.Next = &ListNode{val, nil}
		pre = pre.Next
	}

	return ret.Next
}

// 正向排序
// 输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
// 输出：9 -> 1 -> 2，即912
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// 逆序后，相加，再将结果逆序，即为正序结果
	r1 := reverseList(l1)
	r2 := reverseList(l2)
	r3 := addTwoNumbers(r1, r2)
	return reverseList(r3)
}
