package main

import "github.com/shashtag/dev-test-jul-2023.git/ans1/prefixpercentage"

func main() {
	trie := prefixpercentage.NewTrie()

	trie.Ingest("leoluca:uk:dev")
	trie.Ingest("leoluca:hk:design")
	trie.Ingest("leoluca:hk:pm")
	trie.Ingest("leoluca:hk:dev")
	trie.Ingest("skymaker")

	trie.Appearance("leoluca")
	trie.Appearance("leoluca:hk")

	trie.Ingest("skymaker:london:ealing:dev")
	trie.Ingest("skymaker:london:croydon")
	trie.Ingest("skymaker:london:design")
	trie.Ingest("skymaker:man:pm")
	trie.Ingest("skymaker:man:pm")

	trie.Appearance("skymaker:man")

}
