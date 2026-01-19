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
	default:
		panic(fmt.Sprintf("unexpected ast.NodeType %d", nt))
	}
}

func GenerateHTML(node ASTNODE) string {
	html := ""
	if node.Type == TEXT {
		return node.Text
	}
	switch node.Type {
	case BOLD:
		html += CreateHtmlTag("strong", node.Children)
	case EMPHASIS:
		html += CreateHtmlTag("em", node.Children)
	case CODESPAN:
		html += CreateHtmlTag("code", node.Children)
	case HEADING1:
		html += CreateHtmlTag("h1", node.Children)
	case HEADING2:
		html += CreateHtmlTag("h2", node.Children)
	case HEADING3:
		html += CreateHtmlTag("h3", node.Children)
	case HEADING4:
		html += CreateHtmlTag("h4", node.Children)
	case HEADING5:
		html += CreateHtmlTag("h5", node.Children)
	case HEADING6:
		html += CreateHtmlTag("h6", node.Children)
	case THEMATICBREAK:
		html += "<hr/>"
	case PARAGRAPH:
		html += CreateHtmlTag("p", node.Children)
	case FRAGMENT:
		html += CreateHtmlTag("", node.Children)
	case BODY:
		html += CreateHtmlTag("body", node.Children)
	default:
		log.Fatalf("unexpected ast.NodeType: %s", node.Type)
	}
	return html
}

func CreateHtmlTag(tagName string, children []ASTNODE) string {
	childTags := ""
	for _, c := range children {
		childTags += GenerateHTML(c)
	}
	if tagName == "" {
		return childTags
	}
	parentTag := fmt.Sprintf("<%s>%s</%s>", tagName, childTags, tagName)
	return parentTag
}

type ASTNODE struct {
	Type     NodeType
	Text     string
	Children []ASTNODE
}

func NewAstNode(t NodeType, c []ASTNODE) *ASTNODE {
	return &ASTNODE{
		Type:     t,
		Children: c,
	}
}

func NewTextNode(t string) *ASTNODE {
	return &ASTNODE{
		Type: TEXT,
		Text: t,
	}
}

var NullNode = ASTNODE{}

func (n ASTNODE) String() string {
	if n.Type == TEXT {
		return fmt.Sprintf("[%s: \"%s\"]", TEXT, n.Text)
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

// type LeafNode struct {
// 	Name    string
// 	Content []rune
// }
//
// type ContainerNode struct {
// 	Name    string
// 	Content []LeafNode
// }
//
// type InlineNode struct {
// 	Name  string
// 	Value string
// 	Nodes []InlineNode
// }
//
// func NewNode(name string, nodes []InlineNode) InlineNode {
// 	return InlineNode{
// 		Name:  name,
// 		Nodes: nodes,
// 	}
// }
//
// func NewEmphasisNode(char rune, nodes []InlineNode) *InlineNode {
// 	return &InlineNode{
// 		Name:  "EMPHASIS",
// 		Value: string(char),
// 		Nodes: nodes,
// 	}
// }
//
// func NewStrongNode(char rune, nodes []InlineNode) *InlineNode {
// 	return &InlineNode{
// 		Name:  "STRONG",
// 		Value: string(char),
// 		Nodes: nodes,
// 	}
// }
//
// func NewSentenceNode(nodes []InlineNode) *InlineNode {
// 	return &InlineNode{
// 		Name:  "SENTENCE",
// 		Value: "meow",
// 		Nodes: nodes,
// 	}
// }
//
// func NewHeadingNode(level int, content []rune) *LeafNode {
// 	node := LeafNode{
// 		Name:    "HEADING" + strconv.Itoa(level),
// 		Content: content,
// 	}
// 	return &node
// }
// func NewLeafNode(name string, value string) LeafNode {
// 	return LeafNode{
// 		Name:    name,
// 		Content: []rune{},
// 	}
// }
//
// func NewThematicBreakNode() LeafNode {
// 	return LeafNode{
// 		Name:    "THEMATICBREAK",
// 		Content: []rune{},
// 	}
// }
//
// func NewParagraphNode(content []rune) LeafNode {
// 	return LeafNode{
// 		Name:    "PARAGRAPH",
// 		Content: content,
// 	}
// }
//
// func NewBodyNode() ContainerNode {
// 	return ContainerNode{
// 		Name: "BODY",
// 	}
// }
//
// var NullLeafNode = LeafNode{
// 	Name: "",
// }
//
// var NULLNODE = InlineNode{
// 	Name:  "",
// 	Value: "",
// }
//
// func (ln LeafNode) String() string {
// 	return fmt.Sprintf("node_name: %s, node_content: \"%s\"", ln.Name, string(ln.Content))
// }
//
// func (ln LeafNode) ToHTML() string {
// 	return "HTML"
// }
//
// func (n InlineNode) String() string {
// 	if n.Name == "TEXT" {
// 		return fmt.Sprintf("[%s: \"%s\"]", n.Name, n.Value)
// 	}
//
// 	output := fmt.Sprintf("%s: [", n.Name)
// 	for i, cn := range n.Nodes {
// 		if i > 0 {
// 			output += ", "
// 		}
// 		output += cn.String()
// 	}
// 	output += "]"
// 	return output
// }
//
// func (in InlineNode) ToHTML() string {
// 	return "InlineHTML"
// }
