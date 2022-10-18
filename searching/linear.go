package searching

import "os/exec"

func LinearSearch(v []int, target int) (int, error) {
	for i, item := range v {
		if item == target {
			return i, nil
		}
	}
	return -1, exec.ErrNotFound
}
