package employee

import (
	"database/sql"
	"log"
	"randomize721/go-fiber-server/config/db"
	"randomize721/go-fiber-server/utils/responses"

	"github.com/gofiber/fiber/v2"
)

// @summary		Show All Employees
// @description	Get all employees
// @tags			Employee
// @produce		json
// @success		200	{object}	responses.Data
// @router			/employee [get]
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

// @summary		Show Employee by Id
// @description	Get employee by Id
// @tags			Employee
// @Param			id	path	string	true	"Employee Id"
// @produce		json
// @success		200	{object}	responses.Data
// @router			/employee/{id} [get]
func GetEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")

	query := "SELECT id, name, phone, address FROM app.employees WHERE id=$1"
	row := db.Session.QueryRow(query, id)

	res := Employee{}
	if err := row.Scan(&res.Id, &res.Name, &res.Phone, &res.Address); err != nil || err == sql.ErrNoRows {
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
		Data:            res,
	}
	return c.Status(200).JSON(response)
}

// @summary		Add Employee
// @description	Add new employee
// @tags			Employee
// @accept			json
// @produce		json
// @Param			employee	body		EmployeeField	true	"New Employee"
// @success		200			{object}	responses.Default
// @router			/employee [post]
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

	query := "INSERT INTO app.employees(name, phone, address) VALUES($1, $2, $3)"
	_, err := db.Session.Query(query, newEmployee.Name, newEmployee.Phone, newEmployee.Address)
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

// @summary		Update Employee
// @description	Update employee data
// @tags			Employee
// @accept			json
// @produce		json
// @Param			id			path		string			true	"Employee Id"
// @Param			employee	body		EmployeeField	true	"New Employee Data"
// @success		200			{object}	responses.Default
// @router			/employee/{id} [put]
func UpdateEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")
	employee := new(EmployeeField)

	if err := c.BodyParser(employee); err != nil {
		response := responses.Default{
			ResponseCode:    400,
			ResponseMessage: err.Error(),
		}
		log.Println(err.Error())
		return c.Status(400).JSON(response)
	}

	query := "UPDATE app.employees SET name=$1, phone=$2, address=$3 WHERE id=$4"
	_, err := db.Session.Query(query, employee.Name, employee.Phone, employee.Address, id)
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

// @summary		Delete Employee
// @description	Delete employee by Id
// @tags			Employee
// @Param			id	path	string	true	"Employee Id"
// @produce		json
// @success		200	{object}	responses.Data
// @router			/employee/{id} [delete]
func DeleteEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")

	query := "DELETE FROM app.employees WHERE id=$1"
	_, err := db.Session.Query(query, id)
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
