package main

import (
	"os"
	"time"

	"github.com/lhdv/learn_go_with_tests/clockface"
)

func main() {

	t := time.Now()

	clockface.SVGWriter(os.Stdout, t)
}
