package entity

type RegisterApproval struct {
	ID           uint   `gorm:"primary_key"`
	AdminID      uint   `gorm:"column:admin_id"`
	SuperAdminID uint   `gorm:"column:super_admin_id;default:null"`
	Status       string `gorm:"column:status;"`
}

// TableName overrides the table name used by RegisterApproval to `register_approval`
func (RegisterApproval) TableName() string {
	return "register_approval"
}
