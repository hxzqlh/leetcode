package sorting

func QuickSort(a []int, p int, r int) {
	if p < r {
		pos := partition2(a, p, r)
		QuickSort(a, 0, pos-1)
		QuickSort(a, pos+1, r)
	}
}

func partition(a []int, p int, r int) int {
	key := a[r]
	i := p - 1
	for j := p; j < r; j++ {
		if a[j] < key {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func partition2(a []int, p int, r int) int {
	key := a[p]
	for p < r {
		for p < r && a[r] > key {
			r--
		}
		if p < r {
			a[p] = a[r]
			p++
		}

		for p < r && a[p] < key {
			p++
		}
		if p < r {
			a[r] = a[p]
			r--
		}
	}
	a[p] = key
	return p
}
