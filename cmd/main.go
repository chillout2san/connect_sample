package main

import (
	"example/gen/greet/v1/greetv1connect"
	"example/presentation"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	serviceHandler := presentation.NewGreetServiceHandler()
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(serviceHandler)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
