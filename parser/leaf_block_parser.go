package parser

import (
	"fmt"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

func ThematicBreakParser(text []rune, currentIndex *int) ast.Node {
	var char rune
	tempIndex := *currentIndex
	for ; tempIndex < *currentIndex+4; tempIndex++ {
		peekedRune, err := utils.Peek(text, tempIndex)
		if err != nil {
			return ast.NULLNODE
		}
		if peekedRune == ' ' {
			continue
		}
		if peekedRune == '*' || peekedRune == '-' || peekedRune == '_' {
			char = peekedRune
			break
		}
	}
	if char == 0 {
		fmt.Println("Didnt find any character")
		return ast.NULLNODE
	}
	// making sure atleast 3 same character are in thre
	tempIndex += 1
	length := 1
	for ; length < 3; tempIndex++ {
		peekedRune, err := utils.Peek(text, tempIndex)
		if err != nil || peekedRune == '\n' {
			fmt.Println("Finished before 3 character")
			return ast.NULLNODE
		}
		if peekedRune == ' ' || peekedRune == '\t' {
			continue
		}
		if peekedRune != char {
			fmt.Println("Anotehr character found before 3 char")
			return ast.NULLNODE
		}
		fmt.Println(string(peekedRune))
		length++
		fmt.Println(length)
	}

	// checking there are no other character in the whole line
	for ; ; tempIndex++ {
		peekedRune, err := utils.Peek(text, tempIndex)
		if err != nil || peekedRune == '\n' {
			break
		}
		if peekedRune != ' ' && peekedRune != '\t' && peekedRune != char {
			fmt.Println("Anotehr character found")
			return ast.NULLNODE
		}
	}
	*currentIndex = tempIndex
	return *ast.NewThematicBreakNode()
}

func ATXHeadingParser(text []rune, currentIndex *int) {

}
