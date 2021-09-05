package sorting

import (
	"testing"

	"github.com/hxzqlh/leetcode/list"
)

func TestListQuickSort(t *testing.T) {
	a := []int{1, 4, 7, 4, 5, 3, 6, 5}
	head := list.MakeList(a)
	list.PrintList(head)

	ListQuickSort(head, nil)
	list.PrintList(head)
}
