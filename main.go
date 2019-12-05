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

	addr := fmt.Sprintf(":%d", *port)
	log.Printf("listen %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
