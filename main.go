package main

import (
	"fmt"
	"github.com/0suyog/smrkdwnp/parser"
	// "github.com/0suyog/smrkdwnp/utils"
)

func main() {
	ind := 3
	fmt.Println(parser.EmphasisAndStrongParser([]rune("foo******bar*********baz"), &ind))
}
