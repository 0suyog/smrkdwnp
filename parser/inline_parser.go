package parser

import "github.com/0suyog/smrkdwnp/ast"

type InlineParser func([]rune, *int) (ast.ASTNODE, bool)
