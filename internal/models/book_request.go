package models

type BookRequest struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Year       int    `json:"year"`
	ISBN       string `json:"isbn"`
	OutOfStock bool   `json:"out_of_stock"`
	Read       bool   `json:"read"`
	Rating     int    `json:"rating"`
}
