package models

import "time"

// User ...
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Confirm   string    `json:"confirm"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
