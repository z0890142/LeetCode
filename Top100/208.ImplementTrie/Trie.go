package ImplementTrie

type Trie struct {
	root *trieStorage
}

type trieStorage struct {
	alphabet [26]*trieStorage
	end      bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: &trieStorage{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for _, digit := range word {
		number := digit - 'a'
		if nil == node.alphabet[number] { // if the next position doesn't exist, make a new place
			node.alphabet[number] = &trieStorage{}
		}
		node = node.alphabet[number]
	}
	node.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, digit := range word {
		number := digit - 'a'
		if nil == node.alphabet[number] { // if the next position doesn't exist, not in
			return false
		}
		node = node.alphabet[number]
	}
	return node.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, digit := range prefix {
		number := digit - 'a'
		if nil == node.alphabet[number] { // if the next position doesn't exist, not in
			return false
		}
		node = node.alphabet[number]
	}
	return true
}
