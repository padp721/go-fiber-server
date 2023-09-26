package book

import (
	"github.com/gocql/gocql"
)

type Book struct {
	Id     gocql.UUID `json:"id"`
	Author string     `json:"author"`
	Title  string     `json:"title"`
	Year   int        `json:"year"`
}

type BookField struct {
	Author string `json:"author" example:"New author"`
	Title  string `json:"title" example:"New book title"`
	Year   int    `json:"year" example:"1999"`
}
