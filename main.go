package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/isayme/go-httpecho/app/handler"
)

var port = flag.Uint("p", 3000, "listen port")

func main() {
	flag.Parse()

	http.HandleFunc("/version", handler.Version)
	http.HandleFunc("/", handler.Echo)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
