package parser

import "github.com/0suyog/smrkdwnp/ast"

type Delimiter struct {
	char     rune
	length   int
	position int
	canOpen  bool
	nodes    []ast.Node
}

func NewDelimiter(char rune, length int, position int, canOpen bool) Delimiter {

	d := Delimiter{
		char:     char,
		length:   length,
		position: position,
		canOpen:  canOpen,
	}
	return d
}

func (d Delimiter) Char() rune {
	return d.char
}

func (d Delimiter) Position() int {
	return d.position
}

func (d Delimiter) Length() int {
	return d.length
}

func (d Delimiter) CanOpen() bool {
	return d.canOpen
}
func (d Delimiter) CanClose() bool {
	return !d.canOpen
}

func (d Delimiter) Nodes() []ast.Node {
	return d.Nodes()
}

func (d Delimiter) PushNode(node ast.Node) {
	d.nodes = append(d.nodes, node)
}

func ArePairs(d1 Delimiter, d2 Delimiter) bool {
	if d1.char == d2.char && d1.length == d2.length && d1.canOpen == d2.CanClose() {
		return true
	}
	return false
}

type DelimiterStack struct {
	stack []Delimiter
}

func (ds DelimiterStack) IsEmpty() bool {
	if len(ds.stack) == 0 {
		return true
	}
	return false
}

func (ds DelimiterStack) Push(d Delimiter) {
	ds.stack = append(ds.stack, d)
}

func (ds DelimiterStack) Peek() (Delimiter, bool) {
	if ds.IsEmpty() {
		return Delimiter{}, false
	}
	return ds.stack[0], true
}

func (ds DelimiterStack) Pop() (Delimiter, bool) {
	if ds.IsEmpty() {
		return Delimiter{}, false
	}
	dl := ds.stack[len(ds.stack)-1]
	ds.stack = ds.stack[0 : len(ds.stack)-1]
	return dl, true
}

func (ds DelimiterStack) PushNode(node ast.Node) bool {
	if ds.IsEmpty() {
		return false
	}
	ds.stack[0].nodes = append(ds.stack[0].nodes, node)
	return true
}

func (ds DelimiterStack) PartiallyPop(count int) (Deliy, bool) {
	if ds.IsEmpty() {
		return Delimiter{}, false
	}

}

func NewDelimiterStack() DelimiterStack {
	return DelimiterStack{}
}
