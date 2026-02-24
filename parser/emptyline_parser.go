package parser

import (
	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func EmptyLineParser(f *lines.File) (*Leaf_Block, bool) {
	line, err := f.Line()
	if err == nil && (len(line.Content) == 0 || line.ContainsOnly(' ')) {
		return &Leaf_Block{
			Type:    ast.FRAGMENT,
			Content: []rune{},
		}, true
	}

	return &NullI_Leaf_Block, false
}
