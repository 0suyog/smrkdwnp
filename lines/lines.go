package lines

import (
	"bufio"
	"os"

	"github.com/0suyog/smrkdwnp/custom_errors"
)

type File struct {
	IncompleteSetex      bool
	IncompleteICodeBlock bool
	file                 *os.File
	sc                   *bufio.Scanner
	NotUsedLines         []*Line
}

func NewFile(file *os.File) *File {
	newSc := bufio.NewScanner(file)
	return &File{
		file: file,
		sc:   newSc,
	}
}

func (f *File) Line() (*Line, error) {
	canScan := f.sc.Scan()
	if !canScan {
		return &Line{}, custom_errors.NoNewLine
	}
	return NewLine([]rune(f.sc.Text())), nil
}

func (f *File) Unused(l *Line) {
	f.NotUsedLines = append(f.NotUsedLines, l)
}
