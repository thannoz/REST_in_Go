package firestore

import (
	"clean/entity"
	"clean/repository"
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	projectID      = "gorest-d4e9a"
	collectionName = "posts"
)

type repo struct{}

// NewFirestoreRepository returns a value that satisfies the PostRepository interface
func NewFirestoreRepository() repository.PostRepository {
	return &repo{}
}

func createFirestoreClient(ctx context.Context, projectID string) (*firestore.Client, error) {
	opt := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("error initializing firebase: %v", err)
	}
	client, err := app.Firestore(context.TODO())
	if err != nil {
		log.Fatalf("Failed to instatiate a firestore client: %v", err)
		return nil, err
	}
	return client, nil
}

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := createFirestoreClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed saving post: %v", err)
		return nil, err
	}
	return post, nil
}

func (r *repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := createFirestoreClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	posts := []entity.Post{}
	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
