package sorting

func MergeSort(v []int) []int {
	if len(v) < 2 {
		return v
	}

	left := v[:len(v)/2]
	right := v[len(v)/2:]
	return merge(MergeSort(left), MergeSort(right))

}

func merge(l, r []int) []int {
	pl, pr := 0, 0
	res := []int{}

	for pl < len(l) && pr < len(r) {
		if l[pl] < r[pr] {
			res = append(res, l[pl])
			pl++
		} else {
			res = append(res, r[pr])
			pr++
		}
	}

	for pl < len(l) {
		res = append(res, l[pl])
		pl++
	}

	for pr < len(r) {
		res = append(res, r[pr])
		pr++
	}

	return res
}
