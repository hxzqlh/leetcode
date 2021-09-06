package list

// 单链表归并排序,自顶向下
func ListMergeSort(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	// mid: slow
	// [head, mid) [mid, tail)
	return mergeList(sort(head, slow), sort(slow, tail))
}

// 单链表归并排序,自底向上
func ListMergeSort2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	length := Len(head)
	dummyHead := &ListNode{Next: head}
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		for cur != nil {
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}

			head2 := cur.Next
			cur.Next = nil
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}

			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = mergeList(head1, head2)

			for prev.Next != nil {
				prev = prev.Next
			}
			cur = next
		}
	}

	return dummyHead.Next
}
