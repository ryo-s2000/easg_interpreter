package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

const PROMTP = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMTP)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "" {
			continue
		}
		if line == "exit" {
			return
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const ERROR_AA = `
　 　　　 　　 　＿＿＿_　　　　　　＿ ＿ ＿ ＿ ＿ ＿ ＿ ＿ ＿ ＿ ＿ ＿ ＿ ＿
　　　　　　 ／_ノ　　ヽ､_＼　 　　　　
.　　　　　／　（● ）　(● ）＼　　　Woops! something wrong...
　　　　／//////（__人__)///＼　　 ◯
　　　　|　　 　　　　　　　　　 　|　。O　　￣ ￣ ￣ ￣ ￣ ￣ ￣ ￣ ￣ ￣ ￣
　　　　 ＼　　　　 　　　　　 ／
　　　　ノ　　　　　　　　　　 　＼
　 ／´　　　　　　　　　　　　 　　ヽ
　|　　　　ｌ　　　　　　　　　　　　　　＼
　ヽ　　　 -一''''''"~~｀'ー--､　　　-一'''''''ー-､.
　　ヽ ＿＿＿＿(⌒)(⌒)⌒)　)　　(⌒＿(⌒)⌒)⌒))
`

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
	io.WriteString(out, ERROR_AA)
}
