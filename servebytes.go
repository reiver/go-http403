package http403

import (
	"net/http"
	"unsafe"
)

func ServeBytes(responseWriter http.ResponseWriter, bytes []byte) error {
	var str string = unsafe.String(unsafe.SliceData(bytes), len(bytes))
	return ServeString(responseWriter, str)
}
