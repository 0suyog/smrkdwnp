package parser

import (
	"log"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func IndentedCodeBlockParser(f *lines.File) (*Leaf_Block, bool) {
	log.Println("IndentedCodeBlock")
	if f.StackLength() > 1 {
		log.Println("Failed cus stack isnt empty")
		return &NullI_Leaf_Block, false
	}
	line, err := f.Line()
	if err != nil || line.Indentation < 4 {
		log.Println("Failed cuz err isn't equal to nil or indentation: ", line.Indentation)
		return &NullI_Leaf_Block, false
	}
	content := []*lines.Line{}
	emptyLines := []*lines.Line{}
	for line.Indentation >= 4 && err == nil {
		actualContent := line.Content[4:]
		log.Println("line content: ", string(line.Content), "actual content: ", string(actualContent))
		if line.IsEmpty {
			emptyLines = append(emptyLines, lines.NewLine(actualContent))
		} else {
			if len(emptyLines) > 0 {
				content = append(content, emptyLines...)
				emptyLines = []*lines.Line{}
			}
			content = append(content, lines.NewLine(actualContent))
		}
		f.ParsingSucceeded()
		// f.Next()
		line, err = f.Line()
	}
	return &Leaf_Block{
		Type:    ast.INDENTEDCODEBLOCK,
		Content: lines.CombineContent('\n', content...),
	}, true
}
