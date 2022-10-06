package sorting

func Bubble(v []int) []int {
	if len(v) < 2 {
		return v
	}

	for i := 0; i < len(v)-1; i++ {
		for j := i + 1; j < len(v); j++ {
			if v[i] > v[j] {
				tmp := v[i]
				v[i] = v[j]
				v[j] = tmp
			}
		}
	}
	return v
}
