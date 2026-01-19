package lines

import (
	"fmt"
	"slices"
	"unicode"
)

type Line struct {
	Indentation int
	IsEmpty     bool
	Content     []rune
	FirstRune   rune
}

func NewLine(text []rune) *Line {
	// count indentation and check emptyness
	for i, r := range text {
		if r != ' ' && r != '\t' {
			return &Line{
				Indentation: i,
				IsEmpty:     false,
				Content:     text,
				FirstRune:   r,
			}
		}
	}

	return &Line{
		IsEmpty:   true,
		Content:   text,
		FirstRune: 0,
	}
}

func (l *Line) At(ind int) (rune, error) {
	if ind >= len(l.Content) {
		return 0, fmt.Errorf("Index more than array size ind: %d len: %d", ind, len(l.Content))
	}

	if ind < 0 {
		return 0, fmt.Errorf("Index less than 0 ind: %d", ind)
	}
	return l.Content[ind], nil

}

func (l *Line) StartsWith(r ...rune) bool {
	return slices.Contains(r, l.FirstRune)
}

func (l *Line) ContainsOnly(r ...rune) bool {
	if !l.StartsWith(r...) {
		return false
	}
	for _, ru := range l.Content {
		if unicode.IsSpace(ru) {
			continue
		}
		if !slices.Contains(r, ru) {
			return false
		}
	}
	return true
}

func ConbineContent(seperator_delimiter rune, l ...*Line) []rune {
	content := []rune{}
	for _, li := range l {
		content = append(content, li.Content...)
		content = append(content, seperator_delimiter)
	}
	return content
}
