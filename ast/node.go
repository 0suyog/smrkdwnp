package ast

import (
	"fmt"
)

type LeafNode struct {
	name  string
	value string
}

type Node struct {
	Name  string
	Value string
	Nodes []Node
}

func NewNode(name string, nodes []Node) Node {
	return Node{
		Name:  name,
		Nodes: nodes,
	}
}

func NewTextNode(value string) Node {
	return Node{
		Name:  "TEXT",
		Value: value,
		Nodes: []Node{},
	}
}

func NewEmphasisNode(char rune, nodes []Node) *Node {
	return &Node{
		Name:  "EMPHASIS",
		Value: string(char),
		Nodes: nodes,
	}
}

func NewStrongNode(char rune, nodes []Node) *Node {
	return &Node{
		Name:  "STRONG",
		Value: string(char),
		Nodes: nodes,
	}
}

func NewSentenceNode(nodes []Node) *Node {
	return &Node{
		Name:  "SENTENCE",
		Value: "meow",
		Nodes: nodes,
	}
}

func NewThematicBreakNode() *Node {
	return &Node{
		Name:  "THEMATICBREAK",
		Value: "",
		Nodes: []Node{},
	}
}

func NewLeafNode(name string, value string) LeafNode {
	return LeafNode{
		name:  name,
		value: value,
	}
}

var NULLNODE = Node{
	Name:  "",
	Value: "",
}

func (ln LeafNode) String() string {
	return fmt.Sprintf("node_name: %s, node_value: \"%s\"", ln.name, ln.value)
}

func (n Node) String() string {
	if n.Name == "TEXT" {
		return fmt.Sprintf("[%s: \"%s\"]", n.Name, n.Value)
	}

	output := fmt.Sprintf("%s: [", n.Name)
	for i, cn := range n.Nodes {
		if i > 0 {
			output += ", "
		}
		output += cn.String()
	}
	output += "]"
	return output
}
