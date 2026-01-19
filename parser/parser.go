package parser

import (
	// "log"

	"fmt"
	"os"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func Parse_Leaf_Block(lb *Leaf_Block) *ast.ASTNODE {
	parsers := []InlineParser{CodeSpanParser, EmphasisAndStrongParser}
	counter := 0
	parentNode := ast.NewAstNode(lb.Type, []ast.ASTNODE{})
	finished := false
	notMatchedText := []rune{}
	for !finished {
		didntMatchAny := true
		for _, p := range parsers {
			if counter >= len(lb.Content) {
				finished = true
				didntMatchAny = false
				if len(notMatchedText) > 0 {
					parentNode.Children = append(parentNode.Children, *ast.NewTextNode(string(notMatchedText)))
					notMatchedText = []rune{}
				}
				break
			}
			parsedNode, success := p(lb.Content, &counter)
			if !success {
				continue
			}
			didntMatchAny = false
			if len(notMatchedText) > 0 {
				parentNode.Children = append(parentNode.Children, *ast.NewTextNode(string(notMatchedText)))
				notMatchedText = []rune{}
			}
			parentNode.Children = append(parentNode.Children, parsedNode)
		}
		if didntMatchAny {
			notMatchedText = append(notMatchedText, lb.Content[counter])
			counter += 1
		}
	}
	return parentNode
}

func Scan_Leaf_Block(f *lines.File) *Container_Block {
	bodyNode := Container_Block{ast.BODY, []*Leaf_Block{}}
	parsers := []Leaf_Block_Parser{SetexParser}
	didntMatchAny := false
	for _, p := range parsers {
		parsedNode, success := p(f)
		if !success {
			continue
		}
		didntMatchAny = false
		bodyNode.Children = append(bodyNode.Children, parsedNode)
	}
	if didntMatchAny {
		fmt.Println("Havent handled what happens if no leaf block succeed")
		os.Exit(1)
	}
	return &bodyNode
}

func Parse_Container_Block(container *Container_Block) *ast.ASTNODE {
	bodyNode := ast.NewAstNode(ast.BODY, []ast.ASTNODE{})
	for _, ln := range container.Children {
		bodyNode.Children = append(bodyNode.Children, *Parse_Leaf_Block(ln))
	}
	return bodyNode
}

func Parse(f *os.File) {
	md := lines.NewFile(f)
	body := Scan_Leaf_Block(md)
	parsed_Document := Parse_Container_Block(body)
	fmt.Println(parsed_Document)
}
