package lines

import (
	"bufio"
	"io"
	"log"

	"github.com/0suyog/smrkdwnp/custom_errors"
)

type File struct {
	sc                 *bufio.Scanner
	LatestLine         *Line
	unusedLineStack    *lineStack
	resend             bool
	IsFinishedScanning bool
	IsFinishedParsing  bool
}

func NewFile(file io.Reader) *File {
	newSc := bufio.NewScanner(file)
	return &File{
		sc:              newSc,
		unusedLineStack: NewLineStack(),
	}
}

func (f *File) GetStack() *lineStack {
	return f.unusedLineStack
}

func (f *File) IsStackEmpty() bool {
	return len(f.unusedLineStack.stack) == 0
}

func (f *File) StackLength() int {
	return len(f.unusedLineStack.stack)
}

func (f *File) Next() {
	f.resend = false
}

func (f *File) Line() (*Line, error) {

	if f.resend {
		log.Println("resending")
		log.Println(string(CombineContent(',', f.unusedLineStack.stack[len(f.unusedLineStack.stack)-1])))
		return f.unusedLineStack.stack[len(f.unusedLineStack.stack)-1], nil
	}
	if f.IsFinishedScanning {
		f.IsFinishedParsing = true
		return &Line{}, custom_errors.NoNewLine
	}
	canScan := f.sc.Scan()
	if !canScan {
		f.IsFinishedScanning = true
	}
	log.Println("new sent")
	retLine := NewLine([]rune(f.sc.Text()))
	f.unusedLineStack.add(retLine)
	f.resend = true
	log.Println(string(retLine.Content))
	return retLine, nil
}

func (f *File) ParsingSucceeded() {
	f.resend = false
	f.unusedLineStack.pop()
}

func (f *File) GetAllUnusedLinesCombined() []rune {
	// log.Println("stack: ", string(ConbineContent(' ', f.GetStack().stack...)))
	log.Println("get all stack")
	retRune := CombineContent(' ', f.unusedLineStack.getAll()...)
	log.Println("returned rune: ", string(retRune))
	return retRune
}

type lineStack struct {
	stack   []*Line
	top     int
	pointer int
}

func NewLineStack() *lineStack {
	return &lineStack{
		top:     -1,
		pointer: -1,
	}
}

func (s *lineStack) add(l *Line) {
	log.Println("before adding: ", string(CombineContent(',', s.stack...)))
	s.stack = append(s.stack, l)
	s.top++
	s.pointer = s.top
	log.Println("added: ", string(l.Content))
	log.Println("after adding: ", string(CombineContent(',', s.stack...)))
}

func (s *lineStack) get() (*Line, bool) {
	if s.pointer < 0 {
		return &Line{}, true
	}
	s.pointer--
	return s.stack[s.pointer+1], false
}

func (s *lineStack) pop() {
	log.Println("in pop")
	log.Println("before popping", string(CombineContent(',', s.stack...)))
	s.stack = s.stack[0 : len(s.stack)-1]
	log.Println("after popping", string(CombineContent(',', s.stack...)))
	s.top--
}

func (s *lineStack) getAll() []*Line {
	log.Println("in get all")
	log.Println("stack: ", string(CombineContent(',', s.stack...)))
	ret := s.stack
	s.stack = []*Line{}
	s.top = -1
	s.pointer = s.top
	return ret
}

func (s *lineStack) resetStack() {
	s.stack = []*Line{}
	s.top = -1
	s.pointer = -1
}

func (s *lineStack) GetStack() []*Line {
	return s.stack
}
