package prefixpercentage

import (
	"strings"
)

func (t *trie) Ingest(input string) {
	ingests++

	// record all words in dict
	words := strings.Split(input, ":")
	for _, word := range words {
		dict[word] = true
	}

	//insert string into trie
	current := t.root
	for _, letter := range input {
		if letter == ':' {
			current.eow_count++
			continue
		}
		index := letter - 'a'
		if current.children[index] == nil {
			current.children[index] = &trieNode{eow_count: 0}
		}
		current = current.children[index]

	}

	current.eow_count++

}
