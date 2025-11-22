package parser

import (
	"fmt"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/lines"
)

func ThematicBreakParser(line lines.Line) (ast.LeafNode, bool) {
	var char rune
	tempIndex := 0
	if line.Indentation > 3 {
		return ast.NullLeafNode, false
	}
	if line.FirstRune() == '*' || line.FirstRune() == '-' || line.FirstRune() == '_' {
		char = line.FirstRune()
	}
	if char == 0 {
		fmt.Println("Didnt find any character")
		return ast.NullLeafNode, false
	}
	// making sure atleast 3 same character are in thre
	tempIndex += 1
	length := 1
	for ; length < 3; tempIndex++ {
		peekedRune, err := line.At(tempIndex)
		if err != nil {
			fmt.Println("Finished before 3 character")
			return ast.NullLeafNode, false
		}
		if peekedRune == ' ' || peekedRune == '\t' {
			continue
		}
		if peekedRune != char {
			fmt.Println("Anotehr character found before 3 char")
			return ast.NullLeafNode, false
		}
		fmt.Println(string(peekedRune))
		length++
		fmt.Println(length)
	}

	// checking there are no other character in the whole line
	for ; ; tempIndex++ {
		peekedRune, err := line.At(tempIndex)
		if err != nil || peekedRune == '\n' {
			break
		}
		if peekedRune != ' ' && peekedRune != '\t' && peekedRune != char {
			fmt.Println("Anotehr character found")
			return ast.NullLeafNode, false
		}
	}
	return ast.NewThematicBreakNode(), true
}
