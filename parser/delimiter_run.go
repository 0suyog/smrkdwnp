package parser

import (
	"fmt"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

type Delimiter struct {
	char            rune
	length          int
	isLeftFlanking  bool
	isRightFlanking bool
	nodes           []ast.Node
}

var EmptyDelimiter = Delimiter{}

func NewDelimiter(char rune, length int, canOpen bool, canClose bool) Delimiter {

	d := Delimiter{
		char:            char,
		length:          length,
		isLeftFlanking:  canOpen,
		isRightFlanking: canClose,
	}
	return d
}

func (d *Delimiter) Char() rune {
	return d.char
}

func (d *Delimiter) Length() int {
	return d.length
}

func (d *Delimiter) IsLeftFlanking() bool {
	return d.isLeftFlanking
}
func (d *Delimiter) IsRightFlanking() bool {
	return d.isRightFlanking
}

func (d *Delimiter) Nodes() []ast.Node {
	return d.nodes
}

func (d *Delimiter) PushNode(node []ast.Node) {
	d.nodes = append(d.nodes, node...)
}

// a delimiter that can both open and close cannot form emphasis if the sum of the lengths of the delimiter runs
// containing the opening and closing delimiters is a multiple of 3 unless both lengths are multiples of 3.
// see https://spec.commonmark.org/0.31.2/#example-411
func (d *Delimiter) CanClose(d1 Delimiter) bool {

	if d.char != d1.char {
		return false
	}

	if !d.IsRightFlanking() {
		return false
	}

	if !d.IsLeftFlanking() {
		return true
	}

	// check if the sum of their lengh is going to be a multiple of three provided closing delimiter is both left and right flanking

	if (d1.length%3 != 0 || d.length%3 != 0) && (d1.length+d.length)%3 == 0 {
		return false
	}

	return true
}

func (d Delimiter) ToNode() []ast.Node {
	if d.isLeftFlanking {
		numberOfStrong := d.length / 2
		numberOfEmp := d.length % 2
		var node *ast.Node
		for numberOfStrong > 0 {
			numberOfStrong--
			if node == nil {
				node = ast.NewStrongNode(d.char, d.nodes)
				continue
			}
			newStrongNode := ast.NewStrongNode(d.char, []ast.Node{*node})
			node = newStrongNode
		}
		if numberOfEmp == 1 {
			if node == nil {
				node = ast.NewEmphasisNode(d.char, d.nodes)
			} else {
				node = ast.NewEmphasisNode(d.char, []ast.Node{*node})
			}
		}
		return []ast.Node{*node}
	}
	return d.ToTextNode()
}

func (d Delimiter) ToTextNode() []ast.Node {
	delimiterString := ""
	for d.length > 0 {
		delimiterString += string(d.char)
		d.length--
	}
	return append([]ast.Node{ast.NewTextNode(delimiterString)}, d.nodes...)
}

// should give a delimiter that can close the top delimiter returns final Node if its the final node, leftOver Closing Delimiter,
func (ds *DelimiterStack) PopMatchingDelimiter(closer *Delimiter) ([]ast.Node, bool) {
	// ook for this lets make a loop that checks whether the closer delimiter is empty, and runs till it isnt empty
	// if the closer is empty then return
	returnNodes := []ast.Node{}
	for closer.length > 0 {

		opener, ok := ds.Peek()

		if !ok {
			node := closer.ToNode()
			return node, false
		}
		if !closer.CanClose(*opener) {
			ds.PushNode(closer.ToNode())
			break
		}

		if closer.Length() < opener.Length() {
			// if closer delimiter runs length is less than opener delimiter run then we create a new delimiter that will close the closer delimiter
			// the new delimiter will have properties of opener delimiter cuz its going to be the new opener delimiter for the closer one
			// we will turn the matched delimiter (newly created opener one) into nodes and then push those nodes to the node that is on top of the
			// stack ie opener
			matchedDelimiter := NewDelimiter(opener.char, opener.length-closer.length, opener.isLeftFlanking, opener.isRightFlanking)
			matchedDelimiter.nodes = opener.nodes
			opener.length = opener.length - matchedDelimiter.length
			opener.nodes = matchedDelimiter.ToNode()
			continue
		}

		// if closer.Length is more tha opener.Length then we just pop the stack and turn it into node
		// check if the stack is empty if its empty then we return the node and false, else
		// the closer delimiter will be mutated

		opener, _ = ds.Pop()
		if ds.IsEmpty() {
			return opener.ToNode(), false
		}
		closer.length -= opener.length
		ds.PushNode(opener.ToNode())
	}
	return returnNodes, true
}

type DelimiterStack struct {
	stack []*Delimiter
}

func (ds *DelimiterStack) IsEmpty() bool {
	if len(ds.stack) == 0 {
		return true
	}
	return false
}

func (ds *DelimiterStack) Push(d *Delimiter) {
	ds.stack = append(ds.stack, d)
}

func (ds *DelimiterStack) Peek() (*Delimiter, bool) {
	if ds.IsEmpty() {
		return &EmptyDelimiter, false
	}
	return ds.stack[len(ds.stack)-1], true
}

func (ds *DelimiterStack) Pop() (*Delimiter, bool) {
	if ds.IsEmpty() {
		return &Delimiter{}, false
	}
	dl := ds.stack[len(ds.stack)-1]
	ds.stack = ds.stack[0 : len(ds.stack)-1]
	return dl, true
}

func (ds *DelimiterStack) PushNode(node []ast.Node) bool {
	if ds.IsEmpty() {
		return false
	}
	ds.stack[len(ds.stack)-1].nodes = append(ds.stack[len(ds.stack)-1].nodes, node...)
	return true
}

func NewDelimiterStack() DelimiterStack {
	return DelimiterStack{}
}

func ScanDelimiterRun(text []rune, char rune, index *int) (Delimiter, bool) {
	if text[*index] != char {
		return EmptyDelimiter, false
	}
	length := 1
	*index++
	start := *index - 1
	for {
		if *index >= len(text) || text[*index] != char {
			break
		}
		*index++
		length++
	}
	canOpen := utils.IsLeftFlankingDelimiterRun(text, start, *index)
	fmt.Printf("index: %d canOpen: %v\n", *index, canOpen)
	canClose := utils.IsRightFlankingDelimiterRun(text, start, *index)
	fmt.Printf("index: %d canClose: %v\n", *index, canClose)

	return NewDelimiter(char, length, canOpen, canClose), true
}

func (ds *DelimiterStack) ToNode() []ast.Node {
	arr := []ast.Node{}
	for _, d := range ds.stack {
		arr = append(arr, d.ToTextNode()...)
	}
	return arr
}
