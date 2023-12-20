package handlers

import (
	"fmt"
	"gopoc/internal/api/errors"
	"gopoc/internal/db/repositories"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
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
	e := echo.New()

	absPath, _ := filepath.Abs("./ui/static")
	fmt.Println(absPath)
	e.Static("assets", absPath)

	r.movieHandler.registerRoutes(e)

	return e.Start(":8080")
}

func ExceptionCatchingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)

		if err == nil {
			return nil
		}

		handlerErr, ok := err.(errors.HandlerError)

		if ok {
			return handlerErr.GetResponseError()
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
}
