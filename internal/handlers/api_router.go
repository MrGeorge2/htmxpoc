package handlers

import (
	"gopoc/internal/db/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiRouter struct {
	movieHandler MovieHandler
}

func RegisterRouter(movieRouter repositories.Queries) error {
	var router = ApiRouter{
		movieHandler: MovieHandler{
			movieRepository: movieRouter,
		},
	}

	return router.registerRoutes()
}

func (r *ApiRouter) registerRoutes() error {
	mux := mux.NewRouter()

	r.movieHandler.registerRoutes(mux)

	http.Handle("/", mux)

	return http.ListenAndServe(":8080", nil)
}
