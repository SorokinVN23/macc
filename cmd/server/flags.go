package main

import (
	"flag"
	"fmt"
	dom "macc/internal/domains"
)

func GetFlags(settings *dom.SrvSettings) {
	flag.StringVar(&(settings.Address), "a", "localhost:8080", "host:port")
	flag.Parse()

	for _, s := range flag.Args() {
		panic(fmt.Sprintf("invalid flag %s", s))
	}
}
