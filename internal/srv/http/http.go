package http

import (
	"net/http"

	dom "macc/internal/domains"
	"macc/internal/srv/api"

	"github.com/go-chi/chi/v5"
)

func Start(settings *dom.SrvSettings) {
	router := chi.NewRouter()
	router.Get("/", api.List)
	router.Post("/update/{mtype}/{mname}/{mvalue}", api.Update)
	http.ListenAndServe(settings.Address, router)
}
