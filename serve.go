package http403

import (
	"net/http"
)

func Serve(responseWriter http.ResponseWriter) error {
	return ServeString(responseWriter, DefaultStatusText)
}
