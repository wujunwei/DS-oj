package datastructure

type Trie struct {
	path     string
	children map[byte]*Trie
	isWord   bool
}

func Find(t *Trie, a string) []string {
	for i := 0; i < len(a); i++ {
		if t.children[a[i]] == nil {
			return []string{}
		}
		t = t.children[a[i]]
	}
	return t.GetAll()
}

func (t Trie) GetAll() []string {
	var ans []string
	if t.isWord {
		ans = append(ans, t.path)
	}
	for _, trie := range t.children {
		ans = append(ans, trie.GetAll()...)
	}
	return ans
}

func BuildTrie(dict []string) *Trie {
	t := &Trie{children: map[byte]*Trie{}}
	for _, s := range dict {
		temp := t
		for i := 0; i < len(s); i++ {
			if temp.children[s[i]] == nil {
				temp.children[s[i]] = &Trie{children: map[byte]*Trie{}}
			}
			temp = temp.children[s[i]]
		}
		temp.path = s
		temp.isWord = true
	}
	return t
}
