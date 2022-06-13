package models

import "time"

// CommentVote ...
type CommentVote struct {
	CommentVoteID int       `json:"comment_vote_id"`
	Value         bool      `json:"value"`
	UserID        int       `json:"user_id"`
	CommentID     int       `json:"comment_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
