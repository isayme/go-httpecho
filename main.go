package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/isayme/go-httpecho/app/handler"
)

var port = flag.Uint("p", 3000, "listen port")
var certFile = flag.String("cert_file", "", "cert file")
var keyFile = flag.String("key_file", "", "key file")

func main() {
	flag.Parse()

	http.HandleFunc("/version", handler.Version)
	http.HandleFunc("/", handler.Echo)

	addr := fmt.Sprintf(":%d", *port)
	if *certFile != "" && *keyFile != "" {
		log.Printf("listen https %s ...\n", addr)
		log.Fatal(http.ListenAndServeTLS(addr, *certFile, *keyFile, nil))
	} else {
		log.Printf("listen http %s ...\n", addr)
		log.Fatal(http.ListenAndServe(addr, nil))
	}
}
