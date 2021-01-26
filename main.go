package main

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
