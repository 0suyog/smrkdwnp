package ast

import "fmt"

type LeafNode struct {
	name  string
	value string
}

type Node struct {
	name  string
	value string
	nodes []Node
}

func NewNode(name string, nodes []Node) Node {
	return Node{
		name:  name,
		nodes: nodes,
	}
}

func NewTextNode(value string) Node {
	return Node{
		name:  "TEXT",
		value: value,
		nodes: []Node{},
	}
}

func NewEmphasisNode(nodes []Node) Node {
	return Node{
		name:  "EMPHASIS",
		value: "*",
		nodes: nodes,
	}
}

func NewLeafNode(name string, value string) LeafNode {
	return LeafNode{
		name:  name,
		value: value,
	}
}

var NULLNODE = Node{
	name:  "",
	value: "",
}

func (ln LeafNode) String() string {
	return fmt.Sprintf("node_name: %s, node_value: \"%s\"", ln.name, ln.value)
}
