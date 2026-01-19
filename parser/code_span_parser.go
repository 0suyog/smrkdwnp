package parser

import (
	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

func CodeSpanParser(text []rune, currentIndex *int) (ast.ASTNODE, bool) {
	if text[*currentIndex] != '`' {
		return ast.NullNode, false
	}
	index := *currentIndex
	backTickCount := 1
	index += 1
	matchedRunes := []rune{}
	foundCloseTag := false
	isAllSpace := true
	// Counting number of backticks
	for {
		nextChar, err := utils.Peek(text, index)
		if err != nil {
			return ast.NullNode, false
		}
		if nextChar != '`' {
			break
		}
		backTickCount++
		index++
	}

	// getting text
	for {
		// replace line endign with space
		nextChar, err := utils.Peek(text, index)
		if err != nil {
			return ast.NullNode, false
		}
		if nextChar == '\n' || nextChar == '\r' {
			// if line ending is followed by another line ending then not valid codespan
			nextChar, err := utils.Peek(text, index+1)
			if err != nil {
				return ast.NullNode, false
			}
			if nextChar == '\n' || nextChar == '\r' {
				return ast.NullNode, false
			}
			matchedRunes = append(matchedRunes, ' ')
			index++
			continue
		}

		// if find a backtick and its not preceeded by a backtick then check if the number matches to opening backticks
		if nextChar == '`' && matchedRunes[len(matchedRunes)-1] != '`' {
			closingTag := "`"
			index++
			for {
				char, err := utils.Peek(text, index)
				if err != nil {
					if len(closingTag) != backTickCount {
						return ast.NullNode, false
					}
					foundCloseTag = true
					break
				}
				if char != '`' {
					if len(closingTag) != backTickCount {
						matchedRunes = append(matchedRunes, []rune(closingTag)...)
						isAllSpace = false
						break
					}
					foundCloseTag = true
					break
				}
				closingTag += "`"
				index++
			}
		}
		// break the loop if closing tag is found
		if foundCloseTag {
			break
		}
		// if nothing happened then we append the rune to matched matched runes
		if isAllSpace && text[index] != ' ' {
			isAllSpace = false
		}
		matchedRunes = append(matchedRunes, text[index])
		index++
	}

	// if there is no matched runes then return null node
	if len(matchedRunes) == 0 {
		return ast.NullNode, false
	}

	// remove 1 trailing and preceeding whitespace character if white space present
	// in both front and back if the matched string isnt all space
	matchedText := string(matchedRunes)
	if !isAllSpace {
		if matchedRunes[0] == ' ' && matchedRunes[len(matchedRunes)-1] == ' ' {
			matchedText = string(matchedRunes[1 : len(matchedRunes)-1])
		}
	}
	*currentIndex = index
	return *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode(matchedText)}), true
}
