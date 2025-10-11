package ast

import "fmt"

type LeafNode struct {
	name  string
	value string
}

func NewLeafNode(name string, value string) LeafNode {
	return LeafNode{
		name:  name,
		value: value,
	}
}

var NULLNODE = LeafNode{
	name:  "",
	value: "",
}

func (ln LeafNode) String() string {
	return fmt.Sprintf("node_name: %s, node_value: \"%s\"", ln.name, ln.value)
}
