package core_http_server

import "net/http"

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}
