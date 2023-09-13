package main

import "fmt"

const size = 26

// Node represent each node in the trie it contain an array of 26 length
type TrieNode struct {
	children [size]*TrieNode
	isEnd    bool
}

//trie  represent a trie and has a pointer to root node

type Trie struct {
	root *TrieNode
}

// initTrie will create new trie
func InitTrie() *Trie {
	result := &Trie{}
	result.root = &TrieNode{}

	return result

}

// insert
func (t *Trie) insert(w string) {
	wordLength := len(w)
	temp := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'

		if temp.children[charIndex] == nil {
			newNode := new(TrieNode)
			temp.children[charIndex] = newNode
		}
		temp = temp.children[charIndex]
	}
	temp.isEnd = true
}

//display

// search
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	temp := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if temp.children[charIndex] == nil {
			return false

		}
		temp = temp.children[charIndex]

	}
	if temp.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()

	myTrie.insert("jerin")
	myTrie.insert("jerrin")
	myTrie.insert("sanika")
	myTrie.insert("clever")
	myTrie.insert("awesome")
	fmt.Println(myTrie.Search("jerin"))

}
