package main

import (
	"errors"
	"strings"
)

type Node struct {
	child map[byte]*Node
	word  bool
}

func NewNode(w bool) *Node {
	return &Node{word: w, child: make(map[byte]*Node)}
}

func (n *Node) isWord() bool {
	return n.word
}

func (n *Node) AddWord(w string) {
	addWord(n, strings.ToLower(w), 0)
}

func addWord(n *Node, w string, i int) {
	if len(w) != 0 {
		n := n.getAndPutIfAbsent(w[i], (i+1 == len(w)))
		i++
		if i != len(w) {
			addWord(n, w, i)
		}
	}
}

func (n *Node) GetNodeFromPrefix(p string) (*Node, error) {
	return recursiveDeepIteration(n, p, 0)
}

func recursiveDeepIteration(n *Node, p string, i int) (*Node, error) {
	if val, ok := n.child[p[i]]; ok {
		i++ //srly google
		if len(p) == i {
			return val, nil
		}
		return recursiveDeepIteration(val, p, i)
	} else {
		return &Node{}, errors.New("node not found")
	}
}

func (n *Node) getAndPutIfAbsent(k byte, w bool) *Node {
	if val, ok := n.child[k]; ok {
		return val
	} else {
		n.child[k] = NewNode(w)
		return n.child[k]
	}
}
