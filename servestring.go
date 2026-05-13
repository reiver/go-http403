package http403

import (
	"net/http"
	"io"
)

func ServeString(responseWriter http.ResponseWriter, value string) error {
	if nil == responseWriter {
		return ErrNilHTTPResponseWriter
	}

	responseWriter.WriteHeader(StatusCode)

	if "" != value {
		io.WriteString(responseWriter, value)
	}

	return nil
}
