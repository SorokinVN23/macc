package http

import (
	"net/http"

	"github.com/SorokinVN23/macc/internal/srv/api"
)

func Start() {
	mux := api.NewMux()
	http.ListenAndServe(":8080", mux)
}
