package main

import (
	"bytes"
)

type SearchEngine struct {
	root *Node
}

func NewSearchEngine(n *Node) SearchEngine {
	return SearchEngine{root: n}
}

func (s SearchEngine) FindByPrefix(p string) ([]string, error) {
	f, err := s.root.GetNodeFromPrefix(p)
	if err != nil {
		return nil, err
	}
	return build(f, p), nil
}

func build(n *Node, acc string) []string {
	var match []string
	if n.word {
		match = append(match, acc)
	}
	for k, val := range n.child {
		b := bytes.NewBufferString(acc)
		b.WriteByte(k)
		match = append(match, build(val, b.String())...)
	}
	return match
}
