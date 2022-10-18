package sorting

func partition(v []int) int {
	if len(v) < 2 {
		return 0
	}
	p := 0

	for i := 1; i < len(v); i++ {
		if v[i] < v[p] {
			v[p+1], v[i] = v[i], v[p+1]
			v[p], v[p+1] = v[p+1], v[p]
			p++
		}
	}
	return p
}

func sort(v []int) {
	if len(v) < 2 {
		return
	}
	pivot := partition(v)
	sort(v[:pivot])
	sort(v[pivot+1:])
}

func QuickSort(v []int) []int {
	sort(v)
	return v
}
