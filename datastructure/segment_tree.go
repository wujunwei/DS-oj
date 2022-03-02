package datastructure

// 注：对于指针写法，必要时禁止 GC，能加速不少
// func init() { debug.SetGCPercent(-1) }

// EXTRA: 线段树优化建图
// 每个位置对应着 O(logn) 个线段树上的节点，每个区间可以拆分成至多 O(logn) 个线段树上的区间
// 这个性质可以用于优化建图

// l 和 r 也可以写到方法参数上，实测二者在执行效率上无异
// 考虑到 debug 和 bug free 上的优点，写到结构体参数中
type seg []struct {
	l, r int
	val  int
}

// 单点更新：build 和 update 通用
func (t seg) set(o, val int) {
	t[o].val = val
}

// 合并两个节点上的数据：maintain 和 query 通用
// 要求操作满足区间可加性
// 例如 + * | & ^ min max gcd mulMatrix 摩尔投票 最大子段和 ...
func (seg) op(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].val = t.op(lo.val, ro.val)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// o=1  1<=i<=n
func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	if m := (t[o].l + t[o].r) >> 1; i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].val
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.op(vl, vr)
}

func (t seg) queryAll() int { return t[1].val }

func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

// EXTRA: 查询整个区间小于 v 的最靠左的位置
// 这里线段树维护的是区间最小值
// 需要先判断 t[1].min < v
func (t seg) queryFirstLessPos(o, v int) int {
	if t[o].l == t[o].r {
		return t[o].l
	}
	if t[o<<1].val < v {
		return t.queryFirstLessPos(o<<1, v)
	}
	return t.queryFirstLessPos(o<<1|1, v)
}

// EXTRA: 查询 [l,r] 上小于 v 的最靠左的位置
// 这里线段树维护的是区间最小值
// 不存在时返回 0
func (t seg) queryFirstLessPosInRange(o, l, r, v int) int {
	if t[o].val >= v {
		return 0
	}
	if t[o].l == t[o].r {
		return t[o].l
	}
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		if pos := t.queryFirstLessPosInRange(o<<1, l, r, v); pos > 0 {
			return pos
		}
	}
	if m < r {
		if pos := t.queryFirstLessPosInRange(o<<1|1, l, r, v); pos > 0 { // 注：这里 pos > 0 的判断可以省略，因为 pos == 0 时最后仍然会返回 0
			return pos
		}
	}
	return 0
}

//

// 延迟标记（区间修改）
// 单个更新操作：
// + min/max https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/A
//           https://codeforces.com/problemset/problem/1321/E
//           https://codeforces.com/problemset/problem/52/C
// + min/max 转换 https://codeforces.com/gym/294041/problem/E
// + max DP https://atcoder.jp/contests/dp/tasks/dp_w
// + ∑ https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/D https://www.luogu.com.cn/problem/P3372
// | & https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/C
// = min https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/E
// = ∑ https://codeforces.com/edu/course/2/lesson/5/2/practice/contest/279653/problem/F https://codeforces.com/problemset/problem/558/E
// max max 离散化 https://codeforces.com/contest/1557/problem/D
// https://codeforces.com/problemset/problem/1114/F
// + 某个区间的不小于 x 的最小下标 https://codeforces.com/edu/course/2/lesson/5/3/practice/contest/280799/problem/C
// =max 求和的 O(log^2) 性质 https://codeforces.com/contest/1439/problem/C
// 矩阵乘法 ∑ https://codeforces.com/problemset/problem/718/C
//
// 多个更新操作复合：
// * + ∑ https://www.luogu.com.cn/problem/P3373 https://leetcode-cn.com/problems/fancy-sequence/
// = + ∑ https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/A

// EXTRA: 多项式更新 Competitive Programmer’s Handbook Ch.28
// 比如区间加等差数列（差分法）https://www.luogu.com.cn/problem/P1438 https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/B
type lazyST []struct {
	l, r int
	todo int64
	sum  int64
}

func (lazyST) op(a, b int64) int64 {
	return a + b // % mod
}

func (t lazyST) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].sum = t.op(lo.sum, ro.sum)
}

func (t lazyST) build(a []int64, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t lazyST) do(o int, add int64) {
	to := &t[o]
	to.todo += add                     // % mod
	to.sum += int64(to.r-to.l+1) * add // % mod
}

func (t lazyST) spread(o int) {
	if add := t[o].todo; add != 0 {
		t.do(o<<1, add)
		t.do(o<<1|1, add)
		t[o].todo = 0
	}
}

// 如果维护的数据（或者判断条件）具有单调性，我们就可以在线段树上二分
// 未找到时返回 n+1
// o=1  [l,r] 1<=l<=r<=n
// https://codeforces.com/problemset/problem/1179/C
func (t lazyST) binarySearch(o, l, r int, val int64) int {
	if t[o].l == t[o].r {
		if val <= t[o].sum {
			return t[o].l
		}
		return t[o].l + 1
	}
	t.spread(o)
	// 注意判断对象是当前节点还是子节点
	if val <= t[o<<1].sum {
		return t.binarySearch(o<<1, l, r, val)
	}
	return t.binarySearch(o<<1|1, l, r, val)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) update(o, l, r int, add int64) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o, add)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r, add)
	}
	if m < r {
		t.update(o<<1|1, l, r, add)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t lazyST) query(o, l, r int) int64 {
	if l <= t[o].l && t[o].r <= r {
		return t[o].sum
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.op(vl, vr)
}

func (t lazyST) queryAll() int64 { return t[1].sum }

// a 从 0 开始
func newLazySegmentTree(a []int64) lazyST {
	t := make(lazyST, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}

// EXTRA: 适用于需要提取所有元素值的场景
func (t lazyST) spreadAll(o int) {
	if t[o].l == t[o].r {
		return
	}
	t.spread(o)
	t.spreadAll(o << 1)
	t.spreadAll(o<<1 | 1)
}

//

// 动态开点线段树·其一·单点修改
// LC327 https://leetcode-cn.com/problems/count-of-range-sum/
// rt := &stNode{l: 1, r: 1e9}
type stNode struct {
	lo, ro *stNode
	l, r   int
	sum    int64
}

func (o *stNode) get() int64 {
	if o != nil {
		return o.sum
	}
	return 0 // inf
}

func (stNode) op(a, b int64) int64 {
	return a + b //
}

func (o *stNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

func (o *stNode) build(a []int64, l, r int) {
	o.l, o.r = l, r
	if l == r {
		o.sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	o.lo = &stNode{}
	o.lo.build(a, l, m)
	o.ro = &stNode{}
	o.ro.build(a, m+1, r)
	o.maintain()
}

func (o *stNode) update(i int, add int64) {
	if o.l == o.r {
		o.sum += add //
		return
	}
	m := (o.l + o.r) >> 1
	if i <= m {
		if o.lo == nil {
			o.lo = &stNode{l: o.l, r: m}
		}
		o.lo.update(i, add)
	} else {
		if o.ro == nil {
			o.ro = &stNode{l: m + 1, r: o.r}
		}
		o.ro.update(i, add)
	}
	o.maintain()
}

func (o *stNode) query(l, r int) int64 {
	if o == nil || l > o.r || r < o.l {
		return 0 // inf
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}

// 动态开点线段树·其二·延迟标记（区间修改）
// https://codeforces.com/problemset/problem/915/E（注：此题有多种解法）
// https://codeforces.com/edu/course/2/lesson/5/4/practice/contest/280801/problem/F https://www.luogu.com.cn/problem/P5848
//（内存受限）https://codeforces.com/problemset/problem/1557/D
// rt := &lazyNode{l: 1, r: 1e9}
type lazyNode struct {
	lo, ro *lazyNode
	l, r   int
	sum    int64
	todo   int64
}

func (o *lazyNode) get() int64 {
	if o != nil {
		return o.sum
	}
	return 0 // inf
}

func (lazyNode) op(a, b int64) int64 {
	return a + b //
}

func (o *lazyNode) maintain() {
	o.sum = o.op(o.lo.get(), o.ro.get())
}

func (o *lazyNode) build(a []int64, l, r int) {
	o.l, o.r = l, r
	if l == r {
		o.sum = a[l-1]
		return
	}
	m := (l + r) >> 1
	o.lo = &lazyNode{}
	o.lo.build(a, l, m)
	o.ro = &lazyNode{}
	o.ro.build(a, m+1, r)
	o.maintain()
}

func (o *lazyNode) do(add int64) {
	o.todo += add                   // % mod
	o.sum += int64(o.r-o.l+1) * add // % mod
}

func (o *lazyNode) spread() {
	m := (o.l + o.r) >> 1
	if o.lo == nil {
		o.lo = &lazyNode{l: o.l, r: m}
	}
	if o.ro == nil {
		o.ro = &lazyNode{l: m + 1, r: o.r}
	}
	if add := o.todo; add != 0 {
		o.lo.do(add)
		o.ro.do(add)
		o.todo = 0 // -1
	}
}

func (o *lazyNode) update(l, r int, add int64) {
	if l <= o.l && o.r <= r {
		o.do(add)
		return
	}
	o.spread()
	m := (o.l + o.r) >> 1
	if l <= m {
		o.lo.update(l, r, add)
	}
	if m < r {
		o.ro.update(l, r, add)
	}
	o.maintain()
}

func (o *lazyNode) query(l, r int) int64 {
	// 对于不在线段树中的点，应按照题意来返回
	if o == nil || l > o.r || r < o.l {
		return 0 // inf
	}
	if l <= o.l && o.r <= r {
		return o.sum
	}
	o.spread()
	return o.op(o.lo.query(l, r), o.ro.query(l, r))
}

// EXTRA: 线段树合并
// https://www.luogu.com.cn/problem/P5494
// todo 一些题目 https://www.luogu.com.cn/blog/styx-ferryman/xian-duan-shu-ge-bing-zong-ru-men-dao-fang-qi
//   https://codeforces.com/blog/entry/83969
//   https://www.luogu.com.cn/problem/P4556
//   https://www.luogu.com.cn/problem/P5298
//   https://codeforces.com/problemset/problem/600/E
// rt = rt.merge(rt2)
func (o *stNode) merge(b *stNode) *stNode {
	if o == nil {
		return b
	}
	if b == nil {
		return o
	}
	if o.l == o.r {
		// 按照所需合并，如加法
		o.sum += b.sum
		return o
	}
	o.lo = o.lo.merge(b.lo)
	o.ro = o.ro.merge(b.ro)
	o.maintain()
	return o
}

// EXTRA: 线段树分裂
// 将区间 [l,r] 从 o 中分离到 b 上
// https://www.luogu.com.cn/blog/cyffff/talk-about-segument-trees-split
// https://www.luogu.com.cn/problem/P5494
// rt, rt2 := rt.split(nil, l, r)
func (o *stNode) split(b *stNode, l, r int) (*stNode, *stNode) {
	if o == nil || l > o.r || r < o.l {
		return o, nil
	}
	if l <= o.l && o.r <= r {
		return nil, o
	}
	if b == nil {
		b = &stNode{l: o.l, r: o.r}
	}
	o.lo, b.lo = o.lo.split(b.lo, l, r)
	o.ro, b.ro = o.ro.split(b.ro, l, r)
	o.maintain()
	b.maintain()
	return o, b
}

// 权值线段树求第 k 小
// 调用前需保证 1 <= k <= rt.get()
func (o *stNode) kth(k int64) int {
	if o.l == o.r {
		return o.l
	}
	if cntL := o.lo.get(); k <= cntL {
		return o.lo.kth(k)
	} else {
		return o.ro.kth(k - cntL)
	}
}
