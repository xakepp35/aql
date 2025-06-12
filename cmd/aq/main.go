package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xakepp35/aql/pkg/asf"
	"github.com/xakepp35/aql/pkg/ast"
	"github.com/xakepp35/aql/pkg/lexer"
	"github.com/xakepp35/aql/pkg/vm"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: aq <run|compile|explain> \"program\"\n")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	program := os.Args[1]
	cmd := "run"
	if len(os.Args) > 2 {
		cmd = os.Args[1]
		program = os.Args[2]
	}
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

var main_run = main_compile

// func main_run(src string) error {
// 	res, err := vm.Run([]byte(src), nil)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Fprintln(os.Stdout, res)
// 	return nil
// }

func main_compile(src string) error {
	// parse source code into ast tree
	a, err := ast.Parse([]byte(src))
	if err != nil {
		return fmt.Errorf("compile: %w", err)
	}
	// print tree in terminal
	var sb strings.Builder
	a.BuildString(&sb)
	fmt.Println(sb.String())

	// create virtual machine
	m := vm.New()
	// reserve emitter space, for less reallocs during bytecode emission
	m.Emit = make(asf.Emitter, 0, 256)
	a.P0(&m.Emit)
	a.P1(&m.Emit)
	a.P2(&m.Emit)

	// print vm bytecode as hex
	fmt.Println(m.Emit.AsHex())

	// run vm
	m.Run()

	// print stack dump after run
	fmt.Println(m.Dump())

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
