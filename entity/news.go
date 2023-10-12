package entity

import (
	"time"
)

type News struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Thumbnail   string    `json:"thumbnail"`
	PublishedAt time.Time `json:"published_at"`
}
