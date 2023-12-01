package service

import (
	"clean/entity"
	"clean/repository"
	"clean/repository/firestore"
	"errors"
	"math/rand"
	"strings"
)

var (
	repo repository.PostRepository = firestore.NewFirestoreRepository()
)

type PostService interface {
	Validate(*entity.Post) error
	Create(*entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

func NewPostService() PostService {
	return &service{}
}

func (srv *service) Validate(post *entity.Post) error {
	var errs []string

	if post == nil {
		errs = append(errs, "The post is empty")
	}
	if post.Title == "" {
		errs = append(errs, "The post title is empty")
	}

	if post.Text == "" {
		errs = append(errs, "The post text is empty")
	}
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}

func (srv *service) Create(post *entity.Post) (*entity.Post, error) {

	post.ID = rand.Int63()
	return repo.Save(post)
}

func (srv *service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}
