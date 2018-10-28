package handler

import (
	"encoding/json"
	"net/http"

	"github.com/isayme/go-httpecho/app"
)

var versionInfo []byte

func init() {
	versionInfo, _ = json.Marshal(map[string]string{
		"name":    app.Name,
		"version": app.Version,
	})
}

// Version version hanlder
func Version(w http.ResponseWriter, r *http.Request) {
	resHeader := w.Header()
	resHeader.Set(app.HeaderContentType, app.MIMEApplicationJSON)
	w.Write(versionInfo)
}
