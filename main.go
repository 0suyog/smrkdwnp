package main

import (
	"log"
	"os"

	"github.com/0suyog/smrkdwnp/parser"
)

func main() {
	file, err := os.Open("markdown.md")
	if err != nil {
		log.Fatal(err)
	}
	parser.Parse(file)
}
