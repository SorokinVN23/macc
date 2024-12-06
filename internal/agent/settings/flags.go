package settings

import (
	"flag"
	//"fmt"
	dom "macc/internal/domains"
)

func getFlags(obj *dom.AgentSettings) {
	flag.StringVar(&(obj.Address), "a", "localhost:8080", "host:port")
	flag.IntVar(&(obj.ReportInterval), "r", 10, "ReportInterval int second")
	flag.IntVar(&(obj.PollInterval), "p", 2, "PollInterval int second")
	flag.Parse()

	/* for _, s := range flag.Args() {
		panic(fmt.Sprintf("invalid flag %s", s))
	} */
}
