package models

import (
	"context"
	"time"
)

// Post ...
type Post struct {
	PostID    int       `json:"post_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	Category  Category  `json:"category"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// PostRequestDTO ...
type PostRequestDTO struct {
	PostID    int       `json:"post_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Category  Category  `json:"category"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// PostResponceDTO ...
type PostResponceDTO struct {
	PostID    int       `json:"post_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int       `json:"user_id"`
	Category  Category  `json:"category"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// PostUsecase represents posts's usecases
type PostUsecase struct {
}

// PostRepository represent post's repository contact
type PostRepository interface {
	Create(ctx context.Context, post *Post) (id int, err error)
	Update(ctx context.Context, post *Post) (err error)
	GetAll(ctx context.Context, post *Post) (posts *[]PostRequestDTO, err error)
	GetByID(ctx context.Context, id int) (post *PostRequestDTO, err error)
	Delete(ctx context.Context, id int) (err error)
}
