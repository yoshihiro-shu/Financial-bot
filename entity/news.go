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
	Score       int32     `json:"score"`
	PublishedAt time.Time `json:"published_at"`
	ProviderID  int32     `json:"provider_id"`
	CategoryID  int32     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
