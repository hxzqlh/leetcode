package sorting

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	a := []int{1, 4, 7, 2, 5, 6, 3, 6, 5}
	QuickSort(a, 0, len(a)-1)
	t.Log(a)
}
