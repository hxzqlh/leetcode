package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct{
		n1 []int
		n2 []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{}, []int{}},
	}

	for _, v := range cases {
		l1 := MakeList(v.n1)
		l2 := MakeList(v.n2)
		PrintList(l1)
		r1 := ReverseList(l1)
		PrintList(r1)
		assert.Equal(t, IsEqual(r1, l2), true)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct{
		n1 []int
		n  int
		n2 []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
	}

	for _, v := range cases {
		l1 := MakeList(v.n1)
		l2 := MakeList(v.n2)
		ret := RemoveNthFromEnd(l1, v.n)
		PrintList(l2)
		PrintList(ret)
		assert.Equal(t, IsEqual(ret, l2), true)
	}
}

func TestAddTwoNums(t *testing.T) {
	cases := []struct{
		n1 []int
		n2 []int
		n3 []int
	}{
		{[]int{7, 1, 6}, []int{5, 9, 2}, []int{2, 1, 9}},
		{[]int{5, 4, 3, 2, 1}, []int{6, 4, 2}, []int{1, 9, 5, 2, 1}},
		{[]int{5, 1, 9}, []int{5, 6, 1}, []int{0, 8, 0, 1}},
		{[]int{9, 9, 9}, []int{1}, []int{0, 0, 0, 1}},
	}

	for _, v := range cases {
		l1 := MakeList(v.n1)
		l2 := MakeList(v.n2)
		l3 := MakeList(v.n3)
		ret := AddTwoNumbers(l1, l2)
		PrintList(l3)
		PrintList(ret)
		assert.Equal(t, IsEqual(ret, l3), true)
	}
}

func TestAddTwoNums2(t *testing.T) {
	cases := []struct{
		n1 []int
		n2 []int
		n3 []int
	}{
		{[]int{6, 1, 7}, []int{2, 9, 5}, []int{9, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 4, 6}, []int{1, 2, 5, 9, 1}},
		{[]int{9, 1, 5}, []int{1, 6, 5}, []int{1, 0, 8, 0}},
		{[]int{9, 9, 9}, []int{1}, []int{1, 0, 0, 0}},
	}

	for _, v := range cases {
		l1 := MakeList(v.n1)
		l2 := MakeList(v.n2)
		l3 := MakeList(v.n3)
		ret := AddTwoNumbers2(l1, l2)
		PrintList(l3)
		PrintList(ret)
		assert.Equal(t, IsEqual(ret, l3), true)
	}
}

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
	node := DetectCycle(n1)
	fmt.Println(node.Val)
}

func TestSort(t *testing.T) {
	a := []int{1, 5, 2, 0, 3, 2}

	l1 := MakeList(a)
	PrintList(l1)

	ret := ListMergeSort2(l1)
	PrintList(ret)
}

func TestMergeList(t *testing.T) {
	cases := []struct{
		n1 []int
		n2 []int
		n3 []int
	}{
		{[]int{1, 2}, []int{2, 3, 5}, []int{1, 2, 2, 3, 5}},
		{[]int{1}, []int{2}, []int{1, 2}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{}, []int{}, []int{}},
	}

	for _, v := range cases {
		l1 := MakeList(v.n1)
		l2 := MakeList(v.n2)
		l3 := MakeList(v.n3)
		ret := MergeList(l1, l2)
		PrintList(l3)
		PrintList(ret)
		assert.Equal(t, IsEqual(ret, l3), true)
	}
}

func TestMergeKLists(t *testing.T) {
	cases := []struct{
		n1 [][]int
		n2 []int
	}{
		{[][]int{{1, 2}, {2, 3, 5}}, []int{1, 2, 2, 3, 5}},
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
	}

	for _, v := range cases {
		lists := []*ListNode{}
		for _, elems := range v.n1 {
			lists = append(lists, MakeList(elems))
		}
		ret := MergeKLists(lists)
		l2 := MakeList(v.n2)
		PrintList(l2)
		PrintList(ret)
		assert.Equal(t, IsEqual(ret, l2), true)
	}
}
