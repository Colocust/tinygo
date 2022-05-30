package render

import (
	"encoding/json"
	"net/http"
)

type Json struct {
	Data interface{}
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func (j Json) Render(w http.ResponseWriter) {
	if err := writeJson(w, j.Data); err != nil {
		panic(err)
	}
	return
}

func writeJson(w http.ResponseWriter, data interface{}) error {
	writeContentType(w, jsonContentType)
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(jsonBytes)
	return err
}

func (j Json) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, jsonContentType)
}
