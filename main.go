package main

import (
	"os"

	"github.com/salleaffaire/ynt/repl"
)

var version = "0.0.1"

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
