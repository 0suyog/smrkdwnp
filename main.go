package main

import (
	"fmt"

	"github.com/0suyog/smrkdwnp/utils"
)

func main() {
	fmt.Println(utils.IsRightFlankingDelimiterRun([]rune(" _\"abc\""), 1, 1))
}
