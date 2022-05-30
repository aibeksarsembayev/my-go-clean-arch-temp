package models

import "time"

// Post ...
type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	User      User      `json:"user"`
	Category  Category  `json:"category"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// PostUseCase represents posts's usecases
