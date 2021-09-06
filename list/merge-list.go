package list

// 将所有链表(已升序排列)合并到一个升序链表中，返回合并后的链表。
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//	1->4->5,
//	1->3->4,
//	2->6
// ]
// 将它们合并到一个有序链表中得到：
// 1->1->2->3->4->4->5->6

// 顺序合并
func mergeKLists(lists []*ListNode) *ListNode {
	ret := new(ListNode)
	for _, v := range lists {
		ret = mergeList(ret, v)
	}
	return ret
}

// 归并合并
func mergeKLists2(lists []*ListNode) *ListNode {
	return mergeHelp(lists, 0, len(lists)-1)
}

func mergeHelp(lists []*ListNode, begin int, end int) *ListNode {
	if end == begin {
		return lists[begin]
	}
	if begin > end {
		return new(ListNode)
	}
	mid := (begin + end) >> 1
	return mergeList(mergeHelp(lists, begin, mid), mergeHelp(lists, mid+1, end))
}

func mergeList(a *ListNode, b *ListNode) *ListNode {
	ret := new(ListNode)

	cur := ret
	p, q := a, b
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}

	if p != nil {
		cur.Next = p
	}

	if q != nil {
		cur.Next = q
	}

	return ret.Next
}
