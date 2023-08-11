package repository

import (
	"github.com/qomaindo-dev/go-bookstore/app/models"
	"github.com/qomaindo-dev/go-bookstore/graph/model"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	CreateEmployee(employeeInput *model.EmployeeInput) (*models.Employee, error)
	UpdateEmployee(employeeInput *model.Employee, id int) error
	DeleteEmployee(id int) error
	GetOneEmplyee(id int) (*models.Employee, error)
	GetAllEmployees() ([]*model.Employee, error)
}

type EmployeeService struct {
	Db *gorm.DB
}

var _ EmployeeRepository = &EmployeeService{}

func (es *EmployeeService) CreateEmployee(employeeInput *model.EmployeeInput) (*models.Employee, error) {
	emplyee := &models.Employee{
		Name:     employeeInput.Name,
		Email:    employeeInput.Email,
		Password: employeeInput.Password,
		Photo:    employeeInput.Photo,
		IsActive: employeeInput.IsActive,
	}

	err := es.Db.Create(&emplyee).Error

	return emplyee, err
}

func (es *EmployeeService) UpdateEmployee(employeeInput *model.Employee, id int) error {
	employee := &models.Employee{
		ID:       id,
		Name:     employeeInput.Name,
		Email:    employeeInput.Email,
		Password: employeeInput.Password,
		Photo:    employeeInput.Photo,
		IsActive: employeeInput.IsActive,
	}

	err := es.Db.Model(&employee).Where("id = ?", id).Updates(employee).Error

	return err
}

func (es *EmployeeService) DeleteEmployee(id int) error {
	employee := &models.Book{}
	err := es.Db.Delete(employee, id).Error

	return err
}

func (es *EmployeeService) GetOneEmplyee(id int) (*models.Employee, error) {
	employee := &models.Employee{}
	err := es.Db.Where("id = ?", id).First(employee).Error

	return employee, err
}

func (es *EmployeeService) GetAllEmployees() ([]*model.Employee, error) {
	employees := []*model.Employee{}
	err := es.Db.Find(&employees).Error

	return employees, err
}
