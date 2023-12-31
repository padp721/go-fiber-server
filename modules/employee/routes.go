package employee

import "github.com/gofiber/fiber/v2"

func Setup(parent fiber.Router) {
	parent.Get("/", GetAllEmployees)
	parent.Get("/:id<string>", GetEmployeeById)
	parent.Post("/", AddEmployee)
	parent.Put("/:id<string>", UpdateEmployeeById)
	parent.Delete("/:id<string>", DeleteEmployeeById)
}
