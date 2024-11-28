package http

import (
	"net/http"

	"macc/internal/srv/api"
)

func Start() {
	mux := api.NewMux()
	http.ListenAndServe(":8080", mux)
}
