package parser

import (
	"unicode"

	"github.com/0suyog/smrkdwnp/ast"
	"github.com/0suyog/smrkdwnp/utils"
)

type DelimiterType int

const (
	ASTERIK DelimiterType = iota
	UNDERSCORE
	OPENSQBRACKET
	IMAGEOPENER
)

var CharDelimMap = map[string]DelimiterType{
	"*":  ASTERIK,
	"_":  UNDERSCORE,
	"[":  OPENSQBRACKET,
	"![": IMAGEOPENER,
}

type Delimiter struct {
	Text     *ast.ASTNODE
	Delim    DelimiterType
	Length   int
	CanOpen  bool
	CanClose bool
	IsActive bool
	Below    *Delimiter
	Above    *Delimiter
}

func IsLeftFlankingDelimiterRun(text []rune, from int, to int) bool {
	// A left flanking delimiter run shouldn't be followed by unicode space
	nextRune, err := utils.At(text, to)
	// if the run is followed by unicode space then its not a left flanking delimiter run.
	if err != nil || unicode.IsSpace(nextRune) {
		return false
	}
	// if the run is follwoed by punctuation then it can be left flanking delimiter run if its preceeded by space or punctuation
	if unicode.IsPunct(nextRune) || unicode.IsSymbol(nextRune) {
		prevRune, err := utils.At(text, from-1)
		if err != nil || unicode.IsSpace(prevRune) || unicode.IsPunct(prevRune) || unicode.IsSymbol(prevRune) {
			return true
		}
		return false
	}
	return true
}

func IsRightFlankingDelimiterRun(text []rune, from int, to int) bool {
	prevRune, err := utils.At(text, from-1)
	// if the previous rune is space then its not RightFlankingDelimiterRun
	if err != nil || unicode.IsSpace(prevRune) {
		return false
	}
	// if its preceededd by punctuation then it should be followed by sapce or punctuation to be right flanking delimiter run
	if unicode.IsPunct(prevRune) || unicode.IsSymbol(prevRune) {
		nextRune, err := utils.At(text, to)
		if err != nil || unicode.IsSpace(nextRune) || unicode.IsPunct(nextRune) || unicode.IsSymbol(nextRune) {
			return true
		}
		return false
	}
	// if its not preceeded by punctuation then return true
	return true
}

func ScanTillMatchingDelim(text []rune, delim rune, index *int) []rune {
	matchedRune := []rune{}
	textLen := len(text)
	for ; text[*index] == delim && *index < textLen; *index++ {
		matchedRune = append(matchedRune, text[*index])
	}
	return matchedRune
}

// from is open interval, to is closed interval
func CreateEmOrStrongDelim(text []rune, from int, delimStr *ast.ASTNODE, below *Delimiter) *Delimiter {
	to := from + len(delimStr.Text)
	isLeftFlanking := IsLeftFlankingDelimiterRun(text, from, to)
	isRightFlanking := IsRightFlankingDelimiterRun(text, from, to)
	prevChar, err := utils.At(text, from-1)
	if err != nil {
		prevChar = 0
	}
	followingChar, err := utils.At(text, to)
	if err != nil {
		followingChar = 0
	}
	canOpen := PotentialOpener(text[from], isLeftFlanking, isRightFlanking, prevChar)
	canClose := PotentialCloser(text[from], isLeftFlanking, isRightFlanking, followingChar, to-from)
	return newDelimeter(
		CharDelimMap[string(text[from])],
		delimStr,
		canOpen,
		canClose,
		below,
	)
}

func CreateLinkLikeDelim(delimStr *ast.ASTNODE, below *Delimiter) *Delimiter {
	return newDelimeter(
		CharDelimMap[string(delimStr.Text)],
		delimStr,
		true,
		false,
		below,
	)
}

func newDelimeter(delim DelimiterType, text *ast.ASTNODE, canOpen bool, canClose bool, below *Delimiter) *Delimiter {
	return &Delimiter{
		Delim:    delim,
		CanOpen:  canOpen,
		CanClose: canClose,
		IsActive: true,
		Below:    below,
		Above:    nil,
		Text:     text,
		Length:   0,
	}
}

func PotentialOpener(char rune, isLeftFlanking bool, isRightFlanking bool, prevChar rune) bool {

	if char == '*' {
		if isLeftFlanking {
			return true
		}
	}

	if char == '_' {
		// fmt.Printf("Delimiter is left Flanking %v\n", d.isLeftFlanking)
		// fmt.Printf("Delimiter is right Flanking %v\n", d.isRightFlanking)
		if isLeftFlanking {
			return false
		}
		if isRightFlanking {
			// check if the preceeding character is a punctuation
			if unicode.IsPunct(prevChar) || unicode.IsSymbol(prevChar) {
				return true
			}
			return false
		}
	}
	return true
}

func PotentialCloser(char rune, isLeftFlanking bool, isRightFlanking bool, followingChor rune, lenght int) bool {
	switch char {
	case '*':
		if isRightFlanking {
			return true
		}

	case '_':

		if !isRightFlanking {
			return false
		}

		if isLeftFlanking {
			// check if the following character is a punctuation
			if unicode.IsPunct(followingChor) || unicode.IsSymbol(followingChor) {
				return true
			}
			return false
		}
	}
	return true
}

type DelimiterStack struct {
	top              *Delimiter
	bottom           *Delimiter
	stack_bottom     *Delimiter
	current_position *Delimiter
	openers_bottom   *Delimiter
}

func (ds *DelimiterStack) Push(d *Delimiter) {
	if ds.top != nil {
		ds.top.Above = d
	}
	d.Below = ds.top
}
