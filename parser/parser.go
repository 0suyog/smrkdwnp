package parser

import (
	// "log"

	"fmt"
	"os"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func Parse_Leaf_Block(lb *Leaf_Block) *ast.ASTNODE {

	return Parse_Inline(lb)

	// parsers := []InlineParser{CodeSpanParser}
	// counter := 0
	// parentNode := ast.NewAstNode(lb.Type, []*ast.ASTNODE{})
	// finished := false
	// notMatchedText := []rune{}
	// for !finished {
	// 	didntMatchAny := true
	// 	for _, p := range parsers {
	// 		if counter >= len(lb.Content) {
	// 			finished = true
	// 			didntMatchAny = false
	// 			if len(notMatchedText) > 0 {
	// 				parentNode.Children = append(parentNode.Children, ast.NewTextNode(notMatchedText))
	// 				notMatchedText = []rune{}
	// 			}
	// 			break
	// 		}
	// 		parsedNode, success := p(lb.Content, &counter)
	// 		if !success {
	// 			continue
	// 		}
	// 		didntMatchAny = false
	// 		if len(notMatchedText) > 0 {
	// 			parentNode.Children = append(parentNode.Children, ast.NewTextNode(notMatchedText))
	// 			notMatchedText = []rune{}
	// 		}
	// 		parentNode.Children = append(parentNode.Children, parsedNode)
	// 	}
	// 	if didntMatchAny {
	// 		notMatchedText = append(notMatchedText, lb.Content[counter])
	// 		counter += 1
	// 	}
	// }
	// return parentNode
}

func Scan_Leaf_Block(f *lines.File) *Container_Block {
	bodyNode := Container_Block{ast.BODY, []*Leaf_Block{}}
	parsers := []Leaf_Block_Parser{AtxParser, SetexParser}
	for !f.IsFinishedParsing {
		didntMatchAny := true
		for _, p := range parsers {
			parsedNode, success := p(f)
			if !success {
				f.ParsingFailed()
				continue
			}
			f.ParsingSucceeded()
			didntMatchAny = false
			bodyNode.Children = append(bodyNode.Children, parsedNode)
			break
		}
		if didntMatchAny {
			if content := f.GetAllUnusedLinesCombined(); len(content) > 0 {
				paraNode := Leaf_Block{
					Type:    ast.PARAGRAPH,
					Content: content,
				}
				f.ParsingSucceeded()
				bodyNode.Children = append(bodyNode.Children, &paraNode)
			}
		}
	}
	return &bodyNode
}

func Parse_Container_Block(container *Container_Block) *ast.ASTNODE {
	bodyNode := ast.NewAstNode(ast.BODY, []*ast.ASTNODE{})
	for _, ln := range container.Children {
		parsedLb := Parse_Leaf_Block(ln)
		bodyNode.Children = append(bodyNode.Children, parsedLb)
	}
	return bodyNode
}

func Parse(f *os.File) {
	md := lines.NewFile(f)
	body := Scan_Leaf_Block(md)
	parsed_Document := Parse_Container_Block(body)
	genHtml := ast.GenerateHTML(parsed_Document)
	fmt.Println(genHtml)
}
