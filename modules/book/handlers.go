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
	rows := cassandra.Session.Query(cql).Iter()

	var books []Book
	var book Book
	for rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year) {
		books = append(books, book)
	}

	if err := rows.Close(); err != nil {
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

func GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")

	cql := "SELECT id, title, author, year FROM books WHERE id=?"
	rows := cassandra.Session.Query(cql, id).Iter()

	var books []Book
	var book Book
	for rows.Scan(&book.Id, &book.Title, &book.Author, &book.Year) {
		books = append(books, book)
	}

	if err := rows.Close(); err != nil {
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

func UpdateBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(BookField)

	if err := c.BodyParser(book); err != nil {
		response := responses.Default{
			ResponseCode:    400,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(400).JSON(response)
	}

	cql := "UPDATE books SET author=?, title=?, year=? WHERE id=?"
	err := cassandra.Session.Query(cql, book.Author, book.Title, book.Year, id).Exec()
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
		ResponseMessage: "Berhasil mengubah data.",
	}
	return c.Status(200).JSON(response)
}

func DeleteBookById(c *fiber.Ctx) error {
	id := c.Params("id")

	cql := "DELETE FROM books WHERE id=?"
	err := cassandra.Session.Query(cql, id).Exec()
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
		ResponseMessage: "Berhasil menghapus data.",
	}
	return c.Status(200).JSON(response)
}
