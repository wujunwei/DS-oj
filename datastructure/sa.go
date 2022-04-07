package datastructure

import "sort"

type suffix struct {
	c   rune
	idx int // a tuple (c, idx) indicates a suffix in source string
}

type SA struct {
	src    string
	sa     []int
	rk     []int
	height []int
}

func NewSA(src string) *SA {
	n := len(src)
	sa := make([]suffix, n)
	for i, c := range src {
		sa[i] = suffix{c, i}
	}
	// init sa: sort by ascii
	sort.Slice(sa, func(i, j int) bool {
		return sa[i].c < sa[j].c
	})

	rk := make([]int, n)
	// init rk: incr rk each time sa[i].c changes
	rk[sa[0].idx] = 1
	for i := 1; i < n; i++ {
		rk[sa[i].idx] = rk[sa[i-1].idx]
		if sa[i].c != sa[i-1].c {
			rk[sa[i].idx]++
		}
	}

	// binary lifting
	// k is length of substr needed comparing
	// when the last of rk is n, terminate
	for k := 2; rk[sa[n-1].idx] < n; k <<= 1 {
		// sort sa by rank tuple (first, second)
		sort.Slice(sa, func(i, j int) bool {
			ii, jj := sa[i].idx, sa[j].idx
			// compare first part
			if rk[ii] != rk[jj] {
				return rk[ii] < rk[jj]
			}
			// compare second part
			if ii+k/2 >= n {
				return true // special case: second part not exists
			}
			if jj+k/2 >= n {
				return false
			}
			return rk[ii+k/2] < rk[jj+k/2]
		})
		// update rk
		rk[sa[0].idx] = 1
		for i := 1; i < n; i++ {
			cur, pre := sa[i].idx, sa[i-1].idx
			rk[cur] = rk[pre]
			// incr rk each time the substring changes. notice that an out-of-bound case indicates a change
			if cur+k > n || pre+k > n || src[cur:cur+k] != src[pre:pre+k] {
				rk[cur]++
			}
		}
	}
	realSA := make([]int, n)
	for i, v := range sa {
		realSA[i] = v.idx
	}
	// compute height
	height := make([]int, n)
	k := 0
	for i := 0; i < n; i++ {
		if rk[i] == 1 {
			continue
		}
		j := realSA[rk[i]-2]
		if k > 0 {
			k-- // invariant condition: height[rank[i]]>=height[rank[i-1]]-1
		}
		for i+k < n && j+k < n && src[i+k] == src[j+k] {
			k++
		}
		height[rk[i]-1] = k
	}

	return &SA{src, realSA, rk, height}
}
