package datastructure

import (
	"math/rand"
)

type node struct {
	Ch       [2]*node
	priority int
	Val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.Val:
		return 0
	case b > o.Val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.Ch[d^1]
	o.Ch[d^1] = x.Ch[d]
	x.Ch[d] = o
	return x
}

type Treap struct {
	root *node
}

func (t *Treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), Val: val}
	}
	d := o.cmp(val)
	o.Ch[d] = t._put(o.Ch[d], val)
	if o.Ch[d].priority > o.priority {
		o = o.rotate(d ^ 1)
	}
	return o
}

func (t *Treap) Put(val int) {
	t.root = t._put(t.root, val)
}

func (t *Treap) _delete(o *node, val int) *node {
	if d := o.cmp(val); d >= 0 {
		o.Ch[d] = t._delete(o.Ch[d], val)
		return o
	}
	if o.Ch[1] == nil {
		return o.Ch[0]
	}
	if o.Ch[0] == nil {
		return o.Ch[1]
	}
	d := 0
	if o.Ch[0].priority > o.Ch[1].priority {
		d = 1
	}
	o = o.rotate(d)
	o.Ch[d] = t._delete(o.Ch[d], val)
	return o
}

func (t *Treap) Delete(val int) {
	t.root = t._delete(t.root, val)
}

func (t *Treap) LowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.Ch[0]
		case c > 0:
			o = o.Ch[1]
		default:
			return o
		}
	}
	return
}
func (t *Treap) HighBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			o = o.Ch[0]
		case c > 0:
			lb = o
			o = o.Ch[1]
		default:
			return o
		}
	}
	return
}
