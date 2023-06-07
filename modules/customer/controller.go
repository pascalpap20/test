package customer

import (
	"crud/dto"
)

type ControllerCustomer interface {
	CreateCustomer(req CustomerParam) (any, error)
	GetCustomerById(id uint) (FindCustomer, error)
	GetCustomers(name string, email string, page uint) (FindAllCustomers, error)
	UpdateCustomerById(req UpdateCustomer, id uint) (any, error)
	DeleteCustomerById(id uint) (FindCustomer, error)
}

type controllerCustomer struct {
	customerUseCase UseCaseCustomer
}

func (uc controllerCustomer) CreateCustomer(req CustomerParam) (any, error) {

	customer, err := uc.customerUseCase.CreateCustomer(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create customer",
			Message:      "Success create",
			ResponseTime: "",
		},
		Data: CustomerParam{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}

func (uc controllerCustomer) GetCustomerById(id uint) (FindCustomer, error) {
	var res FindCustomer
	customer, err := uc.customerUseCase.GetCustomerById(id)
	if err != nil {
		return FindCustomer{}, err
	}
	res.Data = customer
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get a customer",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}

func (uc controllerCustomer) GetCustomers(name string, email string, page uint) (FindAllCustomers, error) {
	var res FindAllCustomers
	customer, err := uc.customerUseCase.GetCustomers(name, email, page)
	if err != nil {
		return FindAllCustomers{}, err
	}

	res.Data = customer
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get all customers",
		Message:      "Success",
		ResponseTime: "",
	}

	return res, nil
}

func (uc controllerCustomer) UpdateCustomerById(req UpdateCustomer, id uint) (any, error) {

	customer, err := uc.customerUseCase.UpdateCustomerById(req, id)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success update customer",
			Message:      "Success update",
			ResponseTime: "",
		},
		Data: CustomerParam{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Avatar:    customer.Avatar,
		},
	}
	return res, nil
}

func (uc controllerCustomer) DeleteCustomerById(id uint) (FindCustomer, error) {
	var res FindCustomer
	customer, err := uc.customerUseCase.DeleteCustomerById(id)
	if err != nil {
		return FindCustomer{}, err
	}
	res.Data = customer
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success Delete Admin",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}
