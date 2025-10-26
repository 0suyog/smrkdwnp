package main

import (
	"flag"
	"fmt"
	"github.com/0suyog/smrkdwnp/parser"
)

func main() {

	text := flag.String("t", "", "Markdown text to parse")
	ind := flag.Int("i", 0, "Start position of first delimiter")
	flag.Parse()

	fmt.Println(parser.EmphasisAndStrongParser([]rune(*text), ind))
}
