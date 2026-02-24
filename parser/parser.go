package parser

import (
	"io"
	"log"

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
	multilineParsers := []Leaf_Block_Parser{SetexParser, IndentedCodeBlockParser}
	parsers := []Leaf_Block_Parser{ThematicBlockParser, AtxParser}
	// log.Println("test")
	for !f.IsFinishedParsing {
		parsedBlocks := []*Leaf_Block{}
		// log.Println(string(lines.ConbineContent(' ', f.GetStack().GetStack()...)))
		makeParagraph := false
		parsedSuccessful := false
		for _, p := range multilineParsers {
			log.Println("in dangling parser")
			parsedNode, success := p(f)
			if !success {
				continue
			}
			parsedSuccessful = true
			log.Println("parsedNode: ", string(parsedNode.Content))
			parsedBlocks = append(parsedBlocks, parsedNode)
			break
		}
		if !parsedSuccessful {
			for _, p := range parsers {
				prevStackLen := f.StackLength()
				log.Println("in  parser")
				// log.Println(string(lines.ConbineContent(' ', f.GetStack().GetStack()...)))
				parsedNode, success := p(f)
				log.Println("parsing")
				if !success {
					log.Println("failed")
					continue
				}
				log.Println("passed")
				log.Println("parsed Node: ", string(parsedNode.Content))
				log.Println("prev stack len: ", prevStackLen, "stack len: ", f.StackLength())
				if f.StackLength() != prevStackLen+1 {
					makeParagraph = true
				}
				log.Println("making paragraph")
				f.ParsingSucceeded()
				parsedBlocks = append(parsedBlocks, parsedNode)
				break
			}
			if !makeParagraph {
				log.Println("checking to make paragraph")
				l, err := f.Line()
				if err != nil {
					makeParagraph = true
				} else if l.IsEmpty {
					f.ParsingSucceeded()
					makeParagraph = true
				}
			}

			if makeParagraph {
				log.Println("succeeded")
				log.Println(string(lines.CombineContent(' ', f.GetStack().GetStack()...)))
				if content := f.GetAllUnusedLinesCombined(); len(content) > 0 {
					paraNode := Leaf_Block{
						Type:    ast.PARAGRAPH,
						Content: content,
					}
					parsedBlocks = append([]*Leaf_Block{&paraNode}, parsedBlocks...)
				}
			} else {
				log.Println("failed")
			}
			f.Next()
		}
		bodyNode.Children = append(bodyNode.Children, parsedBlocks...)
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

func Parse(f io.Reader) string {
	md := lines.NewFile(f)
	body := Scan_Leaf_Block(md)
	parsed_Document := Parse_Container_Block(body)
	log.Println("parsed Doc: ", parsed_Document)
	genHtml := ast.GenerateHTML(parsed_Document)
	return genHtml
}
