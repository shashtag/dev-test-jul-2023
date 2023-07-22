package prefixpercentage

import (
	"strings"
)

func (t *trie) Ingest(input string) {
	t.ingests++

	// record all words in dict
	// Time and Space - O ( w ) where w is the number of words in the input string
	words := strings.Split(input, ":")
	for _, word := range words {
		t.dict[word] = true
	}

	//insert string into trie
	// Time and Space - O ( n ) where n is the length of the input string
	current := t.root
	for _, letter := range input {
		index := letter - 'a'
		if letter == ':' {
			current.eow_count++
			index = 26
		}

		if current.children[index] == nil {
			current.children[index] = &trieNode{eow_count: 0}
		}
		current = current.children[index]

	}

	current.eow_count++

}
