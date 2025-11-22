package lines

import "fmt"

type Line struct {
	Indentation int
	IsEmpty     bool
	Content     []rune
}

func NewLine(text []rune) Line {
	// count indentation and check emptyness
	for i, r := range text {
		if r != ' ' && r != '\t' {
			return Line{
				Indentation: i,
				IsEmpty:     false,
				Content:     text,
			}
		}
	}
	return Line{
		IsEmpty: true,
		Content: text,
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

func (l *Line) FirstRune() rune {
	if l.IsEmpty {
		return 0
	}
	return l.Content[l.Indentation]
}
