package http

import (
	"net/http"

	dom "macc/internal/domains"
	"macc/internal/srv/api"

	"github.com/go-chi/chi/v5"
)

func Start(settings *dom.SrvSettings) {
	//mux := api.NewMux()
	//http.ListenAndServe(":8080", mux)

	router := chi.NewRouter()
	router.Get("/", api.List)
	router.Post("/update/{mtype}/{mname}/{mvalue}", api.Update)
	http.ListenAndServe(settings.Address, router)
}
