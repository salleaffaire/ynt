package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/salleaffaire/ynt/parser"

	"github.com/salleaffaire/ynt/lexer"
)

const VERSION = "0.0.1"

const PROMPT = ">>"

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func Start(in io.Reader, out io.Writer) {

	// Can I add a comment here
	scanner := bufio.NewScanner(in)

	fmt.Printf("YNT %s REPL\n", VERSION)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		if l != nil {
			p := parser.New(l)

			document := p.ParseDocument()

			// if len(p.Errors()) != 0 {
			// 	printParserErrors(out, p.Errors())
			// 	continue
			// }

			if document != nil {
				io.WriteString(out, document.String())
				io.WriteString(out, "\n")
			}
		}
	}
}
