package book

import "github.com/gofiber/fiber/v2"

func Setup(parent fiber.Router) {
	parent.Get("/", GetAllBooks)
	parent.Get("/:id<string>", GetBookById)
	parent.Post("/", AddBook)
	parent.Put("/:id<string>", UpdateBookById)
	parent.Delete("/:id<string>", DeleteBookById)
}
