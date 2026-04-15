package handler

import (
	"net/http"

	handler "github.com/isayme/go-httpecho/app/handler"
)

func Echo(w http.ResponseWriter, r *http.Request) {
	handler.Echo(w, r)
}
