package main

import (
	"clean/entity"
	"encoding/json"
	"net/http"
)

var (
	posts []entity.Post
)

func init() {
	posts = []entity.Post{
		entity.Post{ID: 1, Title: "Money on my mind", Text: "Getting money."},
		entity.Post{ID: 2, Title: "Success", Text: "I will make it no matter what!"},
	}
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the posts data"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := entity.Post{}

	// read the request body with NewDecoder & format the json body into the post type
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	post.ID = len(posts) + 1
	posts = append(posts, post)

	w.WriteHeader(http.StatusOK)

	// convert the post object into json
	result, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	w.Write(result)
}
