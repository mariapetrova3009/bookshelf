package models

import "time"

type Book struct {
	Id         uint64    `json:"id"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Year       int       `json:"year"`
	ISBN       string    `json:"isbn,omitempty"`
	OutOfStock bool      `json:"out_of_stock"`
	Read       bool      `json:"read"`
	Rating     int       `json:"rating,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
