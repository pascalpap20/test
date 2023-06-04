package account

import (
	"crud/dto"
	"crud/entity"
	"time"
)

type AdminParam struct {
	RoleID     uint   `json:"role_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data AdminParam `json:"data"`
}

type FindAdmin struct {
	dto.ResponseMeta
	Data entity.Actor `json:"data"`
}

type FindAllAdmins struct {
	dto.ResponseMeta
	Data []entity.Actor `json:"data"`
}

type UpdateAdmin struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
