package parser

import (
	"fmt"

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
	sentenceNode := ast.NewSentenceNode([]ast.Node{})
	for {
		if *currentIndex >= len(text) {
			sentenceNode.Nodes = append(sentenceNode.Nodes, ds.ToNode()...)
			break
		}
		if delimiter, ok := ScanDelimiterRun(text, currentIndex); ok {
			fmt.Printf("scanned Delimiter %s\n", delimiter)
			if delimiter.isRightFlanking {
				fmt.Printf("rightflanking Delimiter %d\n", delimiter.length)
				finalNode, isFinalNode := ds.PopMatchingDelimiter(&delimiter)
				if isFinalNode {
					sentenceNode.Nodes = append(sentenceNode.Nodes, finalNode...)

					if delimiter.length > 0 {
						fmt.Println("Delimiter still has length left")
						if delimiter.CanOpen() {
							fmt.Println("Delimiter isnt fully closed so it is pushed to the stack again")
							ds.Push(&delimiter)
							continue
						}
					}
					fmt.Println("This is final node one")
					break
				}
				if delimiter.length <= 0 {
					continue
				}

				fmt.Println("Not a final node so continuing")
			}

			if !delimiter.CanOpen() {
				if ds.IsEmpty() {
					fmt.Println("Nothing happened apparantly")
					sentenceNode.Nodes = append(sentenceNode.Nodes, delimiter.ToNode()...)
				} else {
					fmt.Println("Nothing happened and pushed to ds")
					ds.PushNode(delimiter.ToNode())
				}
			}
			if delimiter.CanOpen() {
				fmt.Printf("Delimiter is a opener %d\n", delimiter.length)
				ds.Push(&delimiter)
			}
		}
		if ds.IsEmpty() {
			return *sentenceNode
		}
		str := utils.ScanText(text, currentIndex, func(text []rune, index int) bool { return text[index] == '*' || text[index] == '_' })
		fmt.Printf("Scanned Text %s\n", str)
		ds.PushNode([]ast.Node{ast.NewTextNode(str)})
	}
	return *sentenceNode
}
