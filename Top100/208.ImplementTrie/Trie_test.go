package ImplementTrie

import (
	"fmt"
	"testing"
)

func Test_Constructor(t *testing.T) {
	action := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"}
	params := [][]string{[]string{}, []string{"apple"}, []string{"apple"}, []string{"app"}, []string{"app"}, []string{"app"}, []string{"app"}}

	var trie Trie
	for index, act := range action {
		if act == "Trie" {
			trie = Constructor()
		} else if act == "insert" {

			trie.Insert(params[index][0])

			fmt.Println("---------------------------------------------")
		} else if act == "search" {

			fmt.Println(trie.Search(params[index][0]))
			fmt.Println("---------------------------------------------")

		} else {
			fmt.Println(trie.StartsWith(params[index][0]))

		}
	}

}
