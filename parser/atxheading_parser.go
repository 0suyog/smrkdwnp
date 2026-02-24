package parser

import (
	"log"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func AtxParser(f *lines.File) (*Leaf_Block, bool) {
	log.Println("atx")
	line, err := f.Line()
	headingBlock := Leaf_Block{}
	if err != nil || line.Indentation > 3 || line.FirstRune != '#' {
		return &NullI_Leaf_Block, false
	}
	index := line.Indentation + 1
	count := 1
	for line.Content[index] == '#' {
		if count > 6 || index >= len(line.Content) {
			return &NullI_Leaf_Block, false
		}
		index += 1
		count += 1
	}
	if line.Content[index] != ' ' {
		return &NullI_Leaf_Block, false
	}
	index += 1
	if index > len(line.Content) {
		return &NullI_Leaf_Block, false
	}
	for line.Content[index] == ' ' {
		index += 1
		if index > len(line.Content) {
			return &NullI_Leaf_Block, false
		}
	}
	content := line.Content[index:]
	switch count {
	case 1:
		headingBlock.Type = ast.HEADING1
	case 2:
		headingBlock.Type = ast.HEADING2
	case 3:
		headingBlock.Type = ast.HEADING3
	case 4:
		headingBlock.Type = ast.HEADING4
	case 5:
		headingBlock.Type = ast.HEADING5
	case 6:
		headingBlock.Type = ast.HEADING6
	default:
		return &NullI_Leaf_Block, false
	}
	headingBlock.Content = content
	return &headingBlock, true
}
