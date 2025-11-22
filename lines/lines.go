package lines

import ()

type Lines struct {
	Content              []Line
	IncompleteSetex      bool
	IncompleteICodeBlock bool
}

// func (lis *Lines) Parse() ast.Node {
// 	bodyNode := ast.NewBodyNode()
// 	leafParsers := []parser.Parser{}
// 	for _, line := range lis.Content {
// 		// node, ok := MatchFirst(line, leafParsers)
// 		// // in case all parsers fail just make it a paragraph node
// 		// bodyNode.Content = append(bodyNode.Content, node)
//
// 	}
// }
