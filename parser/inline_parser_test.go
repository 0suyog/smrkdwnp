package parser_test

import (
	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/parser"
	"testing"
)

func TestEmphasisAndStrongParser(t *testing.T) {
	type testStruct struct {
		name string // description of this test case
		// Named input parameters for target function.
		text         []rune
		currentIndex *int
		want         ast.Node
	}
	tests := []testStruct{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parser.EmphasisAndStrongParser(tt.text, tt.currentIndex)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("EmphasisAndStrongParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
