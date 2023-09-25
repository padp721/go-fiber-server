package employee

import "github.com/google/uuid"

type Employee struct {
	Id      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}

type EmployeeField struct {
	Name    string `json:"name" example:"Angga"`
	Phone   string `json:"phone" example:"0876543219"`
	Address string `json:"address" example:"Jl. Siulan"`
}
