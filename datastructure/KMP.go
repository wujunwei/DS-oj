package datastructure

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//z 函数，拓展 kmp
func zz(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r && z[i-l] < r-i+1 {
			z[i] = z[i-l]
		} else {
			z[i] = max(0, r-i+1)
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				z[i]++
			}
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}
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
