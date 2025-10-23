package parser

import (
	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

func CodeSpanParser(text []rune, currentIndex *int) ast.Node {
	tempIndex := *currentIndex
	if text[tempIndex] != '`' {
		return ast.NULLNODE
	}
	backTickCount := 1
	tempIndex += 1
	matchedRunes := []rune{}
	foundCloseTag := false
	isAllSpace := true
	// Counting number of backticks
	for {
		if text[tempIndex] != '`' {
			break
		}
		backTickCount++
		tempIndex++
	}

	// getting text
	for {
		// replace line endign with space
		if text[tempIndex] == '\n' || text[tempIndex] == '\r' {
			// if line ending is followed by another line ending then not valid codespan
			nextChar, err := utils.Peek(text, tempIndex)
			if err != nil {
				panic(err.Error())
			}
			if nextChar == '\n' || nextChar == '\r' {
				return ast.NULLNODE
			}
			matchedRunes = append(matchedRunes, ' ')
			tempIndex++
			continue
		}
		// if find a backtick and its not preceeded by a backtick then check if the number matches to opening backticks
		if text[tempIndex] == '`' && matchedRunes[len(matchedRunes)-1] != '`' {

			for i := 0; i < backTickCount; i++ {
				if text[tempIndex+i] != '`' {
					break
				}
				// if this loop doesnt break and next character isnt "`" we found the closing tag
				// for this case ` ``` `

				if i == backTickCount-1 {
					nextCharacter, err := utils.Peek(text, tempIndex+i)
					// if peek gives an index out of error then the string is finished and there is nothing to check for
					if err != nil {
						tempIndex += backTickCount
						foundCloseTag = true
						break
					}

					if nextCharacter != '`' {
						foundCloseTag = true
					} else {
						for j := 0; j < backTickCount; j++ {
							matchedRunes = append(matchedRunes, '`')
						}
					}
					tempIndex += backTickCount
				}
			}

		}
		// break the loop if closing tag is found
		if foundCloseTag {
			break
		}
		// if nothing happened then we append the rune to matched matched runes
		if isAllSpace && text[tempIndex] != ' ' {
			isAllSpace = false
		}
		matchedRunes = append(matchedRunes, text[tempIndex])
		tempIndex++
	}

	// if there is no matched runes then return null node
	if len(matchedRunes) == 0 {
		return ast.NULLNODE
	}

	// remove 1 trailing and preceeding whitespace character if white space present
	// in both front and back if the matched string isnt all space
	matchedText := string(matchedRunes)
	if !isAllSpace {
		if matchedRunes[0] == ' ' && matchedRunes[len(matchedRunes)-1] == ' ' {
			matchedText = string(matchedRunes[1 : len(matchedRunes)-1])
		}
	}
	*currentIndex = tempIndex
	return ast.NewNode("CODESPAN", []ast.Node{ast.NewTextNode(matchedText)})

}

func EmphasisAndStrongParser(text []rune, currentIndex *int) ast.Node {
	ds := NewDelimiterStack()
	for {
		// if index is more than length of text then return node fo stack
		if *currentIndex >= len(text) {
			sentenceNode := ast.NewSentenceNode(ds.ToNode())
			return *sentenceNode
		}
		if delimiter, ok := ScanDelimiterRun(text, currentIndex); ok {
			// _, ok := ds.Peek()
			// if !ok && delimiter.CanOpen() {
			// 	ds.Push(&delimiter)
			// 	continue
			// }
			if ok {
				// if delimiter.CanClose(*recentOpener) {
				finalNode, ok := ds.PopMatchingDelimiter(&delimiter)
				if !ok {
					sentenceNode := ast.NewSentenceNode(finalNode)
					return *sentenceNode
				}
				continue
				// }
			}
			if delimiter.CanOpen() {
				ds.Push(&delimiter)
				continue
			}
			// if delimiter.CanOpen() && delimiter.CanClose(*recentOpener)
			// if delimiter.IsLeftFlanking() && delimiter.IsRightFlanking() {
			// 	recentOpener, ok := ds.Peek()
			// 	if !ok || !delimiter.CanClose(*recentOpener) {
			// 		ds.Push(&delimiter)
			// 		continue
			// 	}
			// }
			// if delimiter.IsRightFlanking() {
			// 	// if delimiter.CanClose(*recentOpener) {
			// 	finalNode, ok := ds.PopMatchingDelimiter(&delimiter)
			// 	if !ok {
			// 		sentenceNode := ast.NewSentenceNode(finalNode)
			// 		return *sentenceNode
			// 	}
			// 	continue
			// }
			// if delimiter.isLeftFlanking {
			// 	ds.Push(&delimiter)
			// 	continue
			// }
			if ds.IsEmpty() {
				senteceNode := ast.NewSentenceNode(delimiter.ToNode())
				return *senteceNode
			}
			ds.PushNode(delimiter.ToNode())
			continue
		}

		if ds.IsEmpty() {
			return ast.NULLNODE
		}

		str := utils.ScanText(text, currentIndex, func(text []rune, index int) bool { return text[index] == '*' || text[index] == '_' })
		ds.PushNode([]ast.Node{ast.NewTextNode(str)})
	}
}
