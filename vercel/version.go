package handler

import (
	"net/http"

	handler "github.com/isayme/go-httpecho/app/handler"
)

func Version(w http.ResponseWriter, r *http.Request) {
	handler.Version(w, r)
}
