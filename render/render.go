package render

import "net/http"

type Render interface {
	Render(http.ResponseWriter) error
	WriteHttpContentType(http.ResponseWriter)
}

func writeContentType(writer http.ResponseWriter, value []string) {
	header := writer.Header()
	if val := header["Content-Type"]; len(val) != 0 {
		header["Content-Type"] = value
	}
}
