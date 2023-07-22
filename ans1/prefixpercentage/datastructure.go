package prefixpercentage

// stores the occurrence of a word in the trie
var dict = make(map[string]bool)

// stores the number of ingests
var ingests = 0

// trie enters all the stings and stores them in form of a tree
type trieNode struct {
	children  [27]*trieNode
	eow_count int
}
type trie struct{ root *trieNode }

func NewTrie() *trie {
	return &trie{root: &trieNode{eow_count: 0}}
}
