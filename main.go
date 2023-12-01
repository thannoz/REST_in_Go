package main

import (
	"clean/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	const port string = ":8080"
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	router.HandleFunc("/posts", controller.GetPosts).Methods("GET")
	router.HandleFunc("/posts", controller.CreatePost).Methods("POST")

	log.Println("Server listening on port", port)
	http.ListenAndServe(port, router)
}
