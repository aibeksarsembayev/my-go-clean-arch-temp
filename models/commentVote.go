package models

import "time"

// Post Vote ...
type CommentVote struct {
	ID        int       `json:"id"`
	Value     bool      `json:"value"`
	User      User      `json:"user"`
	Comment   Comment   `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
