package datastructure

import (
	"math"
	"math/rand"
)

const (
	maxLevel = 16
	maxRand  = 65535.0
)

func randLevel() int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

type skipNode struct {
	value int
	right *skipNode
	down  *skipNode
}

type Skiplist struct {
	head   *skipNode
	length int
}

func (s Skiplist) Len() int {
	return s.length
}

func NewSL() Skiplist {
	left := make([]*skipNode, maxLevel)
	right := make([]*skipNode, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &skipNode{-1, nil, nil}
		right[i] = &skipNode{100001, nil, nil}
	}
	for i := maxLevel - 2; i >= 0; i-- {
		left[i].right = right[i]
		left[i].down = left[i+1]
		right[i].down = right[i+1]
	}
	left[maxLevel-1].right = right[maxLevel-1]
	return Skiplist{head: left[0]}
}

func (s *Skiplist) Search(target int) int {
	node := s.head
	for node != nil {
		if node.right.value > target {
			if node.down == nil {
				return node.down.value
			}
			node = node.down
		} else if node.right.value < target {
			if node.right == nil {
				return node.right.value
			}
			node = node.right
		} else {
			return target
		}
	}
	return -1
}

func (s *Skiplist) Add(num int) {
	prev := make([]*skipNode, maxLevel)
	i := 0
	node := s.head
	for node != nil {
		if node.right.value >= num {
			prev[i] = node
			i++
			node = node.down
		} else {
			node = node.right
		}
	}
	n := randLevel()
	arr := make([]*skipNode, n)
	t := &skipNode{-1, nil, nil}
	for i, a := range arr {
		p := prev[maxLevel-n+i]
		a = &skipNode{num, p.right, nil}
		p.right = a
		t.down = a
		t = a
	}
	s.length++
}

func (s *Skiplist) Erase(num int) (ans bool) {
	node := s.head
	for node != nil {
		if node.right.value > num {
			node = node.down
		} else if node.right.value < num {
			node = node.right
		} else {
			s.length--
			ans = true
			node.right = node.right.right
			node = node.down
		}
	}
	return
}
