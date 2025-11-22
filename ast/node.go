package ast

import (
	"fmt"
	"strconv"
)

type NodeType int

const (
	TEXT NodeType = iota
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
)

func GenerateHTML(node ASTNODE) string {
	html := ""
	if node.Type == TEXT {
		return node.Text
	}
	for _, child := range node.Children {
		switch node.Type {
		case BOLD:
			html += fmt.Sprintf("<strong>%s</strong>", child.Text)
		case EMPHASIS:
			html += fmt.Sprintf("<em>%s</em>", GenerateHTML(child))
		case CODESPAN:
			text := GenerateHTML(child)
			html += fmt.Sprintf("<code>%s</code>", text)
		case HEADING1:
			html += fmt.Sprintf("<h1>%s</h1>", GenerateHTML(child))
		case HEADING2:
			html += fmt.Sprintf("<h2>%s</h2>", GenerateHTML(child))
		case HEADING3:
			html += fmt.Sprintf("<h3>%s</h3>", GenerateHTML(child))
		case HEADING4:
			html += fmt.Sprintf("<h4>%s</h4>", GenerateHTML(child))
		case HEADING5:
			html += fmt.Sprintf("<h5>%s</h5>", GenerateHTML(child))
		case HEADING6:
			html += fmt.Sprintf("<h6>%s</h6>", GenerateHTML(child))
		case THEMATICBREAK:
			html += "<hr/>"
		default:
			panic(fmt.Sprintf("unexpected ast.NodeType: %#v", child.Type))
		}
	}
	return html
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

type Node interface {
	ToHTML() string
}

type LeafNode struct {
	Name    string
	Content []rune
}

type ContainerNode struct {
	Name    string
	Content []LeafNode
}

type InlineNode struct {
	Name  string
	Value string
	Nodes []InlineNode
}

func NewNode(name string, nodes []InlineNode) InlineNode {
	return InlineNode{
		Name:  name,
		Nodes: nodes,
	}
}

func NewEmphasisNode(char rune, nodes []InlineNode) *InlineNode {
	return &InlineNode{
		Name:  "EMPHASIS",
		Value: string(char),
		Nodes: nodes,
	}
}

func NewStrongNode(char rune, nodes []InlineNode) *InlineNode {
	return &InlineNode{
		Name:  "STRONG",
		Value: string(char),
		Nodes: nodes,
	}
}

func NewSentenceNode(nodes []InlineNode) *InlineNode {
	return &InlineNode{
		Name:  "SENTENCE",
		Value: "meow",
		Nodes: nodes,
	}
}

func NewHeadingNode(level int, content []rune) *LeafNode {
	node := LeafNode{
		Name:    "HEADING" + strconv.Itoa(level),
		Content: content,
	}
	return &node
}
func NewLeafNode(name string, value string) LeafNode {
	return LeafNode{
		Name:    name,
		Content: []rune{},
	}
}

func NewThematicBreakNode() LeafNode {
	return LeafNode{
		Name:    "THEMATICBREAK",
		Content: []rune{},
	}
}

func NewParagraphNode(content []rune) LeafNode {
	return LeafNode{
		Name:    "PARAGRAPH",
		Content: content,
	}
}

func NewBodyNode() ContainerNode {
	return ContainerNode{
		Name: "BODY",
	}
}

var NullLeafNode = LeafNode{
	Name: "",
}

var NULLNODE = InlineNode{
	Name:  "",
	Value: "",
}

func (ln LeafNode) String() string {
	return fmt.Sprintf("node_name: %s, node_content: \"%s\"", ln.Name, string(ln.Content))
}

func (ln LeafNode) ToHTML() string {
	return "HTML"
}

func (n InlineNode) String() string {
	if n.Name == "TEXT" {
		return fmt.Sprintf("[%s: \"%s\"]", n.Name, n.Value)
	}

	output := fmt.Sprintf("%s: [", n.Name)
	for i, cn := range n.Nodes {
		if i > 0 {
			output += ", "
		}
		output += cn.String()
	}
	output += "]"
	return output
}

func (in InlineNode) ToHTML() string {
	return "InlineHTML"
}
