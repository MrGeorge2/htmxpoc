package handlers

import (
	"gopoc/internal/api/errors"
	"gopoc/internal/db/repositories"
	"gopoc/pkg/templutils"
	uihtml "gopoc/ui/html"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MovieHandler struct {
	movieRepository repositories.Queries
}

func (mh MovieHandler) registerRoutes(router *echo.Echo) {
	router.GET("/", mh.homePage)

	router.DELETE("/movie/:id", mh.delete)
	router.GET("/add-movie", mh.addMovieGet)
	router.POST("/add-movie", mh.addMoviePost)
}

func (mh MovieHandler) homePage(ctx echo.Context) error {
	movies, err := mh.movieRepository.GetMovies(ctx.Request().Context())
	if err != nil {
		return err
	}

	component := uihtml.Index(movies)
	return templutils.Render(ctx, component, 200)
}

func (mh MovieHandler) delete(ctx echo.Context) error {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		return errors.NewHandlerError(err, http.StatusBadRequest)
	}

	err = mh.movieRepository.DeleteMovie(ctx.Request().Context(), id)

	if err != nil {
		return errors.NewHandlerError(err, http.StatusInternalServerError)
	}

	return nil
}

func (mh MovieHandler) addMovieGet(ctx echo.Context) error {
	componenet := uihtml.AddMovieForm()
	return templutils.Render(ctx, componenet, 200)
}

func (mh MovieHandler) addMoviePost(ctx echo.Context) error {

	movie := repositories.Movie{
		MovieName: ctx.FormValue("name"),
	}

	id, err := mh.movieRepository.InsertMovie(ctx.Request().Context(), movie.MovieName)

	if err != nil {
		return err
	}

	movie.ID = id

	component := uihtml.AddMovieResult(movie)
	return templutils.Render(ctx, component, 200)
}
