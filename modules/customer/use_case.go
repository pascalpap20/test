package customer

import (
	"crud/entity"
	"crud/repository"
	"time"
)

type UseCaseCustomer interface {
	CreateCustomer(user CustomerParam) (entity.Customer, error)
	GetCustomerById(id uint) (entity.Customer, error)
	GetCustomers(name string, email string, page uint) ([]entity.Customer, error)
	UpdateCustomerById(customer UpdateCustomer, id uint) (entity.Customer, error)
	DeleteCustomerById(id uint) (entity.Customer, error)
}

type useCaseCustomer struct {
	customerRepo repository.CustomerInterfaceRepo
}

func (uc useCaseCustomer) CreateCustomer(customer CustomerParam) (entity.Customer, error) {
	var newCustomer *entity.Customer

	//forbiddenCustomer, _ := helper.GetData()
	//err := helper.ValidateCreateCustomer(customer, forbiddenCustomer)
	//if err != nil {
	//	return *newCustomer, err
	//}

	newCustomer = &entity.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}
	return *newCustomer, nil
}

func (uc useCaseCustomer) GetCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	customer, err := uc.customerRepo.GetCustomerById(id)
	return customer, err
}

func (uc useCaseCustomer) GetCustomers(name string, email string, page uint) ([]entity.Customer, error) {
	var customer []entity.Customer
	customer, err := uc.customerRepo.GetCustomers(name, email, page)
	return customer, err
}

func (uc useCaseCustomer) UpdateCustomerById(customer UpdateCustomer, id uint) (entity.Customer, error) {
	var updateCustomer *entity.Customer

	updateCustomer = &entity.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
	}

	_, err := uc.customerRepo.UpdateCustomerById(updateCustomer, id)
	if err != nil {
		return *updateCustomer, err
	}
	return *updateCustomer, nil
}

func (uc useCaseCustomer) DeleteCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	customer, err := uc.customerRepo.DeleteCustomerById(id)
	return customer, err
}
