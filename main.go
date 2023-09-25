package main

import (
	"fmt"
	"log"
	"randomize721/go-fiber-server/config/cassandra"
	"randomize721/go-fiber-server/config/db"
	_ "randomize721/go-fiber-server/docs"
	"randomize721/go-fiber-server/modules/book"
	"randomize721/go-fiber-server/modules/employee"
	"randomize721/go-fiber-server/modules/helloworld"
	"strconv"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

//	@title			GO FIBER SERVER (Go v1.21.1)
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8000
//	@basePath	/api
//	@schemes	http https

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	if _, err := db.Connect(); err != nil {
		log.Fatalf("Cannot connect to Database: %v", err)
	}

	if _, err := cassandra.Connect(); err != nil {
		log.Fatalf("unable to connect Astra DB session: %v", err)
	}

	APP_PRINT_ROUTES, err := strconv.ParseBool(os.Getenv("APP_PRINT_ROUTES"))
	if err != nil {
		log.Fatal("Error loading APP_PRINT_ROUTES from environment variable!")
	}

	fiberConfig := fiber.Config{
		AppName:           "GO FIBER SERVER (Go v1.21.1)",
		CaseSensitive:     true,
		EnablePrintRoutes: APP_PRINT_ROUTES,
	}

	corsConfig := cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "*",
		AllowHeaders:     "*",
	}

	app := fiber.New(fiberConfig)

	app.Use(cors.New(corsConfig))
	app.Use(logger.New())

	app.Static("/", "./public")

	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")
	helloworld.Setup(api)

	employeeRoutes := api.Group("/employee")
	employee.Setup(employeeRoutes)

	bookRoutes := api.Group("/book")
	book.Setup(bookRoutes)

	APP_PORT := os.Getenv("APP_PORT")
	app.Listen(fmt.Sprintf(":%s", APP_PORT))
}
