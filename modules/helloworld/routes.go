package helloworld

import "github.com/gofiber/fiber/v2"

func Setup(parent fiber.Router) {
	parent.Get("/", GetHelloWorld)
	parent.Get("/params/:value<alpha>", GetHelloWorldWithParams)
}
