package handler

import "net/http"

func Greet(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("yo"))
}
