package sorting

import (
	"testing"

	"github.com/hxzqlh/leetcode/list"
)

func TestQuickSort(t *testing.T) {
	a := []int{1, 4, 7, 2, 5, 6, 3, 6, 5}
	QuickSort(a, 0, len(a)-1)
	t.Log(a)
}

func TestListQuickSort(t *testing.T) {
	a := []int{1, 4, 7, 4, 5, 3, 6, 5}
	head := list.MakeList(a)
	list.PrintList(head)

	ListQuickSort(head, nil)
	list.PrintList(head)
}