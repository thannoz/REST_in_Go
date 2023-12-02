package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	muxDispatcher = mux.NewRouter()
)

type muxRouter struct{}

func NewMuxRouter() PostHandler {
	return &muxRouter{}
}

func (r *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	fmt.Println("get posts called...")
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (r *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	fmt.Println("create post called...")
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (r *muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP Server running on port %v\n", port)
	http.ListenAndServe(port, muxDispatcher)
}
