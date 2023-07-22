package prefixpercentage

// stores the occurrence of a word in the trie
var dict = make(map[string]bool)

// trie enters all the stings and stores them in form of a tree
type trieNode struct {
	children  [27]*trieNode
	eow_count int
}
type trie struct {
	root    *trieNode
	ingests int
}

func NewTrie() *trie {
	return &trie{root: &trieNode{eow_count: 0}, ingests: 0}
}
