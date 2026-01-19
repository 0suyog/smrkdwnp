package parser

import (
	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

func EmphasisAndStrongParser(text []rune, currentIndex *int) (ast.ASTNODE, bool) {
	index := *currentIndex
	success := true
	ds := NewDelimiterStack()
	var returnNode ast.ASTNODE
	for {
		if index >= len(text) {
			return ast.NullNode, !success
		}
		if delimiter, ok := ScanDelimiterRun(text, &index); ok {
			if delimiter.isRightFlanking {
				finalNode, isFinalNode := ds.PopMatchingDelimiter(&delimiter)
				if isFinalNode {
					returnNode = finalNode

					break
				}
				if delimiter.length <= 0 {
					continue
				}

			}

			if !delimiter.CanOpen() {
				if ds.IsEmpty() {
					break
				} else {
					ds.PushNode([]ast.ASTNODE{delimiter.ToNode()})
				}
			}
			if delimiter.CanOpen() {
				ds.Push(&delimiter)
			}
		}
		if ds.IsEmpty() {
			success = false
			return ast.NullNode, success
		}
		str := utils.ScanText(text, &index, func(text []rune, index int) bool { return text[index] == '*' || text[index] == '_' })
		ds.PushNode([]ast.ASTNODE{*ast.NewTextNode(str)})
	}

	return *&returnNode, success
}
