package parser

import (
	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func MatchFirst(l lines.Line, parsers ...func(lines.Line) (ast.Node, bool)) (ast.Node, bool) {
	for _, p := range parsers {
		if node, ok := p(l); ok {
			return node, true
		}
	}
	return ast.NULLNODE, false
}
