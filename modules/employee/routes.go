package employee

import "github.com/gofiber/fiber/v2"

func Setup(parent fiber.Router) {
	parent.Get("/", GetAllEmployees)
	parent.Post("/", AddEmployee)
}
