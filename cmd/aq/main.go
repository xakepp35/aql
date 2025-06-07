package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"github.com/xakepp35/aql/pkg/aqc"
	"github.com/xakepp35/aql/pkg/lexer"
	"github.com/xakepp35/aql/pkg/vm"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: aq <run|compile|explain> \"program\"\n")
}

func main() {
	if len(os.Args) < 3 {
		usage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	program := os.Args[2]

	switch cmd {
	case "run":
		err := main_run(program)
		if err != nil {
			fatal(err)
		}

	case "compile":
		err := main_compile(program)
		if err != nil {
			fatal(err)
		}

	case "lexer":
		err := main_lexer(program)
		if err != nil {
			fatal(err)
		}

	default:
		usage()
		os.Exit(1)
	}
}

func main_run(src string) error {
	res, err := vm.Run([]byte(src), nil)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, res)
	return nil
}

func main_compile(src string) error {
	e := vm.NewCompiler()
	if err := aqc.Compile([]byte(src), e); err != nil {
		return fmt.Errorf("compile: %w", err)
	}
	bin, err := e.MarshalBinary()
	if err != nil {
		return err
	}
	// выводим как hex‑строку чтобы не бить терминал бинарщиной
	fmt.Fprintln(os.Stdout, hex.EncodeToString(bin))
	return nil
}

func main_lexer(src string) error {
	lx := lexer.New([]byte(src))
	for {
		tok := lx.Next()
		fmt.Printf("%v\t%q\n", tok.Kind, tok.Lit)
		if tok.Kind == lexer.TEOF {
			break
		}
	}
	return nil
}

func fatal(err error) {
	var perr *os.PathError
	if errors.As(err, &perr) {
		fmt.Fprintln(os.Stderr, perr.Error())
	} else {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	os.Exit(2)
}

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Usage: aqparse 'expression'")
// 		os.Exit(1)
// 	}
// 	expr := os.Args[1]
// 	lx := lexer.New([]byte(expr))
// 	for {
// 		tok := lx.Next()
// 		fmt.Printf("%v\t%q\n", tok.Kind, tok.Lit)
// 		if tok.Kind == lexer.TEOF {
// 			break
// 		}
// 	}
// 	fmt.Fprintf(os.Stdout, "LEXED")
// 	prog := vm.NewProgram()
// 	err := aqc.Compile([]byte(expr), prog)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "compile error: %v\n", err)
// 		os.Exit(2)
// 	}

// 	fmt.Printf("OK: %#v\n", prog)

// 	this := vm.NewState()
// 	prog.Init(this)
// 	err = prog.Run(this)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
// 		os.Exit(2)
// 	}
// 	res := this.Pop()

// 	fmt.Printf("Eval: %v\n", res)
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Usage: aqparse 'expression'")
// 		os.Exit(1)
// 	}
// 	expr := os.Args[1]
// 	lx := lexer.New([]byte(expr))
// 	for {
// 		tok := lx.Next()
// 		fmt.Printf("%v\t%q\n", tok.Kind, tok.Lit)
// 		if tok.Kind == lexer.TEOF {
// 			break
// 		}
// 	}
// 	fmt.Fprintf(os.Stdout, "LEXED")
// 	node, err := ast.Parse([]byte(expr))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
// 		os.Exit(2)
// 	}

// 	fmt.Printf("OK: %#v\n", node)

// 	this := vm.NewState()
// 	this.Set("x", 42.7)
// 	err = node.Eval(this)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
// 		os.Exit(2)
// 	}
// 	res := this.Pop()

// 	fmt.Printf("Eval: %v\n", res)
// }
