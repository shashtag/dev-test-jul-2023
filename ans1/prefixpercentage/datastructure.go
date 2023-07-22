package prefixpercentage

// trie enters all the stings and stores them in form of a tree
type trieNode struct {
	children  [27]*trieNode
	eow_count int
}
type trie struct {
	root *trieNode

	// stores the number of ingests
	ingests int

	// stores the occurrence of a word in the trie
	dict map[string]bool
}

func NewTrie() *trie {
	return &trie{root: &trieNode{eow_count: 0}, ingests: 0, dict: make(map[string]bool)}
}
