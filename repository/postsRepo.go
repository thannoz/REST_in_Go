package repository

import (
	"clean/entity"
)

type PostRepository interface {
	Save(*entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
