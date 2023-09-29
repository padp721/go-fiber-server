package authentication

import "github.com/gofiber/fiber/v2"

func Setup(parent fiber.Router) {
	parent.Post("/login", Login)
	parent.Get("/validate", ValidateJwt)
}
