package entity

import "time"

type Actor struct {
	ID         uint   `gorm:"primary_key"`
	RoleID     uint   `gorm:"column:role_id"`
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	IsVerified string `gorm:"column:is_verified"`
	IsActive   string `gorm:"column:is_active"`
	Salt       string `gorm:"column:salt"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
