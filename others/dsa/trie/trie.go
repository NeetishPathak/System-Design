package main

import "strconv"

type Node struct {
	substr string
	cNodes []Node
}

func Newnode() Node {
	return Node{
		substr: "",
		cNodes: make([]Node, 26),
	}
}

func (n *Node) findWord(word string, idx int, cur string) bool {
	if idx == len(word) {
		return false
	}
	char := n.substr
	cur += string(char)
	if cur == word {
		return true
	}
	if idx > 0 {

	}
	pos, _ := strconv.Atoi(char)

	if idx == 0 || n.cNodes[-65+pos].substr == string(char) {
		idx++
		return n.cNodes[-65+pos].findWord(word, idx, cur)
	}
	return false

}

func (n *Node) addWord(word string, idx int, cur string) {
	if word == cur {
		return
	}
	char := word[idx]
	cur += string(char)
	idx++
	if n.cNodes[int(char)-65].substr == "" {
		n.cNodes[-65+int(char)] = Node{string(char), make([]Node, 26)}
	}
	n.cNodes[-65+int(char)].addWord(word, idx, cur)
}
