package main

import (
	"context"
	"flag"
	"log"

	"github.com/xakepp35/aql/pkg/anp"
)

func main() {
	addr := flag.String("addr", ":7000", "TCP listen address")
	flag.Parse()
	srv := anp.NewServer()
	if err := srv.Listen(context.Background(), "tcp", *addr); err != nil {
		log.Fatal(err)
	}
	if err := srv.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
