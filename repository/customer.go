package repository

import (
	"crud/entity"
	"errors"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}

}

type CustomerInterfaceRepo interface {
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	GetCustomerById(id uint) (entity.Customer, error)
	GetCustomers(name string, email string, page uint) ([]entity.Customer, error)
	UpdateCustomerById(customer *entity.Customer, id uint) (*entity.Customer, error)
	DeleteCustomerById(id uint) (entity.Customer, error)
}

func (repo Customer) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := repo.db.Model(&entity.Customer{}).Create(customer).Error
	return customer, err
}

func (repo Customer) GetCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	repo.db.First(&customer, "id = ? ", id)
	return customer, nil
}

func (repo Customer) GetCustomers(name string, email string, page uint) ([]entity.Customer, error) {
	var customer []entity.Customer

	query := repo.db
	if name != "" {
		query = query.Where("first_name LIKE ?", "%"+name+"%").Or("last_name LIKE ?", "%"+name+"%")
	}

	if email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	limit := 5
	if page > 0 {
		offset := (int(page) - 1) * limit
		query = query.Limit(limit).Offset(offset)
	}

	query.Find(&customer)

	return customer, nil
}

func (repo Customer) UpdateCustomerById(actor *entity.Customer, id uint) (*entity.Customer, error) {
	var err error
	res := repo.db.Model(&actor).Where("id = ?", id).Updates(actor)
	if res.RowsAffected == 0 {
		err = errors.New("id not found")
	}
	return actor, err
}

func (repo Customer) DeleteCustomerById(id uint) (entity.Customer, error) {
	var actor entity.Customer
	var err error
	res := repo.db.Where("id = ? ", id).Delete(&actor)
	if res.RowsAffected == 0 {
		err = errors.New("id not found")
	}
	return actor, err
}
