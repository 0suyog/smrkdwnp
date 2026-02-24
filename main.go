package main

import (
	"fmt"
	"log"
	"os"

	"github.com/0suyog/smrkdwnp/parser"
)

func main() {
	log.Print("\n\n************NEW LOG************\n\n")
	file, err := os.Open("markdown.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println(parser.Parse(file))
}
