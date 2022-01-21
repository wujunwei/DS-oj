package datastructure

func KMP(a, b string) bool {
	next := getNext(a)
	for i, j := 0, 0; i < len(b); i++ {
		for j > 0 && a[j] != b[i] {
			j = next[j]
		}
		if a[j] == b[i] {
			j++
		}
		if j == len(a) {
			return true
		}
	}
	return false
}

func getNext(a string) []int {
	next := make([]int, len(a))
	for i, j := 1, 0; i < len(a); i++ {
		for j > 0 && a[i] != a[j] {
			j = next[j-1]
		}
		if a[i] == a[j] {
			j++
		}
		next[i] = j
	}
	return next
}
