package api

import (
	"net/http"
)

// The interface for the router
type PostHandler interface {
	GET(uri string, f func(http.ResponseWriter, *http.Request))
	POST(uri string, f func(http.ResponseWriter, *http.Request))
	SERVE(port string)
}
