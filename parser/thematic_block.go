package parser

import (
	"log"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func ThematicBlockParser(f *lines.File) (*Leaf_Block, bool) {
	log.Println("thematic")
	line, err := f.Line()

	if err != nil || line.Indentation > 3 || !line.ContainsOnlyWSpace('*', '-', '_') {
		log.Println("failed, line: ", string(line.Content))
		return &NullI_Leaf_Block, false
	}

	delimiter := line.FirstRune
	count := 0
	for _, r := range line.Content {
		if r != delimiter && r != ' ' {
			return &NullI_Leaf_Block, false
		}
		count++
	}
	if count < 3 {
		log.Println("failed, line: ", string(line.Content))
		log.Println("count of delim: ", count)
		return &NullI_Leaf_Block, false
	}
	return &Leaf_Block{
		Type: ast.THEMATICBREAK,
	}, true
}
