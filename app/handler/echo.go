package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"mime"
	"net"
	"net/http"

	"github.com/isayme/go-httpecho/app"
	"github.com/isayme/go-logger"
)

// Echo echo handler
func Echo(w http.ResponseWriter, r *http.Request) {
	resBody := requestInfo{
		Proto:  r.Proto,
		Method: r.Method,
		Path:   r.RequestURI,
		IP:     r.RemoteAddr,
		Host:   r.Host,
	}

	// ip
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		resBody.IP = ip
	} else if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		resBody.IP = ip
	}

	// header
	header := convert(r.Header)
	resBody.Headers = header

	// query
	query := convert(r.URL.Query())
	if len(query) > 0 {
		resBody.Query = query
	}

	// body
	r.Body = http.MaxBytesReader(w, r.Body, app.MAX_BYTES_READER)
	rawBody, _ := io.ReadAll(r.Body)
	resBody.Data = string(rawBody)
	r.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := r.Header.Get(app.HeaderContentType)
	mediaType, _, _ := mime.ParseMediaType(contentType)

	if mediaType == app.MIMEApplicationJSON { // json
		var body M
		json.Unmarshal(rawBody, &body)
		resBody.Body = body
	} else if mediaType == app.MIMEApplicationForm { // form
		r.ParseForm()
		resBody.Form = convert(r.Form)
	} else if mediaType == app.MIMEMultipartForm { // form-data
		r.ParseMultipartForm(1024 * 1024)

		form := M{}
		if r.MultipartForm != nil {
			form["value"] = convert(r.MultipartForm.Value)

			files := M{}
			for k, v := range r.MultipartForm.File {
				infos := []M{}
				for _, file := range v {
					info := M{}
					info["filename"] = file.Filename
					info["size"] = file.Size
					infos = append(infos, info)
				}

				files[k] = infos
			}
			form["file"] = files
		}
		resBody.MultipartForm = form
	}

	// json response
	resHeader := w.Header()
	resHeader.Set(app.HeaderContentType, app.MIMEApplicationJSON)

	data, _ := json.Marshal(resBody)

	logger.Infof(string(data))

	_, pretty := query["pretty"]
	if pretty {
		data, _ := json.MarshalIndent(resBody, "", "    ")
		w.Write(data)
	} else {
		w.Write(data)
	}
}

// M response object map
type M map[string]interface{}

type requestInfo struct {
	Proto         string `json:"proto"`
	Method        string `json:"method"`
	Path          string `json:"path"`
	Headers       M      `json:"headers"`
	IP            string `json:"ip"`
	Host          string `json:"host"`
	Query         M      `json:"query,omitempty"`
	Data          string `json:"data,omitempty"`
	Form          M      `json:"form,omitempty"`
	MultipartForm M      `json:"multipartForm,omitempty"`
	Body          M      `json:"body,omitempty"`
}

func convert(values map[string][]string) M {
	m := M{}
	for k, v := range values {
		if len(v) > 1 {
			m[k] = v
		} else {
			m[k] = v[0]
		}
	}
	return m
}
