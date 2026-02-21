package parser

import "github.com/0suyog/smrkdwnp/lines"

type BlockTree struct {
	head             *IntermediaryBlock
	IsFencedCodeOpen bool
	openBlocks       []*IntermediaryBlock
}

func (bt *BlockTree) ProcessLine(l *lines.Line) {
	// checking for fenced code block
	// if  bt.IsFencedCodeOpen {
	// 	if l.Indentation<3 && l.ContainsOnly([]rune)
	// }

}
