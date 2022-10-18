package searching

// search the first element that matches in a sorted array.
// that is to say: return the lower bound

func BinarySearch(v []int, target int) (int, error) {
	if len(v) == 0 {
		return -1, ErrNotFound
	}

	l := -1
	r := len(v)

	for l+1 != r {
		m := l + (r-l)/2
		if v[m] < target {
			l = m
		} else {
			r = m
		}
	}
	res := r

	if res == -1 || res == len(v) || v[res] != target {
		return -1, ErrNotFound
	}

	return res, nil
}
