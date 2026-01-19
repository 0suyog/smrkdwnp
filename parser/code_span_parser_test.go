package parser_test

import (
	"testing"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/parser"
)

func TestCodeSpanParser(t *testing.T) {
	get_index := func(i int) *int {
		return &i
	}
	get_index(0)
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		text         []rune
		currentIndex *int
		want         ast.ASTNODE
		want2        bool
	}{
		{
			name:         "Only Code Span",
			text:         []rune("`code span`"),
			currentIndex: get_index(0),
			want:         *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode("code span")}),
			want2:        true,
		},
		{
			name:         "Only Space",
			text:         []rune(" "),
			currentIndex: get_index(0),
			want:         ast.NullNode,
			want2:        false,
		}, {
			name:         "Code span with a accent char in text",
			text:         []rune("`` foo ` bar ``"),
			currentIndex: get_index(0),
			want:         *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode("foo ` bar")}),
			want2:        true,
		},
		{
			name:         "Code span with two accent chars only with one space around",
			text:         []rune("` `` `"),
			currentIndex: new(int),
			want:         *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode("``")}),
			want2:        true,
		},
		{
			name:         "the stripping happens if space is in both sides",
			text:         []rune("` a`"),
			currentIndex: new(int),
			want:         *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode(" a")}),
			want2:        true,
		},
		{
			name:         "No stripping occurs if the code span contains only spaces",
			text:         []rune("`  `"),
			currentIndex: new(int),
			want:         *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode("  ")}),
			want2:        true,
		}, {
			name:         "Line endings are treated like spaces:",
			text:         []rune("``foo\nbar  \nbaz``"),
			currentIndex: new(int),
			want:         *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode("foo bar   baz")}),
			want2:        true,
		},
		{
			name:         "Internal space doesnt get collapsed",
			text:         []rune("`foo   bar\n baz`"),
			currentIndex: new(int),
			want:         *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode("foo   bar  baz")}),
			want2:        true,
		}, {
			name:         "Backslash doesnt work",
			text:         []rune("`foo\\`"),
			currentIndex: new(int),
			want:         *ast.NewAstNode(ast.CODESPAN, []ast.ASTNODE{*ast.NewTextNode("foo\\")}),
			want2:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := parser.CodeSpanParser(tt.text, tt.currentIndex)
			if got.String() != tt.want.String() {
				t.Errorf("Input:%s\nCodeSpanParser() = %v, want %v", string(tt.text), got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("CodeSpanParser() = %v, want %v", got2, tt.want2)
			}
		})
	}
}
