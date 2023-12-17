package handlers

import (
	"fmt"
	"gopoc/internal/db/repositories"
	"gopoc/pkg/routerutils"
	uihtml "gopoc/ui/html"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MovieHandler struct {
	movieRepository repositories.Queries
}

func (mh MovieHandler) registerRoutes(router *mux.Router) {
	router.HandleFunc("/", mh.homePage).Methods(http.MethodGet)

	router.HandleFunc("/movie/{id}", mh.delete).Methods(http.MethodDelete)
	router.HandleFunc("/add-movie", mh.addMovie).Methods(http.MethodGet, http.MethodPost)
}

func (mh MovieHandler) homePage(w http.ResponseWriter, r *http.Request) {
	movies, err := mh.movieRepository.GetMovies(r.Context())
	if err != nil {
		routerutils.BadRequest(w, r)
		return
	}

	component := uihtml.Index(movies)

	err = component.Render(r.Context(), w)

	if err != nil {
		routerutils.BadRequest(w, r)
		return
	}
}

func (mh MovieHandler) delete(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Vars %v", mux.Vars(r))

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		routerutils.BadRequest(w, r)
		return
	}

	err = mh.movieRepository.DeleteMovie(r.Context(), id)

	if err != nil {
		routerutils.BadRequest(w, r)
		return
	}

	if err != nil {
		routerutils.BadRequest(w, r)
		return
	}
}

func (mh MovieHandler) addMovie(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		mh.addMovieGet(w, r)
		break

	case http.MethodPost:
		mh.addMoviePost(w, r)
		break

	default:
		routerutils.NotAlowed(w, r)
		break
	}
}

func (mh MovieHandler) addMovieGet(w http.ResponseWriter, r *http.Request) {
	componenet := uihtml.AddMovieForm()
	err := componenet.Render(r.Context(), w)

	if err != nil {
		routerutils.BadRequest(w, r)
		return
	}

}

func (mh MovieHandler) addMoviePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		routerutils.BadRequest(w, r)
		return
	}

	movie := repositories.Movie{
		MovieName: r.FormValue("name"),
	}

	id, err := mh.movieRepository.InsertMovie(r.Context(), movie.MovieName)

	if err != nil {
		routerutils.BadRequest(w, r)
		return
	}

	movie.ID = id

	uihtml.AddMovieResult(movie).Render(r.Context(), w)
}
