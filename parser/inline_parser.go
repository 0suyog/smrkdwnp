package parser

import (
	"github.com/0suyog/smrkdwnp/ascii"
	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

type InlineParser func([]rune, *int) (*ast.ASTNODE, bool)

func Parse_Inline(lb *Leaf_Block) *ast.ASTNODE {
	childrenNode := []*ast.ASTNODE{}
	index := 0
	text := lb.Content
	var delimTextNode *ast.ASTNODE
	plainTextNode := ast.NewTextNode([]rune{})
	delimStack := DelimiterStack{}

	for index < len(text) {
		// escaped character
		if text[index] == '\\' {
			if nextChar, err := utils.At(text, index+1); err == nil {
				// if the text gets escaped
				if ascii.IsPunct(nextChar) {
					plainTextNode.Text = append(plainTextNode.Text, text[index+1])
					index += 2
					continue
				}
			}
			// if char isnt escapable
			plainTextNode.Text = append(plainTextNode.Text, text[index])
			index += 1
			continue
		}
		// parse code span
		if codeSpanNode, success := CodeSpanParser(text, &index); success {
			childrenNode = append(childrenNode, codeSpanNode)
			continue
		}
		// if found Em or Strong delimiter
		if text[index] == '*' || text[index] == '_' {
			delimStr := ScanTillMatchingDelim(text, text[index], &index)
			delimTextNode = ast.NewTextNode(delimStr)
			delimiter := CreateEmOrStrongDelim(text, index, delimTextNode, delimStack.top)
			delimStack.Push(delimiter)
			if len(plainTextNode.Text) > 0 {
				childrenNode = append(childrenNode, plainTextNode)
				plainTextNode = ast.NewTextNode([]rune{})
			}
			childrenNode = append(childrenNode, delimTextNode)
			continue
		}
		if text[index] == '[' {
			delimTextNode = ast.NewTextNode([]rune{'['})
			if len(plainTextNode.Text) > 0 {
				childrenNode = append(childrenNode, plainTextNode)
				plainTextNode = ast.NewTextNode([]rune{})
			}
			childrenNode = append(childrenNode, delimTextNode)
			delimiter := CreateLinkLikeDelim(delimTextNode, delimStack.top)
			delimStack.Push(delimiter)
			index += 1
			continue
		}
		if text[index] == '!' {
			if nextChar, _ := utils.At(text, index+1); nextChar == '[' {
				delimTextNode = ast.NewTextNode([]rune{'!', '['})
				if len(plainTextNode.Text) > 0 {
					childrenNode = append(childrenNode, plainTextNode)
					plainTextNode = ast.NewTextNode([]rune{})
				}
				childrenNode = append(childrenNode, delimTextNode)
				delimStack.Push(CreateLinkLikeDelim(delimTextNode, delimStack.top))
				index += 1
				continue
			}
		}
		plainTextNode.Text = append(plainTextNode.Text, text[index])
		index += 1
	}
	if len(plainTextNode.Text) > 0 {
		childrenNode = append(childrenNode, plainTextNode)
	}
	return ast.NewAstNode(lb.Type, childrenNode)
}
