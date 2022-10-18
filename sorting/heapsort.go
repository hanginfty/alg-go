package sorting

func HeapSort(v []int) []int {
	if len(v) < 2 {
		return v
	}

	for i := len(v)/2 - 1; i >= 0; i-- {
		heapify(v, i)
	}

	for i := len(v) - 1; i > 0; i-- {
		v[i], v[0] = v[0], v[i]
		heapify(v[:i], 0)
	}

	return v
}

func heapify(v []int, root int) {
	max := root
	l, r := 2*root+1, 2*root+2

	if l < len(v) && v[l] > v[max] {
		max = l
	}
	if r < len(v) && v[r] > v[max] {
		max = r
	}

	if max != root {
		v[max], v[root] = v[root], v[max]
		heapify(v, max)
	}
}
