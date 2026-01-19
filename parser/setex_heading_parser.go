package parser

import (
	"fmt"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func SetexParser(f *lines.File) (*Leaf_Block, bool) {
	line, err := f.Line()

	if err != nil || line.Indentation > 3 {
		fmt.Println("Either indentation or there is no line at all")
		return &NullI_Leaf_Block, false
	}

	var content []*lines.Line

	for {
		if err != nil {
			return &NullI_Leaf_Block, false
		}
		// fmt.Println(string(line.Content))
		if line.IsEmpty {
			return &NullI_Leaf_Block, false
		}

		if line.ContainsOnly('-', '=') && len(content) > 0 {
			var headingBlock *Leaf_Block

			if line.FirstRune == '-' {
				headingBlock = &Leaf_Block{
					Type: ast.HEADING2,
				}
			} else {
				headingBlock = &Leaf_Block{
					Type: ast.HEADING1,
				}
			}

			headingBlock.Content = lines.ConbineContent(' ', content...)
			return headingBlock, true
		}
		content = append(content, line)

		line, err = f.Line()
	}
}
