package prefixpercentage

var dict = make(map[string]bool)
var ingests = 0

type trieNode struct {
	children  [26]*trieNode
	eow_count int
}
type trie struct{ root *trieNode }

func NewTrie() *trie {
	return &trie{root: &trieNode{eow_count: 0}}
}
