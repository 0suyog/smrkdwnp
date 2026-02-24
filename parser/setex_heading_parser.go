package parser

import (
	"log"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func SetexParser(f *lines.File) (*Leaf_Block, bool) {
	log.Println("setex")
	line, err := f.Line()
	if f.StackLength() < 2 {
		return &NullI_Leaf_Block, false
	}
	if err != nil || line.Indentation > 3 {
		// log.Println("failed cuz nil")
		return &NullI_Leaf_Block, false
	}
	if line.ContainsOnly('-') {
		f.ParsingSucceeded()
		content := f.GetAllUnusedLinesCombined()
		log.Println("this succeeded the - one")
		return &Leaf_Block{
			Type:    ast.HEADING2,
			Content: content,
		}, true

	}

	if line.ContainsOnly('=') {
		// log.Println("meow")
		f.ParsingSucceeded()
		log.Println("this succeeded the = one")
		content := f.GetAllUnusedLinesCombined()
		return &Leaf_Block{
			Type:    ast.HEADING1,
			Content: content,
		}, true
	}
	// log.Println("failed cuz not = or -")
	// log.Println(string(line.Content))
	return &NullI_Leaf_Block, false
}
