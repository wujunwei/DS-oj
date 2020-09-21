package data_struct

type Segment struct {
	lazy       int
	l, r       int
	val        int
	leftChild  *Segment
	rightChild *Segment
}

func (s *Segment) Update(start, end, c int) {
	if s.l == start && s.r == end {
		if start != end {
			s.lazy = c
		} else {
			s.val += c
		}
		return
	}
	s.val += (end - start + 1) * c
	middle := (s.r + s.l) >> 1
	if middle >= start {
		if end > middle {
			s.leftChild.Update(start, middle, c)
		} else {
			s.leftChild.Update(start, end, c)
		}
	}
	if end > middle {
		if start > middle {
			s.rightChild.Update(start, end, c)
		} else {
			s.rightChild.Update(middle+1, end, c)
		}
	}
}

func (s *Segment) Get(start, end int) int {
	if s.l == start && s.r == end {
		return s.val + s.lazy*(end-start+1)
	}
	middle := (s.r + s.l) >> 1
	var l, r int
	if middle >= start {
		if end > middle {
			l = s.leftChild.Get(start, middle)
		} else {
			l = s.leftChild.Get(start, end)
		}
	}
	if end > middle {
		if start > middle {
			r = s.rightChild.Get(start, end)
		} else {
			r = s.rightChild.Get(middle+1, end)
		}
	}
	return l + r + (end-start+1)*s.lazy
}

func NewSegment(l int, r int, a []int) *Segment {
	if l == r {
		return &Segment{l: l, r: r, val: a[l]}
	}
	var s = &Segment{0, l, r, 0, NewSegment(l, (l+r)/2, a), NewSegment((l+r)/2+1, r, a)}
	s.val = s.leftChild.val + s.rightChild.val
	return s
}
