package employee

import (
	"database/sql"
	"log"
	"randomize721/go-fiber-server/config/db"
	"randomize721/go-fiber-server/utils/responses"

	"github.com/gofiber/fiber/v2"
)

//	@summary		Show All Employees
//	@description	Get all employees
//	@tags			Employee
//	@produce		json
//	@success		200	{object}	responses.Data
//	@router			/employee [get]
func GetAllEmployees(c *fiber.Ctx) error {
	query := "SELECT id, name, phone, address FROM app.employees ORDER BY id"
	rows, err := db.Session.Query(query)
	if err != nil || err == sql.ErrNoRows {
		response := responses.Default{
			ResponseCode:    500,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(500).JSON(response)
	}
	defer rows.Close()

	res := Employees{}
	for rows.Next() {
		employee := Employee{}
		if err := rows.Scan(&employee.Id, &employee.Name, &employee.Phone, &employee.Address); err != nil {
			response := responses.Default{
				ResponseCode:    500,
				ResponseMessage: err.Error(),
			}
			log.Println(err.Error())
			return c.Status(500).JSON(response)
		}
		res.Employees = append(res.Employees, employee)
	}

	response := responses.Data{
		ResponseCode: 200,
	}
	if len(res.Employees) < 1 {
		response.ResponseMessage = "Tidak ada data ditemukan."
	} else {
		response.ResponseMessage = "Berhasil mengambil data."
		response.Data = res.Employees
	}

	return c.Status(200).JSON(response)
}

//	@summary		Add Employee
//	@description	Add new employee
//	@tags			Employee
//	@accept			json
//	@produce		json
//	@Param			employee	body		EmployeeField	true	"New Employee"
//	@success		200			{object}	responses.Default
//	@router			/employee [post]
func AddEmployee(c *fiber.Ctx) error {
	newEmployee := new(EmployeeField)

	if err := c.BodyParser(newEmployee); err != nil {
		response := responses.Default{
			ResponseCode:    400,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(400).JSON(response)
	}
	log.Println(newEmployee)

	query := "INSERT INTO app.employees(name, phone, address) VALUES($1, $2, $3)"
	res, err := db.Session.Query(query, newEmployee.Name, newEmployee.Phone, newEmployee.Address)
	if err != nil {
		response := responses.Default{
			ResponseCode:    500,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(500).JSON(response)
	}
	log.Println(res)

	response := responses.Default{
		ResponseCode:    200,
		ResponseMessage: "Berhasil menginputkan data.",
	}
	return c.Status(200).JSON(response)
}
