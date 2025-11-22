package parser

import (
	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

func CodeSpanParser(text []rune) (ast.ASTNODE, bool) {
	index := 0
	mightBeValid := false
	for i, r := range text {
		if r == '`' {
			index = i
			mightBeValid = true
			break
		}
	}
	if !mightBeValid {
		return ast.NullNode, false
	}
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
			nextChar, err := utils.Peek(text, index)
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
						break
					}
					foundCloseTag = true
					break
				}
				index++
			}

			// for i := 0; i < backTickCount; i++ {
			// 	char, err := utils.Peek(text, index+i)
			// 	if err != nil {
			// 		return ast.NullNode, false
			// 	}
			// 	if char != '`' {
			// 		break
			// 	}
			// 	// if this loop doesnt break and next character isnt "`" we found the closing tag
			// 	// for this case ` ``` `
			//
			// 	if i == backTickCount-1 {
			// 		nextCharacter, err := utils.Peek(text, index+i)
			// 		// if peek gives an index out of error then the string is finished and there is nothing to check for
			// 		if err != nil {
			// 			index += backTickCount
			// 			foundCloseTag = true
			// 			break
			// 		}
			//
			// 		if nextCharacter != '`' {
			// 			foundCloseTag = true
			// 		} else {
			// 			for j := 0; j < backTickCount; j++ {
			// 				matchedRunes = append(matchedRunes, '`')
			// 			}
			// 		}
			// 		index += backTickCount
			// 	}
			// }

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
	return *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode(matchedText)}), true
}

// func EmphasisAndStrongParser(text []rune, currentIndex *int) ast.InlineNode {
// 	ds := NewDelimiterStack()
// 	sentenceNode := ast.NewSentenceNode([]ast.InlineNode{})
// 	for {
// 		if *currentIndex >= len(text) {
// 			sentenceNode.Nodes = append(sentenceNode.Nodes, ds.ToNode()...)
// 			break
// 		}
// 		if delimiter, ok := ScanDelimiterRun(text, currentIndex); ok {
// 			fmt.Printf("scanned Delimiter %s\n", delimiter)
// 			if delimiter.isRightFlanking {
// 				fmt.Printf("rightflanking Delimiter %d\n", delimiter.length)
// 				finalNode, isFinalNode := ds.PopMatchingDelimiter(&delimiter)
// 				if isFinalNode {
// 					sentenceNode.Nodes = append(sentenceNode.Nodes, finalNode...)
//
// 					if delimiter.length > 0 {
// 						fmt.Println("Delimiter still has length left")
// 						if delimiter.CanOpen() {
// 							fmt.Println("Delimiter isnt fully closed so it is pushed to the stack again")
// 							ds.Push(&delimiter)
// 							continue
// 						}
// 					}
// 					fmt.Println("This is final node one")
// 					break
// 				}
// 				if delimiter.length <= 0 {
// 					continue
// 				}
//
// 				fmt.Println("Not a final node so continuing")
// 			}
//
// 			if !delimiter.CanOpen() {
// 				if ds.IsEmpty() {
// 					fmt.Println("Nothing happened apparantly")
// 					sentenceNode.Nodes = append(sentenceNode.Nodes, delimiter.ToNode()...)
// 				} else {
// 					fmt.Println("Nothing happened and pushed to ds")
// 					ds.PushNode(delimiter.ToNode())
// 				}
// 			}
// 			if delimiter.CanOpen() {
// 				fmt.Printf("Delimiter is a opener %d\n", delimiter.length)
// 				ds.Push(&delimiter)
// 			}
// 		}
// 		if ds.IsEmpty() {
// 			return *sentenceNode
// 		}
// 		str := utils.ScanText(text, currentIndex, func(text []rune, index int) bool { return text[index] == '*' || text[index] == '_' })
// 		fmt.Printf("Scanned Text %s\n", str)
// 		ds.PushNode([]ast.InlineNode{ast.NewTextNode(str)})
// 	}
// 	return *sentenceNode
// }
