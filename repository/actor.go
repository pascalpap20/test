package repository

import (
	"crud/entity"
	"errors"
	"gorm.io/gorm"
)

type Actor struct {
	db *gorm.DB
}

func NewActor(dbCrud *gorm.DB) Actor {
	return Actor{
		db: dbCrud,
	}

}

type ActorInterfaceRepo interface {
	CreateActor(actor *entity.Actor) (*entity.Actor, error)
	GetActorById(id uint) (entity.Actor, error)
	GetActors(username string, page uint) ([]entity.Actor, error)
	UpdateActorById(actor *entity.Actor, id uint) (*entity.Actor, error)
	DeleteActorById(id uint) (entity.Actor, error)
	Login(actor *entity.Actor) (*entity.Actor, error)
	Register(actor *entity.Actor) (*entity.Actor, error)
}

func (repo Actor) CreateActor(actor *entity.Actor) (*entity.Actor, error) {
	err := repo.db.Model(&entity.Actor{}).Create(actor).Error
	return actor, err
}

func (repo Actor) GetActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	repo.db.First(&actor, "id = ? ", id)
	return actor, nil
}

func (repo Actor) GetActors(username string, page uint) ([]entity.Actor, error) {
	var actor []entity.Actor

	query := repo.db
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}

	limit := 5
	if page > 0 {
		offset := (int(page) - 1) * limit
		query = query.Limit(limit).Offset(offset)
	}

	query.Find(&actor)

	return actor, nil
}

func (repo Actor) UpdateActorById(actor *entity.Actor, id uint) (*entity.Actor, error) {
	var err error
	res := repo.db.Model(&actor).Where("id = ?", id).Updates(actor)
	if res.RowsAffected == 0 {
		err = errors.New("id not found")
	}
	return actor, err
}

func (repo Actor) DeleteActorById(id uint) (entity.Actor, error) {
	var actor entity.Actor
	var err error
	res := repo.db.Where("id = ? ", id).Delete(&actor)
	if res.RowsAffected == 0 {
		err = errors.New("id not found")
	}
	return actor, err
}

func (repo Actor) Login(actor *entity.Actor) (*entity.Actor, error) {

	// Check if the user exists in the database
	var admin *entity.Actor
	if err := repo.db.Where("username = ?", actor.Username).First(&admin).Error; err != nil {
		err = errors.New("invalid username or password")
		return actor, err
	}

	// Verify the password
	//if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(actor.Password)); err != nil {
	//	err = errors.New("invalid username or password")
	//	return actor, err
	//}
	if admin.Password != actor.Password {
		err := errors.New("invalid username or password")
		return actor, err
	}

	return admin, nil
}

func (repo Actor) Register(actor *entity.Actor) (*entity.Actor, error) {

	// Check if the user exists in the database
	var admin *entity.Actor
	if err := repo.db.Where("username = ?", actor.Username).First(&admin).Error; err == nil {
		err = errors.New("username already exists")
		return admin, err
	}

	tx := repo.db.Begin()

	if err := tx.Model(&entity.Actor{}).Create(&actor).Error; err != nil {
		tx.Rollback()
		return admin, err
	}

	// Initiate the register approval
	err := tx.Model(&entity.RegisterApproval{}).Create(&entity.RegisterApproval{
		AdminID: actor.ID,
		Status:  "pending",
	}).Error

	if err != nil {
		tx.Rollback()
		return admin, err
	}

	tx.Commit()

	return admin, err
}
