package handler

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"net/url"

	"github.com/isayme/go-httpecho/app"
)

// Echo echo handler
func Echo(w http.ResponseWriter, r *http.Request) {
	resBody := requestInfo{
		Method: r.Method,
		Path:   r.RequestURI,
		IP:     r.RemoteAddr,
	}

	// ip
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		resBody.IP = ip
	}

	// header
	header := M{}
	for k := range r.Header {
		header[k] = r.Header.Get(k)
	}
	resBody.Headers = header

	// query
	query := M{}
	for k, v := range r.URL.Query() {
		if len(v) > 1 {
			query[k] = v
		} else {
			query[k] = v[0]
		}
	}
	if len(query) > 0 {
		resBody.Query = query
	}

	// body
	rawBody, _ := ioutil.ReadAll(r.Body)
	resBody.Data = string(rawBody)

	contentType := r.Header.Get(app.HeaderContentType)
	mediaType, _, _ := mime.ParseMediaType(contentType)
	if mediaType == app.MIMEApplicationJSON { // json
		var body M
		json.Unmarshal(rawBody, &body)
		resBody.Body = body
	} else if mediaType == app.MIMEApplicationForm { // form
		form := M{}
		formValues, _ := url.ParseQuery(string(rawBody))
		for k, v := range formValues {
			if len(v) > 1 {
				form[k] = v
			} else {
				form[k] = v[0]
			}
		}
		resBody.Form = form
	}

	// json response
	resHeader := w.Header()
	resHeader.Set(app.HeaderContentType, app.MIMEApplicationJSON)

	data, _ := json.Marshal(resBody)
	w.Write(data)
}

// M response object map
type M map[string]interface{}

type requestInfo struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Headers M      `json:"headers"`
	IP      string `json:"ip"`
	Query   M      `json:"query,omitempty"`
	Data    string `json:"data,omitempty"`
	Form    M      `json:"form,omitempty"`
	Body    M      `json:"body,omitempty"`
}
