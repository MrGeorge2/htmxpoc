package templutils

import (
	"gopoc/internal/api/errors"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, cmp templ.Component, status int) error {
	c.Response().Writer.WriteHeader(status)

	err := cmp.Render(c.Request().Context(), c.Response().Writer)

	if err != nil {
		return errors.NewHandlerError(err, http.StatusInternalServerError)
	}

	return nil
}
