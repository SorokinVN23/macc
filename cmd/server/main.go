package main

import (
	"macc/internal/srv/http"
	setpack "macc/internal/srv/settings"
)

func main() {
	settings := setpack.GetSettings()
	http.Start(settings)
}
