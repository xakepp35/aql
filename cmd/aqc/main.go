package aqc

import (
	"context"
	"flag"
	"log"

	"github.com/xakepp35/aql/pkg/anp"
)

func main() {
	addr := flag.String("addr", "localhost:7000", "Server address")
	query := flag.String("exec", "", "AQL program to execute")
	flag.Parse()

	ctx := context.Background()
	cli, err := anp.Dial(ctx, "tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	_ = query
	code := []byte{}
	// code := CompileAQL(*query) // или загружаем из файла
	if err := cli.Send(code); err != nil {
		log.Fatal(err)
	}
}
