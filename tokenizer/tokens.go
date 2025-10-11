package tokenizer

import (
	"fmt"
	"unicode/utf8"
)

type Token struct {
	name      string
	character string
}

func (t Token) Name() string {
	return t.name
}

func (t Token) Character() string {
	return t.character
}

func (t Token) Length() int {
	return utf8.RuneCountInString(t.character)
}

func (t Token) String() string {
	return fmt.Sprintf("-- TokenName: %s, TokenValue: %s--", t.name, t.character)
}

var AsteriskToken = Token{
	name:      "ASTERISK",
	character: "*",
}

var ExclamationToken = Token{
	name:      "EXCLAMATION",
	character: "!",
}

var DoublequoteToken = Token{
	name:      "DOUBLEQUOTE",
	character: "\"",
}

var HashToken = Token{
	name:      "HASH",
	character: "#",
}

var DollarToken = Token{
	name:      "DOLLAR",
	character: "$",
}

var PercentToken = Token{
	name:      "PERCENT",
	character: "%",
}

var AmpersandToken = Token{
	name:      "AMPERSAND",
	character: "&",
}

var SingleQuoteToken = Token{
	name:      "SINGLEQUOTE",
	character: "'",
}

var OpeningParenthesisToken = Token{
	name:      "OPENINGPARENTHESIS",
	character: "(",
}

var ClosingParenthesisToken = Token{
	name:      "CLOSINGPARENTHESIS",
	character: ")",
}

var PlusToken = Token{
	name:      "PLUS",
	character: "+",
}

var CommaToken = Token{
	name:      "COMMA",
	character: ",",
}

var DashToken = Token{
	name:      "DASH",
	character: "-",
}

var DotToken = Token{
	name:      "DOT",
	character: ".",
}

var ForwardSlashToken = Token{
	name:      "FORWARDSLASH",
	character: "/",
}

var ColonToken = Token{
	name:      "COLON",
	character: ":",
}

var SemiColonToken = Token{
	name:      "SEMICOLON",
	character: ";",
}

var OpeningAngleBracketToken = Token{
	name:      "OPENINGANGLEBRACKET",
	character: "<",
}

var EqualsToken = Token{
	name:      "EQUALS",
	character: "=",
}

var ClosingAngleBracketToken = Token{
	name:      "CLOSINGANGLEBRACKET",
	character: ">",
}

var QuestionMarkToken = Token{
	name:      "QUESTIONMARK",
	character: "?",
}

var CommercialAtToken = Token{
	name:      "COMMERCIALAT",
	character: "@",
}

var OpeningSquareBracketToken = Token{
	name:      "OPENINGSQUAREBRACKET",
	character: "[",
}

var BackSlashToken = Token{
	name:      "BACKSLASH",
	character: "\\",
}

var ClosingSquareBracketToken = Token{
	name:      "CLOSINGSQUAREBRACKET",
	character: "]",
}

var CaretToken = Token{
	name:      "CARET",
	character: "^",
}

var UnderscoreToken = Token{
	name:      "UNDERSCORE",
	character: "_",
}

var BacktickToken = Token{
	name:      "BACKTICK",
	character: "`",
}

var OpeningCurlyBracketToken = Token{
	name:      "OPENINGCURLYBRACKET",
	character: "{",
}

var PipeToken = Token{
	name:      "PIPE",
	character: "|",
}

var ClosingCurlyBracketToken = Token{
	name:      "CLOSINGCURLYBRACKET",
	character: "}",
}

var TildeToken = Token{
	name:      "TILDE",
	character: "~",
}

var NewLineToken = Token{
	name:      "NEWLINE",
	character: "\n",
}

var HardLineBreakToken = Token{
	name:      "HARDLINEBREAK",
	character: "\\\n",
}

// Escaped ASCII punctuations donot know whether the will be necessary.

var EscapedAsteriskToken = Token{
	name:      "ESCAPEDASTERISK",
	character: "\\*",
}

var EscapedExclamationToken = Token{
	name:      "ESCAPEDEXCLAMATION",
	character: "\\!",
}

var EscapedDoublequoteToken = Token{
	name:      "ESCAPEDDOUBLEQUOTE",
	character: "\\\"",
}

var EscapedHashToken = Token{
	name:      "ESCAPEDHASH",
	character: "\\#",
}

var EscapedDollarToken = Token{
	name:      "ESCAPEDDOLLAR",
	character: "\\$",
}

var EscapedPercentToken = Token{
	name:      "ESCAPEDPERCENT",
	character: "\\%",
}

var EscapedAmpersandToken = Token{
	name:      "ESCAPEDAMPERSAND",
	character: "\\&",
}

var EscapedSingleQuoteToken = Token{
	name:      "ESCAPEDSINGLEQUOTE",
	character: "\\'",
}

var EscapedOpeningParenthesisToken = Token{
	name:      "ESCAPEDOPENINGPARENTHESIS",
	character: "\\(",
}

var EscapedClosingParenthesisToken = Token{
	name:      "ESCAPEDCLOSINGPARENTHESIS",
	character: "\\)",
}

var EscapedPlusToken = Token{
	name:      "ESCAPEDPLUS",
	character: "\\+",
}

var EscapedCommaToken = Token{
	name:      "ESCAPEDCOMMA",
	character: "\\,",
}

var EscapedDashToken = Token{
	name:      "ESCAPEDDASH",
	character: "\\-",
}

var EscapedDotToken = Token{
	name:      "ESCAPEDDOT",
	character: "\\.",
}

var EscapedForwardSlashToken = Token{
	name:      "ESCAPEDFORWARDSLASH",
	character: "\\/",
}

var EscapedColonToken = Token{
	name:      "ESCAPEDCOLON",
	character: "\\:",
}

var EscapedSemiColonToken = Token{
	name:      "ESCAPEDSEMICOLON",
	character: "\\;",
}

var EscapedOpeningAngleBracketToken = Token{
	name:      "ESCAPEDOPENINGANGLEBRACKET",
	character: "\\<",
}

var EscapedEqualsToken = Token{
	name:      "ESCAPEDEQUALS",
	character: "\\=",
}

var EscapedClosingAngleBracketToken = Token{
	name:      "ESCAPEDCLOSINGANGLEBRACKET",
	character: "\\>",
}

var EscapedQuestionMarkToken = Token{
	name:      "ESCAPEDQUESTIONMARK",
	character: "\\?",
}

var EscapedCommercialAtToken = Token{
	name:      "ESCAPEDCOMMERCIALAT",
	character: "\\@",
}

var EscapedOpeningSquareBracketToken = Token{
	name:      "ESCAPEDOPENINGSQUAREBRACKET",
	character: "\\[",
}

var EscapedBackSlashToken = Token{
	name:      "ESCAPEDBACKSLASH",
	character: "\\\\",
}

var EscapedClosingSquareBracketToken = Token{
	name:      "ESCAPEDCLOSINGSQUAREBRACKET",
	character: "\\]",
}

var EscapedCaretToken = Token{
	name:      "ESCAPEDCARET",
	character: "\\^",
}

var EscapedUnderscoreToken = Token{
	name:      "ESCAPEDUNDERSCORE",
	character: "\\_",
}

var EscapedBacktickToken = Token{
	name:      "ESCAPEDBACKTICK",
	character: "\\`",
}

var EscapedOpeningCurlyBracketToken = Token{
	name:      "ESCAPEDOPENINGCURLYBRACKET",
	character: "\\{",
}

var EscapedPipeToken = Token{
	name:      "ESCAPEDPIPE",
	character: "\\|",
}

var EscapedClosingCurlyBracketToken = Token{
	name:      "ESCAPEDCLOSINGCURLYBRACKET",
	character: "\\}",
}

var EscapedTildeToken = Token{
	name:      "ESCAPEDTILDE",
	character: "\\~",
}
