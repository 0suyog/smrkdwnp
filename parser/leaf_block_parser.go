package parser

import "github.com/0suyog/smrkdwnp/lines"

type Leaf_Block_Parser func(f *lines.File) (*Leaf_Block, bool)

//
// import (
// 	"github.com/0suyog/smrkdwnp/ast"
// 	"github.com/0suyog/smrkdwnp/lines"
// )
//
// type LeafBlockParser func(f *lines.File) (*ast.ASTNODE, bool)
//
// func ParseLeaf(f *lines.File) (*ast.ASTNODE, bool) {
// 	_ = []LeafBlockParser{AtxHeadingParser}
// }

//
// import (
// 	"fmt"
//
// 	"github.com/0suyog/smrkdwnp/ast"
// 	"github.com/0suyog/smrkdwnp/lines"
// 	"github.com/0suyog/smrkdwnp/utils"
// )
//
// type Parser func(l lines.Line) (ast.LeafNode, bool)
//
// func ATXHeadingParser(text []rune, currentIndex *int) (_ ast.LeafNode, success bool) {
// 	tempIndex := *currentIndex
// 	defer func() {
// 		if success {
// 			*currentIndex = tempIndex
// 		}
// 	}()
// 	var char rune
// 	fmt.Printf("First scan\n")
// 	for ; tempIndex < *currentIndex+4; tempIndex++ {
// 		fmt.Printf("tempIndex is %d", tempIndex)
// 		peekedRune, err := utils.Peek(text, tempIndex)
// 		if err != nil {
// 			return ast.NullLeafNode, success
// 		}
// 		if peekedRune == ' ' {
// 			continue
// 		}
// 		if peekedRune == '#' {
// 			char = peekedRune
// 			break
// 		}
// 	}
// 	if char != '#' {
// 		return ast.NullLeafNode, success
// 	}
// 	length := 1
// 	tempIndex++
// 	// countingnumber of hashes
// 	fmt.Printf("Scanning for number of hashes\n ")
// 	for ; ; tempIndex++ {
// 		fmt.Println(tempIndex)
// 		if length > 6 {
// 			return ast.NullLeafNode, success
// 		}
// 		peekedRune, err := utils.Peek(text, tempIndex)
// 		if err != nil {
// 			success = true
// 			return *ast.NewHeadingNode(length, []rune{}), success
// 		}
// 		if peekedRune != '#' {
// 			break
// 		}
//
// 		length++
// 	}
//
// 	// make sure there is atleast 1 space between hash and content
//
// 	fmt.Printf("Making sure there is atleast 1 space between hash and content \n ")
// 	if peekedRune, _ := utils.Peek(text, tempIndex); peekedRune != ' ' && peekedRune != '\t' {
// 		fmt.Println(tempIndex)
// 		return ast.NullLeafNode, success
// 	}
// 	tempIndex++
//
// 	fmt.Printf("Trimming starting space and tab\n")
// 	fmt.Println(tempIndex)
// 	// trim starting space and tab
// 	for ; ; tempIndex++ {
// 		fmt.Println("Trimming")
// 		peekedRune, err := utils.Peek(text, tempIndex)
//
// 		if err != nil {
// 			return ast.NullLeafNode, success
// 		}
//
// 		if peekedRune != ' ' && peekedRune != '\t' {
// 			break
// 		}
// 	}
// 	fmt.Println(tempIndex)
// 	contentStartPosition := tempIndex
// 	contentEndPosition := tempIndex
// 	isPotentialEndingSequence := false
//
// 	// get ocntent until newline or a hash is found
// 	for ; ; tempIndex++ {
// 		peekedRune, err := utils.Peek(text, tempIndex)
// 		if err != nil || peekedRune == '\n' {
// 			fmt.Println()
// 			success = true
// 			return *ast.NewHeadingNode(length, text[contentStartPosition:contentEndPosition]), success
// 		}
// 		// if found a # put it in potentiallClosingSequence
// 		if peekedRune == '#' || peekedRune == ' ' || peekedRune == '\t' {
// 			if isPotentialEndingSequence {
// 				continue
// 			}
// 			prevRune, _ := utils.PeekPrev(text, tempIndex)
// 			if (peekedRune == '#' && (prevRune == ' ' || prevRune == '\t')) || peekedRune == ' ' || peekedRune == '\t' {
// 				isPotentialEndingSequence = true
// 				continue
// 			}
// 		}
// 		if isPotentialEndingSequence {
// 			isPotentialEndingSequence = false
// 		}
// 		contentEndPosition = tempIndex + 1
// 	}
// }
//
// func SetexHeadingParser(text []rune, currentIndex *int) (_ ast.LeafNode, success bool) {
// 	tempIndex := *currentIndex
// 	defer func() {
// 		if success {
// 			*currentIndex = tempIndex
// 		}
// 	}()
// 	indentationLT3 := false
// 	// upto 3 space of indentation
// 	for ; tempIndex < *currentIndex+4; tempIndex++ {
// 		peekedRune, err := utils.Peek(text, tempIndex)
// 		if err != nil {
// 			success = true
// 			return ast.NewParagraphNode(text[*currentIndex:tempIndex]), success
// 		}
// 		if peekedRune != ' ' && peekedRune != '\t' {
// 			indentationLT3 = true
// 			break
// 		}
// 	}
// 	if !indentationLT3 {
// 		return ast.NullLeafNode, success
// 	}
// 	var content []rune
// 	isSpaceAtStart := false
// 	canBeStart := true
// 	isPotentialSpaceAtEnd := false
// 	// scan for first line
// 	for ; ; tempIndex++ {
// 		if text[tempIndex] == ' ' || text[tempIndex] == '\t' {
// 			if canBeStart {
// 				isSpaceAtStart = true
// 				canBeStart = false
// 				continue
// 			}
// 			if !isPotentialSpaceAtEnd {
// 				isPotentialSpaceAtEnd = true
// 			}
// 			continue
// 		}
// 		if text[tempIndex] == '\n' {
// 			if isSpaceAtStart || canBeStart {
// 				// this is line break so it should retuon a paragraph
// 				success = true
// 				return ast.NewParagraphNode(content), true
// 			}
// 			if isPotentialSpaceAtEnd {
// 				content = append(content, ' ')
// 				canBeStart = true
// 				isSpaceAtStart = false
// 				isPotentialSpaceAtEnd = false
// 				indentationLT3
// 				for ; tempIndex < tempIndex+4; tempIndex++ {
// 					peekedRune, err := utils.Peek(text, tempIndex)
// 					if err != nil {
// 						success = true
// 						return ast.NewParagraphNode(content), success
// 					}
// 					if peekedRune != ' ' && peekedRune != '\t' {
// 						indentationLT3 = true
// 						break
// 					}
// 				}
// 			}
// 		}
// 		if text[tempIndex] == '=' || text[tempIndex] == '-' {
//
// 		}
//
// 	}
// }
