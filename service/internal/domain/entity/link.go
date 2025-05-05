package entity

import "time"

type Link struct {
	ID        int       `json:"id,omitempty"`
	URL       string    `json:"url"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateLinkRequest struct {
	URL string `json:"url"`
}
