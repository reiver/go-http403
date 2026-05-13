package http403

import (
	"net/http"
)

func Forbidden(responseWriter http.ResponseWriter, request *http.Request) {
	Serve(responseWriter, methods...)
}
