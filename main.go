package main

import (
	"clean/api"
	"clean/controller"
	"fmt"
	"net/http"
)

var (
	ctrl      controller.PostController = controller.NewPostController()
	muxRouter                           = api.NewMuxRouter()
)

func main() {

	const port string = ":8080"
	muxRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	muxRouter.GET("/posts", ctrl.GetPosts)
	muxRouter.POST("/posts", ctrl.CreatePost)
	//router.HandleFunc("/posts", ctrl.GetPosts).Methods("GET")
	//router.HandleFunc("/posts", ctrl.CreatePost).Methods("POST")

	muxRouter.SERVE(port)
	//log.Println("Server listening on port", port)
	//http.ListenAndServe(port, router)
}
