package settings

import (
	"flag"
	//"fmt"
	dom "macc/internal/domains"
)

func getFlags(obj *dom.SrvSettings) {
	flag.StringVar(&(obj.Address), "a", "localhost:8080", "host:port")
	flag.Parse()

	/* for _, s := range flag.Args() {
		panic(fmt.Sprintf("invalid flag %s", s))
	} */
}
