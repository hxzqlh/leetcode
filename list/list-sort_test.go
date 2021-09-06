package list

import (
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{1, 5, 2, 0, 3, 2}

	l1 := MakeList(a)
	PrintList(l1)

	ret := ListMergeSort2(l1)
	PrintList(ret)
}
