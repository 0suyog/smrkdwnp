package parser

import (
	"fmt"
	"unicode"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

type Delimiter struct {
	char            rune
	text            []rune
	position        int
	length          int
	isLeftFlanking  bool
	isRightFlanking bool
	nodes           []ast.Node
}

var EmptyDelimiter = Delimiter{}

func NewDelimiter(char rune, length int, isLeftFlanking bool, isRightFlanking bool, text []rune, position int) Delimiter {

	d := Delimiter{
		char:            char,
		text:            text,
		length:          length,
		isLeftFlanking:  isLeftFlanking,
		isRightFlanking: isRightFlanking,
		position:        position,
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
	fmt.Println("Testing can close")
	fmt.Printf("The opener is %s", d1)
	fmt.Printf("The closer is %s", d)
	if d.char != d1.char {
		fmt.Println("Apparantly char arent same")
		return false
	}

	switch d.char {
	case '*':

		// check if the sum of their lengh is going to be a multiple of three provided delimiter is both left and right flanking
		if (d.IsLeftFlanking() && d.IsRightFlanking()) || (d1.IsLeftFlanking() && d1.isRightFlanking) {
			if (d1.length%3 != 0 || d.length%3 != 0) && (d1.length+d.length)%3 == 0 {
				fmt.Println("Sum of three goit ")
				return false
			}
			return true
		}

		if !d.IsRightFlanking() {
			fmt.Println("It isnt right flanking")
			return false
		}

		if !d.IsLeftFlanking() {
			return true
		}

	case '_':

		if !d.isLeftFlanking {
			return false
		}

		if d.isRightFlanking {
			// check if the following character is a punctuation
			if followingChar, err := utils.Peek(d.text, d.position+d.length); err != nil {
				if unicode.IsPunct(followingChar) || unicode.IsSymbol(followingChar) {
					return true
				}
			}
			return false
		}
	}
	return false
}

func (d Delimiter) CanOpen() bool {

	if d.char == '*' {
		if d.isLeftFlanking {
			return true
		}
	}

	if d.char == '_' {
		if d.isRightFlanking {
			// check if the preceeding character is a punctuation
			if preceedingChar, err := utils.PeekPrev(d.text, d.position+d.length); err != nil {
				if unicode.IsPunct(preceedingChar) || unicode.IsSymbol(preceedingChar) {
					return true
				}
			}
			return false
		}
	}

	return false

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

func (d Delimiter) String() string {
	return fmt.Sprintf("char: %s\nposition: %d\n length: %d\n isLeftFlanking: %v\n isRightFlanking: %v\n nodes: %s\n", string(d.char), d.position, d.length, d.isLeftFlanking, d.isRightFlanking, d.nodes)
}

// should give a delimiter that can close the top delimiter returns final Node if its the final node, leftOver Closing Delimiter,
func (ds *DelimiterStack) PopMatchingDelimiter(closer *Delimiter) (returnNodes []ast.Node, isFinalNode bool) {
	// need to chek the wnhole stack and close any first one that can be closed,
	index := len(ds.stack) - 1
	for closer.length > 0 {
		fmt.Printf("Index: %d, lengtho if stack : %d\n", index, len(ds.stack))
		opener, ok := ds.PeekAt(index)
		if !ok {
			if closer.CanOpen() {
				fmt.Println("this rang")
				fmt.Println(closer.length)
				return returnNodes, false
			}
			fmt.Println("This is returns final node")
			node := closer.ToTextNode()
			return node, true
		}
		fmt.Printf("Closer is %s \n", closer)
		fmt.Printf("Opener is %s\n", opener)
		if !closer.CanClose(*opener) {
			fmt.Printf("Closer cnat close Opener")
			index--
			continue
		}
		fmt.Println("DCloser can close opener")
		if closer.Length() < opener.Length() {
			fmt.Println("Closer is shorter than opener")
			// lets change the openers length to be equals to closers length, and make a new node whose length is leftover of the original opener
			newDelimiter := NewDelimiter(
				opener.char,
				opener.length-closer.length,
				opener.isLeftFlanking,
				opener.isRightFlanking,
				opener.text,
				opener.position,
			)
			fmt.Printf("Stack is %s\n", ds.stack)
			opener.length = closer.length
			closer.length = 0
			newDelimiter.nodes = ds.ToNodeUpto(index)
			fmt.Println("dog")
			fmt.Println(newDelimiter.nodes)
			ds.Push(&newDelimiter)
			fmt.Printf("The delimiterstack is %s\n", ds.stack)
			continue
		}

		// if closer.Length is more tha opener.Length then we just pop the stack and turn it into node
		// check if the stack is empty if its empty then we return the node and false, else change the length of closer delimiter
		fmt.Println("Opener is shorter or equal than cloeser")
		closer.length -= opener.length
		nodes := ds.ToNodeUpto(index)
		if index == 0 {
			fmt.Println("Idnex got to 0")
			return nodes, true
		}
		ds.PushNode(nodes)
		index--
	}
	fmt.Println("This shouldnt react logically")
	return returnNodes, false
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

func (ds *DelimiterStack) PeekAt(ind int) (*Delimiter, bool) {
	if ds.IsEmpty() || ind < 0 || ind >= len(ds.stack) {
		return &EmptyDelimiter, false
	}
	return ds.stack[ind], true
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

func ScanDelimiterRun(text []rune, index *int) (Delimiter, bool) {
	if text[*index] != '*' && text[*index] != '_' {
		return EmptyDelimiter, false
	}
	char := text[*index]
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
	isLeftFlanking := utils.IsLeftFlankingDelimiterRun(text, start, *index)
	// fmt.Printf("index: %d canOpen: %v\n", *index, canOpen)
	isRightFlanking := utils.IsRightFlankingDelimiterRun(text, start, *index)
	// fmt.Printf("index: %d canClose: %v\n", *index, canClose)

	return NewDelimiter(char, length, isLeftFlanking, isRightFlanking, text, start), true
}

func (ds *DelimiterStack) ToNode() []ast.Node {
	arr := []ast.Node{}
	for _, d := range ds.stack {
		arr = append(arr, d.ToTextNode()...)
	}
	ds.stack = []*Delimiter{}
	return arr
}

func (ds *DelimiterStack) ToNodeUpto(ind int) []ast.Node {
	arr := []ast.Node{}
	fmt.Println("Turning stack into node")
	fmt.Printf("Indes: %d\n", ind)
	fmt.Printf("The stack is %s\n", ds.stack)
	fmt.Printf("Length of stack: %d\n", len(ds.stack))
	peekedDelimiter, ok := ds.PeekAt(ind)
	if !ok {
		panic("Invalid index provided")
	}
	if ind > 0 {
		for _, d := range ds.stack[ind+1:] {
			fmt.Printf("Delimiter Length %d\n", d.length)
			arr = append(arr, d.ToTextNode()...)
			fmt.Printf("Array is %s\n", arr)
		}
	}
	peekedDelimiter.nodes = append(peekedDelimiter.nodes, arr...)
	ds.stack = ds.stack[:ind]
	fmt.Printf("Length of stack: %d\n", len(ds.stack))
	return peekedDelimiter.ToNode()
}
