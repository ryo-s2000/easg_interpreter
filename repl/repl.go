package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
	"monkey/token"
)

const PROMTP = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
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
