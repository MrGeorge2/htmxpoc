package errors

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	notFound            = "404 Not Found"
	badRequest          = "400 Bad Request"
	internalServerError = "500 Internal Server Error"
)

var responseCodes = map[int]string{
	http.StatusNotFound:            notFound,
	http.StatusBadRequest:          badRequest,
	http.StatusInternalServerError: internalServerError,
}

type HandlerError struct {
	err          error
	responseCode int
}

func NewHandlerError(err error, responseCode int) HandlerError {
	return HandlerError{
		err:          err,
		responseCode: responseCode,
	}
}

func (he HandlerError) Error() string {
	return he.err.Error()
}

func (he HandlerError) GetResponseError() *echo.HTTPError {
	val := responseCodes[he.responseCode]

	return echo.NewHTTPError(he.responseCode, val)
}
