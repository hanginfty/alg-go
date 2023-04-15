package searching

import (
	"math"
)

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

func Calcsqrt2() float64 {
	l, r := 1.414, 1.415

	precision := math.Pow10(-6)

	for l+precision <= r {
		m := l + (r-l)/2

		if m*m < 2 {
			l = m
		} else {
			r = m
		}
	}

	res := r

	return res
}
