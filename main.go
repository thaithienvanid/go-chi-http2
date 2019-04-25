package main

import (
	"net/http"
	"golang.org/x/net/http2"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Payload interface{} `json:"payload"` 
}

func main() {
	handler := chi.NewRouter()

	res := Response{Code: "0", Message: "OK", Payload: "Hello"}

	handler.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, res)
	})

	server := http.Server{Addr: ":8080", Handler: handler}
	http2.ConfigureServer(&server, nil)

	server.ListenAndServeTLS("./certs/localhost.cert", "./certs/localhost.key")
}
