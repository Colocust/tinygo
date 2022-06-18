package tinygo

import "net/http"

const (
	noWritten     = -1
	defaultStatus = 200
)

type ResponseWriter interface {
	http.ResponseWriter

	// 写入httpStatus
	WriteHeaderNow()

	// 判断当前是否已经写入过数据
	Written() bool

	// 获取当前的网络状态码
	Status() int

	// 获取当前已经写入的字节数量
	Size() int

	Reset()
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
		w.ResponseWriter.WriteHeader(w.status)
	}
}

func (w *responseWriter) Written() bool {
	return w.size == noWritten
}

func (w *responseWriter) Write(data []byte) (n int, err error) {
	w.WriteHeaderNow()

	w.size = 0
	n, err = w.ResponseWriter.Write(data)
	w.size += n

	return
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) Size() int {
	return w.size
}
