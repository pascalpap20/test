package account

import (
	"crud/dto"
	"crud/entity"
	"time"
)

type AdminParam struct {
	RoleID     uint   `json:"role_id" binding:"required"`
	Username   string `json:"username" binding:"required,max=255"`
	Password   string `json:"password" binding:"required,max=255"`
	IsVerified string `json:"is_verified" binding:"oneof='true' 'false'"`
	IsActive   string `json:"is_active" binding:"oneof='true' 'false'"`
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
	Username   string `json:"username" binding:"required,max=255"`
	Password   string `json:"password" binding:"required,max=255"`
	IsVerified string `json:"is_verified" binding:"oneof='true' 'false'"`
	IsActive   string `json:"is_active" binding:"oneof='true' 'false'"`
}

type LoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterParam struct {
	Username string `json:"username" binding:"required,max=255"`
	Password string `json:"password" binding:"required,max=255"`
}

type SuccessLoginParam struct {
	RoleID      uint   `json:"role_id"`
	Username    string `json:"username"`
	IsVerified  string `json:"is_verified"`
	IsActive    string `json:"is_active"`
	AccessToken string `json:"access_token"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SuccessLogin struct {
	dto.ResponseMeta
	Data SuccessLoginParam `json:"data"`
}

type RegisterApprovalParam struct {
	ID           uint   `json:"id"`
	AdminID      uint   `json:"admin_id"`
	SuperAdminID uint   `json:"super_admin_id"`
	Status       string `json:"status"`
}

type FindAllRegisterApproval struct {
	dto.ResponseMeta
	Data []entity.RegisterApproval `json:"data"`
}

type UpdateRegisterApproval struct {
	SuperAdminID uint   `json:"super_admin_id"`
	Status       string `json:"status" binding:"oneof='pending' 'rejected' 'approved'"`
}

type SuccessUpdateRegisterApproval struct {
	dto.ResponseMeta
	Data RegisterApprovalParam `json:"data"`
}
