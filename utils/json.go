package utils

import "net/http"

func WriteJsonError(w *http.ResponseWriter, err error) {
	WriteJson(w, []byte(`{"error": true, "message": "`+err.Error()+`"}`))
}

func WriteJsonString(w *http.ResponseWriter, s string) {
	WriteJson(w, []byte(s))
}

func WriteJson(w *http.ResponseWriter, b []byte) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Write(b)
}
