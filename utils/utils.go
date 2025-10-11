package utils

import (
	"errors"
	"unicode"
)

// Peek gives what the next item is
func Peek(text []rune, currentIndex int) (rune, error) {
	if currentIndex+1 >= len(text) {
		return rune(0), errors.New("Current index out of range")
	}
	return text[currentIndex+1], nil

}

// PeekPrev gives what was before currentindex
func PeekPrev(text []rune, currentIndex int) (rune, error) {

	if currentIndex <= 0 {
		return rune(0), errors.New("Current Index cant be less or equals to 0")
	}
	return text[currentIndex-1], nil
}

// IsEscaped tells whether the rune in given index is backslash escaped
func IsEscaped(text []rune, index int) bool {

	if nextRune, _ := PeekPrev(text, index); nextRune == '\\' {
		return true
	}
	return false

}

func IsDelimiterRun(text []rune, from int, to int) bool {
	if IsEscaped(text, from) || IsEscaped(text, to) {
		return false
	}
	return true
}

func IsLeftFlankingDelimiterRun(text []rune, from int, to int) bool {
	if !IsDelimiterRun(text, from, to) {
		return false
	}
	// A left flanking delimiter run shouldn't be followed by unicode space
	nextRune, err := Peek(text, to)
	// if the run is followed by unicode space then its not a left flanking delimiter run.
	if err != nil {
		return false
	}
	if unicode.IsSpace(nextRune) {
		return false
	}
	// if the run is follwoed by punctuation then it can be left flanking delimiter run if its preceeded by space or punctuation
	if unicode.IsPunct(nextRune) {
		prevRune, err := PeekPrev(text, from)
		if err != nil || unicode.IsSpace(prevRune) || unicode.IsPunct(prevRune) {
			return true
		}
		return false
	}
	return true
}

func IsRightFlankingDelimiterRun(text []rune, from int, to int) bool {
	if !IsDelimiterRun(text, from, to) {
		return false
	}
	prevRune, err := PeekPrev(text, from)
	// if the previous rune is space then its not RightFlankingDelimiterRun
	if err != nil || unicode.IsSpace(prevRune) {
		return false
	}
	// if its preceededd by punctuation then it should be followed by sapce or punctuation to be right flanking delimiter run
	if unicode.IsPunct(prevRune) {
		nextRune, err := Peek(text, to)
		if err != nil || unicode.IsSpace(nextRune) || unicode.IsPunct(nextRune) {
			return true
		}
		return false
	}
	// if its not preceeded by punctuation then return true
	return true
}

func GetEscapedCharacter(text []rune, index int) (rune, bool) {
	nextRune, err := Peek(text, index)
	if err != nil {
		return rune(0), false
	}
	if unicode.IsPunct(nextRune) {
		return nextRune, true
	}
	return rune(0), false
}
