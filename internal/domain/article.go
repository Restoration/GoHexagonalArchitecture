package domain

import (
	"time"
)

//Article .
type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Articles .
type Articles struct {
	Articles []Article `json:"articles"`
}

// ArticleCommonResponse .
type ArticleCommonResponse struct {
	LastID  string      `json:"last_id"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//ArticleParams .
type ArticleParams struct {
	ID        int
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
