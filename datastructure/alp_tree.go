package datastructure

type Trie struct {
	path     string
	children [26]*Trie
	isWord   bool
}

func Find(t *Trie, a string) []string {
	for i := 0; i < len(a); i++ {
		if t.children[a[i]-'a'] == nil {
			return []string{}
		}
		t = t.children[a[i]-'a']
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
	t := &Trie{}
	for _, s := range dict {
		temp := t
		for i := 0; i < len(s); i++ {
			if temp.children[s[i]-'a'] == nil {
				temp.children[s[i]-'a'] = &Trie{}
			}
			temp = temp.children[s[i]-'a']
		}
		temp.path = s
		temp.isWord = true
	}
	return t
}
