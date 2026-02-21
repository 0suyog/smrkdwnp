package utils

import (
	"errors"
	"unicode"
)

// At gives what the next item is
func At(text []rune, currentIndex int) (rune, error) {
	if currentIndex >= len(text) || currentIndex < 0 {
		return rune(0), errors.New("Current index out of range")
	}
	return text[currentIndex], nil

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

//**foo

func GetEscapedPunctuation(text []rune, index int) (rune, bool) {
	nextRune, err := At(text, index)
	if err != nil {
		return rune(0), false
	}
	if unicode.IsPunct(nextRune) {
		return nextRune, true
	}
	return rune(0), false
}

// Scan text until the given function provided returns true(mostly for handling multiple delimiters)
func ScanText(text []rune, ind *int, until func(text []rune, index int) bool) string {

	value := []rune{}
	for {
		if *ind >= len(text) || until(text, *ind) {
			return string(value)
		}
		value = append(value, text[*ind])
		*ind++
	}

}

func Upto3Indentation(text []rune, ind *int, until func(text []rune, index int) bool) (firstRune rune, ok bool) {
	for ; *ind < *ind+3; *ind++ {
		if until(text, *ind) {
			return text[*ind], true
		}
	}
	return 0, false
}
