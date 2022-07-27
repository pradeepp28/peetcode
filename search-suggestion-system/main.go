package main

import "fmt"

type Node struct {
	isWord   bool
	children []*Node
}

type Trie struct {
	Root         *Node
	resultBuffer []string
}

func (t *Trie) dfsWithPrefix(curr *Node, word string) {
	if len(t.resultBuffer) == 3 {
		return
	}

	if curr.isWord {
		t.resultBuffer = append(t.resultBuffer, word)
	}

	for i := byte('a'); i <= 'z' && len(t.resultBuffer) < 3; i++ {
		if curr.children[i-'a'] != nil {
			// word = string(append([]byte(word), i))
			t.dfsWithPrefix(curr.children[i-'a'], string(append([]byte(word), i)))
		}
	}
}

func NewTrie() *Trie {
	t := &Trie{}
	t.Root = &Node{}
	t.Root.children = make([]*Node, 26)
	return t
}

func (t *Trie) insert(word string) {
	curr := t.Root
	for _, r := range []byte(word) {
		if curr.children[r-'a'] == nil {
			curr.children[r-'a'] = &Node{}
			curr.children[r-'a'].children = make([]*Node, 26)
		}
		curr = curr.children[r-'a']
		// curr.children = make([]*Node, 26)
	}
	curr.isWord = true
}

func (t *Trie) getWordsStartingWith(prefix string) []string {
	curr := t.Root
	for _, r := range []byte(prefix) {
		if curr.children[r-'a'] == nil {
			return nil
		}
		curr = curr.children[r-'a']
	}
	t.resultBuffer = nil
	t.dfsWithPrefix(curr, prefix)
	return t.resultBuffer
}

func suggestedProducts(products []string, searchWord string) [][]string {
	t := NewTrie()
	for _, s := range products {
		t.insert(s)
	}

	var prefix []byte
	var result [][]string
	for _, c := range []byte(searchWord) {
		prefix = append(prefix, c)
		result = append(result, t.getWordsStartingWith(string(prefix)))
	}
	return result
}
func main() {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	result := suggestedProducts(products, "mouse")
	fmt.Printf("%v\n", result)
}
