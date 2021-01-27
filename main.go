package main

import "math"

type union struct {
	parent *union
	val    int
}

func (u *union) getParent() *union {
	if u == nil || u == u.parent {
		return u
	}
	u.parent = u.parent.getParent()
	return u.parent
}

func main() {

}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if len(s1) < len(s2) {
		return 0
	}
	a1, a2 := [26]int{}, [26]int{}
	for i := 0; i < len(s1); i++ {
		a1[s1[i]-'a']++
	}
	for i := 0; i < len(s2); i++ {
		a2[s1[i]-'a']++
	}
	c1, c2 := math.MaxInt32, 1
	for i := 0; i < 26; i++ {
		if a2[i] == 0 {
			continue
		}
		if a1[i] < a2[i] {
			return 0
		}
		if float64(a1[i])/float64(a2[i]) < float64(c1)/float64(c2) {
			c1, c2 = a1[i], a2[i]
		}
	}
	return (n1 * c1) / (n2 * c2)
}
