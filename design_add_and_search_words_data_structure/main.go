package main

import "fmt"

type Node struct {
	isWord   bool
	children []*Node
}

type WordDictionary struct {
	Root *Node
}

func Constructor() WordDictionary {
	d := WordDictionary{
		Root: &Node{},
	}
	d.Root.children = make([]*Node, 26)
	return d
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.Root
	for _, b := range []byte(word) {
		if curr.children[b-'a'] == nil {
			curr.children[b-'a'] = &Node{}
			curr.children[b-'a'].children = make([]*Node, 26)
		}
		curr = curr.children[b-'a']
	}
	curr.isWord = true
}

func dfs(curr *Node, word string) bool {

	if len(word) == 0 && curr.isWord {
		return true
	} else if len(word) == 0 {
		return false
	}
	c := word[0]
	if c == '.' {
		for i := 'a'; i <= 'z'; i++ {
			if curr.children[i-'a'] != nil {
				if dfs(curr.children[i-'a'], word[1:]) {
					return true
				}
			}
		}
	} else {
		if curr.children[c-'a'] != nil {
			return dfs(curr.children[c-'a'], word[1:])
		}
	}
	return false
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(this.Root, word)
}

func main() {

	obj := Constructor()
	obj.AddWord("at")
	obj.AddWord("and")
	obj.AddWord("an")
	obj.AddWord("add")
	fmt.Println(obj.Search("a"))
	fmt.Println(obj.Search(".at"))
	obj.AddWord("bat")
	fmt.Println(obj.Search(".at"))
	fmt.Println(obj.Search("an."))
	fmt.Println(obj.Search("a.d."))
	fmt.Println(obj.Search("b."))
	fmt.Println(obj.Search("a.d"))
	fmt.Println(obj.Search("."))
}
