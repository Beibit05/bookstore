package models

type Book struct {
	Id         int     `json:"id"`
	Title      string  `json:"title"`
	AuthorId   int     `json:"author_id"`
	CategoryId int     `json:"category_id"`
	Price      float64 `json:"price"`
}
