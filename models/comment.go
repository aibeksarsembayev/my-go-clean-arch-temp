package models

import "time"

type Comment struct {
	CommentID int       `json:"comment_id"`
	Content   string    `json:"content"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
