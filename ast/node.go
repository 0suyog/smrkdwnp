package ast

import (
	"fmt"
	"log"
	"strings"
)

type NodeType int

const (
	NULL NodeType = iota
	TEXT
	BOLD
	CODESPAN
	EMPHASIS
	THEMATICBREAK
	HEADING1
	HEADING2
	HEADING3
	HEADING4
	HEADING5
	HEADING6
	PARAGRAPH
	FRAGMENT
	FENCEDCODEBLOCK
	INDENTEDCODEBLOCK
	BODY
)

func (nt NodeType) String() string {
	switch nt {
	case BOLD:
		return "BOLD"
	case CODESPAN:
		return "CODESPAN"
	case EMPHASIS:
		return "EMPHASIS"
	case HEADING1:
		return "HEADING1"
	case HEADING2:
		return "HEADING2"
	case HEADING3:
		return "HEADING3"
	case HEADING4:
		return "HEADING4"
	case HEADING5:
		return "HEADING5"
	case HEADING6:
		return "HEADING6"
	case PARAGRAPH:
		return "PARAGRAPH"
	case TEXT:
		return "TEXT"
	case THEMATICBREAK:
		return "THEMATICBREAK"
	case FRAGMENT:
		return "FRAGMENT"
	case BODY:
		return "BODY"
	case NULL:
		return "NULL?"
	case FENCEDCODEBLOCK:
		return "FENCEDCODEBLOCK"
	case INDENTEDCODEBLOCK:
		return "INDENTEDCODEBLOCK"
	default:
		panic(fmt.Sprintf("unexpected ast.NodeType %d", nt))
	}
}

type HTMLTag struct {
	tags  []string
	class string
}

func GenerateHTML(node *ASTNODE) string {
	log.Println("node type: ", node.Type)
	html := ""
	if node.Type == TEXT {
		return string(node.Text)
	}
	switch node.Type {
	case BOLD:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"strong"}}, node.Children)
	case EMPHASIS:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"em"}}, node.Children)
	case CODESPAN:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"code"}}, node.Children)
	case HEADING1:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"h1"}}, node.Children)
	case HEADING2:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"h2"}}, node.Children)
	case HEADING3:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"h3"}}, node.Children)
	case HEADING4:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"h4"}}, node.Children)
	case HEADING5:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"h5"}}, node.Children)
	case HEADING6:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"h6"}}, node.Children)
	case THEMATICBREAK:
		html += "<hr />"
	case PARAGRAPH:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"p"}}, node.Children)
	case FRAGMENT:
		html += CreateHtmlTag(&HTMLTag{tags: []string{}}, node.Children)
	case BODY:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"body"}}, node.Children)
	case FENCEDCODEBLOCK:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"pre", "code"}}, node.Children)
	case INDENTEDCODEBLOCK:
		html += CreateHtmlTag(&HTMLTag{tags: []string{"pre", "code"}}, node.Children)
	default:
		log.Fatalf("unexpected ast.NodeType: %s", node.Type)
	}
	return html
}

func CreateHtmlTag(htmlTag *HTMLTag, children []*ASTNODE) string {
	log.Println("creating html tag for", htmlTag)
	childTags := ""
	for _, c := range children {
		childTags += GenerateHTML(c)
	}
	if len(htmlTag.tags) == 0 {
		return childTags
	}
	opener, closer := "", ""
	for _, t := range htmlTag.tags {
		opener = fmt.Sprintf("%s<%s>", opener, t)
		closer = fmt.Sprintf("</%s>%s", t, closer)
	}
	log.Println("opener: ", opener, "closer: ", closer)
	parentTag := fmt.Sprintf("%s%s%s", opener, childTags, closer)
	log.Println("parent tag", parentTag)
	return parentTag
}

// func MultiTagBlock(children []*ASTNODE, tags []*string) {
// 	opener, closer := "", ""
// 	for _, t := range tags{
// 		opener += ""
// 	}
// }

type ASTNODE struct {
	Type     NodeType
	Text     []rune
	Children []*ASTNODE
}

func NewAstNode(t NodeType, c []*ASTNODE) *ASTNODE {
	return &ASTNODE{
		Type:     t,
		Children: c,
	}
}

func NewTextNode(t []rune) *ASTNODE {
	return &ASTNODE{
		Type: TEXT,
		Text: t,
	}
}

var NullNode = ASTNODE{}

func (n ASTNODE) String() string {
	if n.Type == TEXT {
		return fmt.Sprintf("[%s: \"%s\"]", TEXT, string(n.Text))
	}

	var output strings.Builder
	fmt.Fprintf(&output, "%s: [", n.Type)
	for i, ch := range n.Children {
		if i > 0 {
			output.WriteString(", ")
		}
		output.WriteString(ch.String())
	}
	output.WriteString("]")
	return output.String()
}
