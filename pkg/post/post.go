package post

import "time"

type Post struct {
	ID        int `json:"id,omitempty"`
	Body      string `json:"body,omitempty"`
	Image     string `json:"Image,omitempty"`
	UserID    int `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}