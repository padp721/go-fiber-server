package helloworld

import (
	"randomize721/go-fiber-server/utils/responses"

	"github.com/gofiber/fiber/v2"
)

// @summary		Show hello world
// @description	return hello world string
// @tags			Hello World
// @produce		json
// @Success		200	{object}	responses.Default
// @router			/ [get]
func GetHelloWorld(c *fiber.Ctx) error {
	res := responses.Default{}
	res.ResponseCode = 200
	res.ResponseMessage = "Hello world!"

	return c.JSON(res)
}

// @summary		Show greetings
// @description	return greetings with sent parameter
// @tags			Hello World
// @produce		json
// @success		200		{object}	responses.Default
// @param			value	path		string	true	"Greetings to:"
// @router			/params/{value} [get]
func GetHelloWorldWithParams(c *fiber.Ctx) error {
	res := responses.Default{}
	res.ResponseCode = 200
	res.ResponseMessage = "Hello there, " + c.Params("value") + "!"

	return c.JSON(res)
}
