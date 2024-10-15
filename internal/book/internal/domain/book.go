package domain

import (
	"time"
)

type Book struct {
	ID          string    `json:"id" db:"id"`
	Author_id   string    `json:"author_id" db:"author_id"`
	Category_id string    `json:"category_id" db:"category_id"`
	Title       string    `json:"title" db:"title"`
	Author      string    `json:"author" db:"author"`
	Category    string    `json:"category" db:"category"`
	Stock       int       `json:"stock" db:"stock"`
	Borrowed    int       `json:"borrowed" db:"borrowed"`
	IsPopular   bool      `json:"is_popular" db:"is_popular"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
