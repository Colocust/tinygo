package tinygo

import (
	"net/http"
)

const (
	noWritten     = -1
	defaultStatus = 200
)

type ResponseWriter interface {
	http.ResponseWriter

	// 写入httpStatus
	WriteHeaderNow()

	Reset()

	Written() bool
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

func (w *responseWriter) Reset() {
	w.size = noWritten
	w.status = defaultStatus
}

func (w *responseWriter) WriteHeader(status int) {
	if status > 0 && w.status != status {
		if !w.Written() {
			w.status = status
		}
	}
}

func (w *responseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.size = 0
		w.ResponseWriter.WriteHeader(w.status)
	}
}

func (w *responseWriter) Written() bool {
	return w.size != noWritten
}

func (w *responseWriter) Write(data []byte) (n int, err error) {
	w.WriteHeaderNow()

	n, err = w.ResponseWriter.Write(data)
	w.size += n
	return
}
