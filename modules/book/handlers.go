package book

import (
	"log"
	"randomize721/go-fiber-server/config/cassandra"
	"randomize721/go-fiber-server/utils/responses"

	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

func GetAllBooks(c *fiber.Ctx) error {
	cql := "SELECT id, title, author, year FROM books"

	iter := cassandra.Session.Query(cql).Iter()

	var books []Book
	var book Book
	for iter.Scan(&book.Id, &book.Title, &book.Author, &book.Year) {
		books = append(books, book)
	}

	if err := iter.Close(); err != nil {
		response := responses.Default{
			ResponseCode:    500,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(500).JSON(response)
	}

	response := responses.Data{
		ResponseCode:    200,
		ResponseMessage: "Berhasil mengambil data.",
		Data:            books,
	}
	return c.Status(200).JSON(response)
}

func AddBook(c *fiber.Ctx) error {
	newBook := new(BookField)

	if err := c.BodyParser(newBook); err != nil {
		response := responses.Default{
			ResponseCode:    400,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(400).JSON(response)
	}

	id, err := gocql.RandomUUID()
	if err != nil {
		response := responses.Default{
			ResponseCode:    500,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(500).JSON(response)
	}

	cql := "INSERT INTO books(id, author, title, year) VALUES(?, ?, ?, ?)"
	err = cassandra.Session.Query(cql, id, newBook.Author, newBook.Title, newBook.Year).Exec()
	if err != nil {
		response := responses.Default{
			ResponseCode:    500,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(500).JSON(response)
	}

	response := responses.Default{
		ResponseCode:    200,
		ResponseMessage: "Berhasil menambah data.",
	}
	return c.Status(200).JSON(response)
}
