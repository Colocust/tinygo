package render

import (
	"encoding/json"
	"net/http"
)

type Json struct {
	Data any
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func (r *Json) Render(writer http.ResponseWriter) error {
	r.WriteHttpContentType(writer)
	bytes, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	return err
}

func (r *Json) WriteHttpContentType(writer http.ResponseWriter) {
	writeContentType(writer, jsonContentType)
}
