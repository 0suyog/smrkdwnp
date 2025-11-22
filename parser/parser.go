package parser

import (
	"github.com/0suyog/smrkdwnp/ast"
)

type InlineParser func([]rune) ast.ASTNODE

type Block interface {
	ToNode() ast.ASTNODE
}

type LeafBlock struct {
	Type ast.NodeType
	Text []rune
}

var inlineParsers = []InlineParser{}
