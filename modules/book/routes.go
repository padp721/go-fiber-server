package book

import "github.com/gofiber/fiber/v2"

func Setup(parent fiber.Router) {
	parent.Get("/", GetAllBooks)
	parent.Post("/", AddBook)
	parent.Get("/:id<string>", GetBookById)
	parent.Put("/:id<string>", UpdateBookById)
}
