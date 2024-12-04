package main

import (
	"fmt"
	dom "macc/internal/domains"
	"macc/internal/srv/http"
)

func main() {
	settings := &dom.SrvSettings{}
	GetFlags(settings)
	fmt.Printf("%+v", settings)
	http.Start(settings)
}
