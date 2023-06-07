package customer

import (
	"crud/dto"
	"crud/entity"
	"time"
)

type CustomerParam struct {
	FirstName string `json:"first_name" binding:"required,max=255"`
	LastName  string `json:"last_name" binding:"required,max=255"`
	Email     string `json:"email" binding:"required,email"`
	Avatar    string `json:"avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}

type FindCustomer struct {
	dto.ResponseMeta
	Data entity.Customer `json:"data"`
}

type FindAllCustomers struct {
	dto.ResponseMeta
	Data []entity.Customer `json:"data"`
}

type UpdateCustomer struct {
	FirstName string `json:"first_name" binding:"required,max=255"`
	LastName  string `json:"last_name" binding:"required,max=255"`
	Email     string `json:"email" binding:"required,email"`
	Avatar    string `json:"avatar"`
}
