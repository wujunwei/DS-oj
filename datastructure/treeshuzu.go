package datastructure

var c = make([]int, len([]int{})+5)

func lowBit(x int) int {
	return x & (-x)
}

func update(pos int) {
	for pos < len(c) {
		c[pos]++
		pos += lowBit(pos)
	}
}

func query(pos int) int {
	ret := 0
	for pos > 0 {
		ret += c[pos]
		pos -= lowBit(pos)
	}
	return ret
}
