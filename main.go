package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/isayme/go-httpecho/app/handler"
)

var port = flag.Uint("p", 3000, "listen port")

func main() {
	flag.Parse()

	http.HandleFunc("/version", handler.Version)
	http.HandleFunc("/", handler.Echo)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
