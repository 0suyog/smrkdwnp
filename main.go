package main

import (
	"fmt"

	"github.com/0suyog/smrkdwnp/parser"
	// "github.com/0suyog/smrkdwnp/utils"
)

func main() {
	ind := 0
	fmt.Println(parser.EmphasisAndStrongParser([]rune("*foo *bar**"), &ind))
}
