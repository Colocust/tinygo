package render

import "net/http"

type Render interface {
	Render(http.ResponseWriter) error
	WriteHttpContentType(http.ResponseWriter)
}
