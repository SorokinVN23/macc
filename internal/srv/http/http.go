package http

import (
	"net/http"

	"macc/internal/srv/api"

	"github.com/go-chi/chi/v5"
)

func Start() {
	//mux := api.NewMux()
	//http.ListenAndServe(":8080", mux)

	router := chi.NewRouter()
	router.Get("/", api.List)
	router.Get("/update/{mtype}/{mname}/{mvalue}", api.Update)
	http.ListenAndServe(":8080", router)
}
