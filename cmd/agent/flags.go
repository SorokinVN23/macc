package main

import (
	"flag"
	"fmt"
	dom "macc/internal/domains"
)

func GetFlags(settings *dom.AgentSettings) {
	flag.StringVar(&(settings.Address), "a", "localhost:8080", "host:port")
	flag.IntVar(&(settings.ReportInterval), "r", 10, "ReportInterval int second")
	flag.IntVar(&(settings.PollInterval), "p", 2, "PollInterval int second")
	flag.Parse()

	for _, s := range flag.Args() {
		panic(fmt.Sprintf("invalid flag %s", s))
	}
}
