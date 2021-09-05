package main

import (
	"fmt"

	"github.com/hxzqlh/leetcode/sorting"
)

func main() {
	a := []int{1, 4, 7, 2, 5, 6, 3, 6, 5}
	sorting.QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
