package prefixpercentage

import (
	"fmt"
	"strings"
)

func (t *trie) Appearance(input string) {

	// edge case - empty input string
	// Time - O ( 1 )
	if input == "" {
		if t.ingests == 0 {
			fmt.Println("prefix percentage : 0 - No ingests")
			return
		}
		fmt.Println("prefix percentage : 1 - Empty string is prefix of all strings")
		return
	}

	// check if all words are in the map
	// Time - O ( w ) where w is the number of words in the input string
	words := strings.Split(input, ":")
	for _, word := range words {

		if _, ok := t.dict[word]; !ok {
			fmt.Printf("prefix percentage : 0 - Word '%s' not found part of the words in the ingests ", word)
			return
		}
	}

	// check if the prefix is in the trie and print the appearance percentage
	// Time - O ( n ) where n is the length of the input string
	current := t.root
	for _, letter := range input {
		index := letter - 'a'
		if letter == ':' {
			index = 26
		}

		if current.children[index] == nil {
			fmt.Println("prefix percentage : 0 - No match found")
			return
		}
		current = current.children[index]
	}
	if current.eow_count > 0 {
		s := fmt.Sprintf("%.2f", float64(current.eow_count)/float64(t.ingests))
		fmt.Printf("prefix percentage : %s - prefix count = %d, ingests = %d\n", s, current.eow_count, t.ingests)
		return
	}
	fmt.Println("prefix percentage : 0 - No match found")

}
