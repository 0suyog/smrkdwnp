package parser

import "github.com/0suyog/smrkdwnp/ast"

type Leaf_Block struct {
	Type    ast.NodeType
	Content []rune
}

var NullI_Leaf_Block = Leaf_Block{
	Type:    ast.NULL,
	Content: []rune{},
}

type Container_Block struct {
	Type     ast.NodeType
	Children []*Leaf_Block
}
