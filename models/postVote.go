package models

import "time"

// Post Vote ...
type PostVote struct {
	ID        int       `json:"id"`
	Value     bool      `json:"value"`
	User      User      `json:"user"`
	Post      Post      `json:"post"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
