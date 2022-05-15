package tinygo

import "net/http"

const (
	noWritten     = -1
	defaultStatus = http.StatusOK
)

type ResponseWriter struct {
	status int
	size   int
	http.ResponseWriter
}

func (w *ResponseWriter) reset(writer http.ResponseWriter) {
	w.ResponseWriter = writer
	w.status = defaultStatus
	w.size = noWritten
}

func (w *ResponseWriter) Write(data []byte) (n int, err error) {
	w.WriteHeaderNow()
	n, err = w.ResponseWriter.Write(data)
	w.size += n
	return
}

func (w *ResponseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.ResponseWriter.WriteHeader(w.status)
		w.size = 0
	}
}

func (w *ResponseWriter) Written() bool {
	return w.size != noWritten
}

func (w *ResponseWriter) Status() int {
	return w.status
}

func (w *ResponseWriter) Size() int {
	return w.size
}