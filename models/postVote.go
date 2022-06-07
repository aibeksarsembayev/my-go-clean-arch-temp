package models

import "time"

// Post Vote ...
type PostVote struct {
	PostVoteID int       `json:"post_vote_id"`
	Value      bool      `json:"value"`
	UserID     int       `json:"user_id"`
	PostID     int       `json:"post_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
